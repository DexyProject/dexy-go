package history

import (
	"testing"
	"github.com/DexyProject/dexy-go/types"
	"fmt"
	"gopkg.in/mgo.v2/bson"
)

const (
	connection = "mongodb://127.0.0.1:27017"
	block = 4862998
)


var trans1 = []types.Transaction{
	{
		TransactionID: BytesNew("0x87012a0d870d47c3c93526c05c4a2f494054c3f4dd8584e94af7d8dd90a535f8"),
		TransactionIndex: 1,
		OrderHash: types.NewHash("0xEEAD6DBFC7340A56CAEDC044696A168870549A6A7F6F56961E84A54BD9970B8A"),
		BlockNumber: 4862998,
		Timestamp: types.NewInt(1515233752),
		Taker: types.HexToAddress("0x997919a608788621dd48b3896f78dcda682fe91d"),
		Maker: types.HexToAddress("0x9f612fcb422d1971c1be7416c37e3ebc77c0de19"),
		Give: types.Trade{Token: types.HexToAddress("0x997919a608788621dd48b3896f78dcda682fe91d"), Amount: types.NewInt(300)},
		Get:  types.Trade{Token: types.HexToAddress("0xd26114cd6EE289AccF82350c8d8487fedB8A0C07"), Amount: types.NewInt(3000)},
	},

	{
		TransactionID: BytesNew("0x89012a0d870d47c3c93526c05c4a2f494054c3f4dd8584e94af7d8dd90a535f8"),
		TransactionIndex: 2,
		OrderHash: types.NewHash("0xEEAD6DBFC7340A56CAEDC044696A168870549A6A7F6F56961E84A54BD9970B8A"),
		BlockNumber: 4862998,
		Timestamp: types.NewInt(1515233753),
		Taker: types.HexToAddress("0x997919a608788621dd48b3896f78dcda682fe91d"),
		Maker: types.HexToAddress("0x9f612fcb422d1971c1be7416c37e3ebc77c0de19"),
		Give: types.Trade{Token: types.HexToAddress("0x0000000000000000000000000000000000000000"), Amount: types.NewInt(40)},
		Get:  types.Trade{Token: types.HexToAddress("0xd26114cd6EE289AccF82350c8d8487fedB8A0C07"), Amount: types.NewInt(3000)},
	},

	{
		TransactionID: BytesNew("0x98012a0d870d47c3c93526c05c4a2f494054c3f4dd8584e94af7d8dd90a535f8"),
		TransactionIndex: 3,
		OrderHash: types.NewHash("0xEEAD6DBFC7340A56CAEDC044696A168870549A6A7F6F56961E84A54BD9970B8A"),
		BlockNumber: 4862998,
		Timestamp: types.NewInt(1515233754),
		Taker: types.HexToAddress("0x997919a608788621dd48b3896f78dcda682fe91d"),
		Maker: types.HexToAddress("0x9f612fcb422d1971c1be7416c37e3ebc77c0de19"),
		Give: types.Trade{Token: types.HexToAddress("0x0000000000000000000000000000000000000000"), Amount: types.NewInt(760)},
		Get:  types.Trade{Token: types.HexToAddress("0xd26114cd6EE289AccF82350c8d8487fedB8A0C07"), Amount: types.NewInt(1000)},
	},

	{
		TransactionID: BytesNew("0x98012a0d870d47c3c93526c05c4a2f494054c3f4dd8584e94af7d8dd90a535f8"),
		TransactionIndex: 4,
		OrderHash: types.NewHash("0xEEAD6DBFC7340A56CAEDC044696A168870549A6A7F6F56961E84A54BD9970B8A"),
		BlockNumber: 4862998,
		Timestamp: types.NewInt(1515233755),
		Taker: types.HexToAddress("0x997919a608788621dd48b3896f78dcda682fe91d"),
		Maker: types.HexToAddress("0x9f612fcb422d1971c1be7416c37e3ebc77c0de19"),
		Give: types.Trade{Token: types.HexToAddress("0x9f612fcb422d1971c1be7416c37e3ebc77c0de19"), Amount: types.NewInt(303)},
		Get:  types.Trade{Token: types.HexToAddress("0xd26114cd6EE289AccF82350c8d8487fedB8A0C07"), Amount: types.NewInt(3003)},
	},

	{
		TransactionID: BytesNew("0x98012a0d870d47c3c93526c05c4a2f494054c3f4dd8584e94af7d8dd90a535f8"),
		TransactionIndex: 5,
		OrderHash: types.NewHash("0xEEAD6DBFC7340A56CAEDC044696A168870549A6A7F6F56961E84A54BD9970B8A"),
		BlockNumber: 4862998,
		Timestamp: types.NewInt(1515233756),
		Taker: types.HexToAddress("0x997919a608788621dd48b3896f78dcda682fe91d"),
		Maker: types.HexToAddress("0x9f612fcb422d1971c1be7416c37e3ebc77c0de19"),
		Give: types.Trade{Token: types.HexToAddress("0x9f612fcb422d1971c1be7416c37e3ebc77c0de19"), Amount: types.NewInt(1235)},
		Get:  types.Trade{Token: types.HexToAddress("0xd26114cd6EE289AccF82350c8d8487fedB8A0C07"), Amount: types.NewInt(884532)},
	},

	{
		TransactionID: BytesNew("0x98012a0d870d47c3c93526c05c4a2f494054c3f4dd8584e94af7d8dd90a535f8"),
		TransactionIndex: 6,
		OrderHash: types.NewHash("0xEEAD6DBFC7340A56CAEDC044696A168870549A6A7F6F56961E84A54BD9970B8A"),
		BlockNumber: 4862998,
		Timestamp: types.NewInt(1515233757),
		Taker: types.HexToAddress("0x997919a608788621dd48b3896f78dcda682fe91d"),
		Maker: types.HexToAddress("0x9f612fcb422d1971c1be7416c37e3ebc77c0de19"),
		Give: types.Trade{Token: types.HexToAddress("0x0000000000000000000000000000000000000000"), Amount: types.NewInt(100)},
		Get:  types.Trade{Token: types.HexToAddress("0xd26114cd6EE289AccF82350c8d8487fedB8A0C07"), Amount: types.NewInt(123491)},
	},

	{
		TransactionID: BytesNew("0x98012a0d870d47c3c93526c05c4a2f494054c3f4dd8584e94af7d8dd90a535f8"),
		TransactionIndex: 7,
		OrderHash: types.NewHash("0xEEAD6DBFC7340A56CAEDC044696A168870549A6A7F6F56961E84A54BD9970B8A"),
		BlockNumber: 4862999,
		Timestamp: types.NewInt(1515233751),
		Taker: types.HexToAddress("0x997919a608788621dd48b3896f78dcda682fe91d"),
		Maker: types.HexToAddress("0x9f612fcb422d1971c1be7416c37e3ebc77c0de19"),
		Give: types.Trade{Token: types.HexToAddress("0x0000000000000000000000000000000000000000"), Amount: types.NewInt(1324)},
		Get:  types.Trade{Token: types.HexToAddress("0xd26114cd6EE289AccF82350c8d8487fedB8A0C07"), Amount: types.NewInt(123415)},
	},


 }

 var historyData = []Transactions {
 	{trans1},
 	{trans1},
 }

