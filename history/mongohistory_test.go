package history

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/DexyProject/dexy-go/types"
	"gopkg.in/mgo.v2/bson"
)

const (
	connection = "mongodb://127.0.0.1:27017"
	block      = 4862998
)

var trans1 = []types.Transaction{
	{
		TransactionID:    BytesNew("0x87012a0d870d47c3c93526c05c4a2f494054c3f4dd8584e94af7d8dd90a535f8"),
		TransactionIndex: 1,
		OrderHash:        types.NewHash("0xEEAD6DBFC7340A56CAEDC044696A168870549A6A7F6F56961E84A54BD9970B8A"),
		BlockNumber:      4862998,
		Timestamp:        types.NewInt(1515233752),
		Taker:            types.HexToAddress("0x997919a608788621dd48b3896f78dcda682fe91d"),
		Maker:            types.HexToAddress("0x9f612fcb422d1971c1be7416c37e3ebc77c0de19"),
		Give:             types.Trade{Token: types.HexToAddress("0x997919a608788621dd48b3896f78dcda682fe91d"), Amount: types.NewInt(300)},
		Get:              types.Trade{Token: types.HexToAddress("0xd26114cd6EE289AccF82350c8d8487fedB8A0C07"), Amount: types.NewInt(3000)},
	},

	{
		TransactionID:    BytesNew("0x89012a0d870d47c3c93526c05c4a2f494054c3f4dd8584e94af7d8dd90a535f8"),
		TransactionIndex: 2,
		OrderHash:        types.NewHash("0xEEAD6DBFC7340A56CAEDC044696A168870549A6A7F6F56961E84A54BD9970B8A"),
		BlockNumber:      4862998,
		Timestamp:        types.NewInt(1515233753),
		Taker:            types.HexToAddress("0x997919a608788621dd48b3896f78dcda682fe91d"),
		Maker:            types.HexToAddress("0x9f612fcb422d1971c1be7416c37e3ebc77c0de19"),
		Give:             types.Trade{Token: types.HexToAddress("0x0000000000000000000000000000000000000000"), Amount: types.NewInt(40)},
		Get:              types.Trade{Token: types.HexToAddress("0xd26114cd6EE289AccF82350c8d8487fedB8A0C07"), Amount: types.NewInt(3000)},
	},

	{
		TransactionID:    BytesNew("0x98012a0d870d47c3c93526c05c4a2f494054c3f4dd8584e94af7d8dd90a535f8"),
		TransactionIndex: 3,
		OrderHash:        types.NewHash("0xEEAD6DBFC7340A56CAEDC044696A168870549A6A7F6F56961E84A54BD9970B8A"),
		BlockNumber:      4862998,
		Timestamp:        types.NewInt(1515233754),
		Taker:            types.HexToAddress("0x997919a608788621dd48b3896f78dcda682fe91d"),
		Maker:            types.HexToAddress("0x9f612fcb422d1971c1be7416c37e3ebc77c0de19"),
		Give:             types.Trade{Token: types.HexToAddress("0x0000000000000000000000000000000000000000"), Amount: types.NewInt(760)},
		Get:              types.Trade{Token: types.HexToAddress("0xd26114cd6EE289AccF82350c8d8487fedB8A0C07"), Amount: types.NewInt(1000)},
	},

	{
		TransactionID:    BytesNew("0x98012a0d870d47c3c93526c05c4a2f494054c3f4dd8584e94af7d8dd90a535f8"),
		TransactionIndex: 4,
		OrderHash:        types.NewHash("0xEEAD6DBFC7340A56CAEDC044696A168870549A6A7F6F56961E84A54BD9970B8A"),
		BlockNumber:      4862998,
		Timestamp:        types.NewInt(1515233755),
		Taker:            types.HexToAddress("0x997919a608788621dd48b3896f78dcda682fe91d"),
		Maker:            types.HexToAddress("0x9f612fcb422d1971c1be7416c37e3ebc77c0de19"),
		Give:             types.Trade{Token: types.HexToAddress("0x9f612fcb422d1971c1be7416c37e3ebc77c0de19"), Amount: types.NewInt(303)},
		Get:              types.Trade{Token: types.HexToAddress("0xd26114cd6EE289AccF82350c8d8487fedB8A0C07"), Amount: types.NewInt(3003)},
	},

	{
		TransactionID:    BytesNew("0x98012a0d870d47c3c93526c05c4a2f494054c3f4dd8584e94af7d8dd90a535f8"),
		TransactionIndex: 5,
		OrderHash:        types.NewHash("0xEEAD6DBFC7340A56CAEDC044696A168870549A6A7F6F56961E84A54BD9970B8A"),
		BlockNumber:      4862998,
		Timestamp:        types.NewInt(1515233756),
		Taker:            types.HexToAddress("0x997919a608788621dd48b3896f78dcda682fe91d"),
		Maker:            types.HexToAddress("0x9f612fcb422d1971c1be7416c37e3ebc77c0de19"),
		Give:             types.Trade{Token: types.HexToAddress("0x9f612fcb422d1971c1be7416c37e3ebc77c0de19"), Amount: types.NewInt(1235)},
		Get:              types.Trade{Token: types.HexToAddress("0xd26114cd6EE289AccF82350c8d8487fedB8A0C07"), Amount: types.NewInt(884532)},
	},

	{
		TransactionID:    BytesNew("0x98012a0d870d47c3c93526c05c4a2f494054c3f4dd8584e94af7d8dd90a535f8"),
		TransactionIndex: 6,
		OrderHash:        types.NewHash("0xEEAD6DBFC7340A56CAEDC044696A168870549A6A7F6F56961E84A54BD9970B8A"),
		BlockNumber:      4862998,
		Timestamp:        types.NewInt(1515233757),
		Taker:            types.HexToAddress("0x997919a608788621dd48b3896f78dcda682fe91d"),
		Maker:            types.HexToAddress("0x9f612fcb422d1971c1be7416c37e3ebc77c0de19"),
		Give:             types.Trade{Token: types.HexToAddress("0x0000000000000000000000000000000000000000"), Amount: types.NewInt(100)},
		Get:              types.Trade{Token: types.HexToAddress("0xd26114cd6EE289AccF82350c8d8487fedB8A0C07"), Amount: types.NewInt(123491)},
	},

	{
		TransactionID:    BytesNew("0x98012a0d870d47c3c93526c05c4a2f494054c3f4dd8584e94af7d8dd90a535f8"),
		TransactionIndex: 7,
		OrderHash:        types.NewHash("0xEEAD6DBFC7340A56CAEDC044696A168870549A6A7F6F56961E84A54BD9970B8A"),
		BlockNumber:      4862999,
		Timestamp:        types.NewInt(1515233751),
		Taker:            types.HexToAddress("0x997919a608788621dd48b3896f78dcda682fe91d"),
		Maker:            types.HexToAddress("0x9f612fcb422d1971c1be7416c37e3ebc77c0de19"),
		Give:             types.Trade{Token: types.HexToAddress("0x0000000000000000000000000000000000000000"), Amount: types.NewInt(1324)},
		Get:              types.Trade{Token: types.HexToAddress("0xd26114cd6EE289AccF82350c8d8487fedB8A0C07"), Amount: types.NewInt(123415)},
	},
}
var multiToken = []types.Transaction{
	{
		TransactionID:    BytesNew("0x87012a0d870d47c3c93526c05c4a2f494054c3f4dd8584e94af7d8dd90a535f8"),
		TransactionIndex: 1,
		OrderHash:        types.NewHash("0xEEAD6DBFC7340A56CAEDC044696A168870549A6A7F6F56961E84A54BD9970B8A"),
		BlockNumber:      4862998,
		Timestamp:        types.NewInt(1515233752),
		Taker:            types.HexToAddress("0x997919a608788621dd48b3896f78dcda682fe91d"),
		Maker:            types.HexToAddress("0x9f612fcb422d1971c1be7416c37e3ebc77c0de19"),
		Give:             types.Trade{Token: types.HexToAddress("0x0000000000000000000000000000000000000000"), Amount: types.NewInt(300)},
		Get:              types.Trade{Token: types.HexToAddress("09dfd26114cd6EE289AccF82350c8d8487fedB8A0C"), Amount: types.NewInt(3000)},
	},
	{
		TransactionID:    BytesNew("0x87012a0d870d47c3c93526c05c4a2f494054c3f4dd8584e94af7d8dd90a535f8"),
		TransactionIndex: 1,
		OrderHash:        types.NewHash("0xEEAD6DBFC7340A56CAEDC044696A168870549A6A7F6F56961E84A54BD9970B8A"),
		BlockNumber:      4862998,
		Timestamp:        types.NewInt(1515233752),
		Taker:            types.HexToAddress("0x997919a608788621dd48b3896f78dcda682fe91d"),
		Maker:            types.HexToAddress("0x9f612fcb422d1971c1be7416c37e3ebc77c0de19"),
		Give:             types.Trade{Token: types.HexToAddress("0x0000000000000000000000000000000000000000"), Amount: types.NewInt(300)},
		Get:              types.Trade{Token: types.HexToAddress("0xd26114cd6EE289AccF82350c8d8487fedB8A0C07"), Amount: types.NewInt(3000)},
	},
	{
		TransactionID:    BytesNew("0x87012a0d870d47c3c93526c05c4a2f494054c3f4dd8584e94af7d8dd90a535f8"),
		TransactionIndex: 1,
		OrderHash:        types.NewHash("0xEEAD6DBFC7340A56CAEDC044696A168870549A6A7F6F56961E84A54BD9970B8A"),
		BlockNumber:      4862998,
		Timestamp:        types.NewInt(1515233752),
		Taker:            types.HexToAddress("0x997919a608788621dd48b3896f78dcda682fe91d"),
		Maker:            types.HexToAddress("0x9f612fcb422d1971c1be7416c37e3ebc77c0de19"),
		Give:             types.Trade{Token: types.HexToAddress("0x997919a608788621dd48b3896f78dcda682fe91d"), Amount: types.NewInt(300)},
		Get:              types.Trade{Token: types.HexToAddress("0x0000000000000000000000000000000000000000"), Amount: types.NewInt(3000)},
	},
	{
		TransactionID:    BytesNew("0x87012a0d870d47c3c93526c05c4a2f494054c3f4dd8584e94af7d8dd90a535f8"),
		TransactionIndex: 1,
		OrderHash:        types.NewHash("0xEEAD6DBFC7340A56CAEDC044696A168870549A6A7F6F56961E84A54BD9970B8A"),
		BlockNumber:      4862998,
		Timestamp:        types.NewInt(1515233752),
		Taker:            types.HexToAddress("0x997919a608788621dd48b3896f78dcda682fe91d"),
		Maker:            types.HexToAddress("0x9f612fcb422d1971c1be7416c37e3ebc77c0de19"),
		Give:             types.Trade{Token: types.HexToAddress("0x997919a608788621dd48b3896f78dcda682fe91d"), Amount: types.NewInt(300)},
		Get:              types.Trade{Token: types.HexToAddress("0xd26114cd6EE289AccF82350c8d8487fedB8A0C07"), Amount: types.NewInt(3000)},
	},
	{
		TransactionID:    BytesNew("0x87012a0d870d47c3c93526c05c4a2f494054c3f4dd8584e94af7d8dd90a535f8"),
		TransactionIndex: 1,
		OrderHash:        types.NewHash("0xEEAD6DBFC7340A56CAEDC044696A168870549A6A7F6F56961E84A54BD9970B8A"),
		BlockNumber:      4862998,
		Timestamp:        types.NewInt(1515233752),
		Taker:            types.HexToAddress("0x997919a608788621dd48b3896f78dcda682fe91d"),
		Maker:            types.HexToAddress("0x9f612fcb422d1971c1be7416c37e3ebc77c0de19"),
		Give:             types.Trade{Token: types.HexToAddress("0x0000000000000000000000000000000000000000"), Amount: types.NewInt(300)},
		Get:              types.Trade{Token: types.HexToAddress("0xd26114cd6EE289AccF82350c8d8487fedB8A0C07"), Amount: types.NewInt(3000)},
	},
}

