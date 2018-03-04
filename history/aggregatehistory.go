package history

import (
	"fmt"
	"math/big"

	"github.com/DexyProject/dexy-go/types"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/DexyProject/dexy-go/contracts"
	"math"
)

type HistoryAggregation struct {
	connection string
	session    *mgo.Session
}

type ERC20DecimalsCalc struct {
	erc20 contracts.ERC20
	decimals map[types.Address]uint8
}

func NewHistoryAggregation(connection string) (*HistoryAggregation, error) {
	session, err := mgo.Dial(connection)
	if err != nil {
		return nil, fmt.Errorf("could not connect to mongo database")
	}

	return &HistoryAggregation{connection: connection, session: session}, nil
}

func NewERC20DecimalsCalc(erc20 contracts.ERC20) *ERC20DecimalsCalc {
	return &ERC20DecimalsCalc{
		erc20: erc20,
		decimals: make(map[types.Address]uint8),
	}
}

func (history *HistoryAggregation) AggregateTransactions(block int64) ([]types.Tick, error) {
	ERC20DecimalsCalc := NewERC20DecimalsCalc(contracts.ERC20{})
	session := history.session.Clone()
	defer session.Close()
	c := session.DB(DBName).C(FileName)


	var ticks []types.Tick
	var transactions []types.Transaction

	matchBlock := bson.M{"$match": bson.M{"$transactions.block": block}}

	err := c.Pipe([]bson.M{matchBlock}).All(&transactions)

	if err != nil {
		return nil, fmt.Errorf("could not retrieve transactions")
	}

	mappedTokens := groupTokens(transactions)
	for token := range mappedTokens {
		pair := getPair(token)
		volume := calcVolume(mappedTokens[token])
		prices, txindex := getPrices(mappedTokens[token], ERC20DecimalsCalc)
		openIndex, closeIndex := calcOpenCloseIndex(mappedTokens[token])
		openPrice, closePrice := calcOpenClosePrice(prices, txindex, openIndex, closeIndex)
		high, low := calcHighLow(prices)

		ticks = append(
			ticks, types.Tick{
				Pair:      pair,
				Block:     block,
				Volume:    types.Int{*volume},
				Open:      openPrice,
				Close:     closePrice,
				High:      high,
				Low:       low,
				Timestamp: mappedTokens[token][1].Timestamp,
			},
		)
	}

	return ticks, nil
}

func calcVolume(transactions []types.Transaction) *big.Int {
	volume := new(big.Int)
	for _, tt := range transactions {
		switch {
		case tt.Give.Token != types.HexToAddress(types.ETH_ADDRESS):
			volume.Add(volume, &tt.Give.Amount.Int)
		case tt.Get.Token != types.HexToAddress(types.ETH_ADDRESS):
			volume.Add(volume, &tt.Get.Amount.Int)
		}
	}

	return volume
}

func calcHighLow(prices []float64) (float64, float64) {
	high, low := prices[0], prices[0]
	for _, p := range prices {
		if high < p {
			high = p
		}
		if low > p {
			low = p
		}
	}

	return high, low
}

func getPrices(transactions []types.Transaction, erc20 *ERC20DecimalsCalc) ([]float64, []uint) {
	var prices []float64
	var txindex []uint
	for _, tt := range transactions {
		newPrice, _ := erc20.CalcPrice(tt, types.HexToAddress(types.ETH_ADDRESS))
		prices = append(prices, newPrice)
		txindex = append(txindex, tt.TransactionIndex)
	}

	return prices, txindex
}

func calcOpenCloseIndex(transactions []types.Transaction) (uint, uint) {
	openIndex, closeIndex := transactions[0].TransactionIndex, transactions[0].TransactionIndex
	for _, tt := range transactions {
		if openIndex > tt.TransactionIndex {
			openIndex = tt.TransactionIndex
		}
		if closeIndex < tt.TransactionIndex {
			closeIndex = tt.TransactionIndex
		}
	}

	return openIndex, closeIndex
}

func calcOpenClosePrice(prices []float64, txindex []uint, OpenIndex, CloseIndex uint) (float64, float64) {
	var openPrice, closePrice float64
	for i, tt := range txindex {
		switch tt {
		case OpenIndex:
			openPrice = prices[i]
		case CloseIndex:
			closePrice = prices[i]
		}
	}

	return openPrice, closePrice
}

func getPair(token types.Address) types.Pair {
	return types.Pair{types.HexToAddress(types.ETH_ADDRESS), token}
}

func groupTokens(transactions []types.Transaction) map[types.Address][]types.Transaction {
	m := make(map[types.Address][]types.Transaction)
	for _, t := range transactions {
		if t.Get.Token == types.HexToAddress(types.ETH_ADDRESS) {
			m[t.Give.Token] = append(m[t.Give.Token], t)
		} else {
			m[t.Get.Token] = append(m[t.Get.Token], t)
		}
	}

	return m
}

func (e *ERC20DecimalsCalc) CalcPrice(t types.Transaction, base types.Address) (float64, error) {
	decimals, err := e.erc20.Decimals(nil)
	if err != nil {
		return 0.0, fmt.Errorf("could not get decimals() from contract")
	}
	if t.Give.Amount.Sign() <= 0 || t.Get.Amount.Sign() <= 0 {
		return 0.0, fmt.Errorf("can not divide by zero")
	}

	giveFloat := new(big.Float).SetInt(&t.Give.Amount.Int)
	getFloat := new(big.Float).SetInt(&t.Get.Amount.Int)

	decimalsFloat := float64(decimals)
	if t.Get.Token == base {
		price, _ := new(big.Float).Quo(getFloat, giveFloat).Float64()
		e.decimals[t.Give.Token] = decimals
		return (price / math.Pow(10.0, decimalsFloat)), nil
	}
	price, _ := new(big.Float).Quo(giveFloat, getFloat).Float64()
	e.decimals[t.Get.Token] = decimals
	return (price / math.Pow(10.0, decimalsFloat)), nil
}
