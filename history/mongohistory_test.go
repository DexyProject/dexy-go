package history

import (
	"testing"
	"github.com/DexyProject/dexy-go/types"
	"fmt"
)

const (
	connection = "mongodb://127.0.0.1:27017"
	block = 4862998
)

type Transactions struct {
	Transactions   []types.Transaction `json:"transactions" bson:"transactions"`
}

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
	mgoConnection.session.Clone()
	defer mgoConnection.session.Close()

	c := mgoConnection.session.DB(DBName).C(FileName)
	c.Insert(historyData)

	response, err := (mgoConnection).AggregateTransactions(block)
	if response == nil {
		t.Errorf("Aggregation returned empty")
	}
	if err != nil {
		t.Errorf("Query not completed")
	}
}

func TestCalcOpenCloseIndex(t *testing.T) {
	var openIndex, closeIndex uint
	for _, tt := range historyData {
		openIndex, closeIndex = CalcOpenCloseIndex(tt.Transactions)
	}
	if openIndex, closeIndex == 0 {
		t.Errorf("could not calculate open and close indices")
	}

	fmt.Println(openIndex, closeIndex)
}

func TestCalcOpenClosePrice(t *testing.T) {
	var openPrice, closePrice float64
	prices := GetPrices(historyData[1].Transactions)
	openIndex, closeIndex := CalcOpenCloseIndex(historyData[1].Transactions)
	openPrice, closePrice = CalcOpenClosePrice(prices, openIndex, closeIndex)
	if openPrice, closePrice == 0 {
		t.Errorf("could not calculate open and close prices")
	}



}

func TestGetPrices(t *testing.T) {
	err := GetPrices(historyData[1].Transactions)
	if err == nil {
		t.Errorf("could not generate prices")
	}

	fmt.Println(err)
}

func TestGetPair(t *testing.T) {
	newPair := GetPair(historyData[1].Transactions)
	if (types.Pair{}) == newPair {
		t.Errorf("could not generate pair from transactions")
	}

	fmt.Println(newPair)
}

func TestCalcHighLow(t *testing.T) {

}

func BytesNew(bytes string) (types.Bytes) {
	b := types.Bytes{}
	b.UnmarshalText([]byte(bytes))
	return b
}