func TestMongoHistory_AggregateTransactions(t *testing.T) {
	mgoConnection, err := NewHistoryAggregation(connection)
	if err != nil {
		t.Errorf("could not establish new connection")
	}
	mgoConnection.session.Clone()
	defer mgoConnection.session.Close()

	var matchTicks []types.Transaction

	c := mgoConnection.session.DB(DBName).C(FileName)
	// Test mgo queries
	matchBlock := bson.M{"$match": bson.M{"transactions.block": block}}

	err = c.Pipe([]bson.M{matchBlock}).All(&matchTicks)
	if err != nil {
		t.Errorf("could not match data")
	} else {
		fmt.Println(matchTicks)
	} // test mongo block matching

	ticks, err := mgoConnection.AggregateTransactions(block, trans1)
	if ticks == nil {
		t.Errorf("could not aggregate transactions")
	}
	if err != nil {
		fmt.Println(err)
	}
	b, err := json.Marshal(ticks)
	fmt.Println(string(b))
}

func TestMultiToken(t *testing.T) {
	mgoConnection, err := NewHistoryAggregation(connection)
	if err != nil {
		t.Errorf("could not establish new connection")
	}
	mgoConnection.session.Clone()
	defer mgoConnection.session.Close()

	var matchTicks []types.Transaction

	c := mgoConnection.session.DB(DBName).C(FileName)

	matchBlock := bson.M{"$match": bson.M{"transactions.block": block}}
	err = c.Pipe([]bson.M{matchBlock}).All(&matchTicks)
	if err != nil {
		t.Errorf("could not match data")
	}
	ticks, err := mgoConnection.AggregateTransactions(block, multiToken)
	if ticks == nil {
		t.Errorf("could not aggregate transactions")
	}
	if err != nil {
		fmt.Println(err)
	}
	multiTokenMap := groupTokens(multiToken)
	if len(multiTokenMap) <= 1 {
		t.Errorf("tokens not grouped properly")
	}

}

func TestCalcOpenCloseIndex(t *testing.T) {
	var openIndex, closeIndex uint
	openIndex, closeIndex = calcOpenCloseIndex(trans1)
	if openIndex == 0 || closeIndex == 0 {
		t.Errorf("could not calculate open and close indices")
	}
}

func TestGetPrices(t *testing.T) {
	err, txindex := getPrices(trans1)
	if err, txindex == nil {
		t.Errorf("could not generate prices")
	}
}

func TestGroupTokens(t *testing.T) {
	m := groupTokens(trans1)
	fmt.Println(m)
}

func BytesNew(bytes string) types.Bytes {
	b := types.Bytes{}
	b.UnmarshalText([]byte(bytes))
	return b
}