func TestMongoHistory_AggregateTransactions(t *testing.T) {
	mgoConnection, err := NewMongoHistory(connection)
	if err != nil {
		t.Errorf("could not establish new connection")
	}
	mgoConnection.session.Clone()
	defer mgoConnection.session.Close()

	matchTicks := []bson.M{}
	sortTicks := []bson.M{}
	filterTicks := []bson.M{}

	c := mgoConnection.session.DB(DBName).C(FileName)
	// Test mgo queries
	matchBlock := bson.M{"$match": bson.M{"transactions.block": block}}
	sortTimestamp := bson.M{"$sort": bson.M{"transactions.timestamp": -1}}
	groupGetTokens := bson.M{
		"$group": bson.M{
			"filter": bson.M{
				"input": "$transactions",
				"as":    "tt",
				"cond": bson.M{"$and": []interface{}{
					bson.M{"$eq": []interface{}{"$$tt.get.token", "$$tt.get.token"}},
					bson.M{"$ne": []interface{}{"$$tt.get.token", types.ETH_ADDRESS}},
				},
				},
			},
		},
	}
	groupGiveTokens := bson.M{
		"$group": bson.M{
			"filter": bson.M{
				"input": "$transactions",
				"as":    "tt",
				"cond": bson.M{"$and": []interface{}{
					bson.M{"$eq": []interface{}{"$$tt.give.token", "$$tt.give.token"}},
					bson.M{"$ne": []interface{}{"$$tt.give.token", types.ETH_ADDRESS}},
				},
				},
			},
		},
	}
	groupTokens := bson.M{
		"$group": bson.M{
			"filter": bson.M{
				"input": "$transactions",
				"as":    "tt",
				"cond": bson.M{"$and": []interface{}{
					bson.M{"$eq": []interface{}{"$$tt.give.token", "$$tt.get.token"}},
					bson.M{"$ne": []interface{}{"$$tt.get.token", types.ETH_ADDRESS}},
				},
				},
			},
		},
	}
	err = c.Pipe([]bson.M{matchBlock}).All(&matchTicks)
	if err != nil {
		t.Errorf("could not match data")
	} else {
		fmt.Println(matchTicks)
	}
	err = c.Pipe([]bson.M{sortTimestamp}).All(&sortTicks)
	if err != nil {
		t.Errorf("could not sort by timestamp")
	} else {
		fmt.Println(sortTimestamp)
	}
	mgoError := c.Pipe([]bson.M{groupGetTokens, groupGiveTokens, groupTokens}).All(&filterTicks).Error()
	if mgoError != "" {
		t.Errorf(mgoError)
	} else {
		fmt.Println(filterTicks)
	}
}


func TestCalcOpenCloseIndex(t *testing.T) {
	var openIndex, closeIndex uint
	for _, tt := range historyData {
		openIndex, closeIndex = calcOpenCloseIndex(tt.Transactions)
	}
	if openIndex == 0 || closeIndex == 0 {
		t.Errorf("could not calculate open and close indices")
	}

	//fmt.Println(openIndex, closeIndex)
}

func TestCalcOpenClosePrice(t *testing.T) {
	var openPrice, closePrice float64
	prices := getPrices(historyData[1].Transactions)
	openIndex, closeIndex := calcOpenCloseIndex(historyData[1].Transactions)
	openPrice, closePrice = calcOpenClosePrice(prices, openIndex, closeIndex)
	if openPrice == 0 || closePrice == 0 {
		t.Errorf("could not calculate open and close prices")
	}
	//fmt.Println("open:", openPrice, "close:", closePrice)

}

func TestGetPrices(t *testing.T) {
	err := getPrices(historyData[1].Transactions)
	if err == nil {
		t.Errorf("could not generate prices")
	}

	//fmt.Println("transactionindex, price:", err)
}

func TestGetPair(t *testing.T) {
	newPair := getPair(historyData[1].Transactions)
	if (types.Pair{}) == newPair {
		t.Errorf("could not generate pair from transactions")
	}

	//fmt.Println("pair:", newPair)
}

func TestCalcHighLow(t *testing.T) {
	prices := getPrices(historyData[1].Transactions)
	high, low := calcHighLow(prices)
	if high == 0 || low == 0 {
		t.Errorf("could not retrieve high and low prices")
	}

	//fmt.Println("high:",high, "low:",low)
}

func BytesNew(bytes string) (types.Bytes) {
	b := types.Bytes{}
	b.UnmarshalText([]byte(bytes))
	return b
}