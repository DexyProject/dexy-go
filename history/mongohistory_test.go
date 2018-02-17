package history

import (
	"testing"
)

const (
	connection = "mongodb://127.0.0.1:27017"
	block = 4862998
)

//var historyData = []types.Transaction{
//	{
//		TransactionID: types.NewBytes("0x87012a0d870d47c3c93526c05c4a2f494054c3f4dd8584e94af7d8dd90a535f8"),
//		OrderHash: types.NewHash("0xEEAD6DBFC7340A56CAEDC044696A168870549A6A7F6F56961E84A54BD9970B8A"),
//		BlockNumber: 4862998,
//		Timestamp: types.NewInt(1515233752),
//		Taker: types.HexToAddress("0x997919a608788621dd48b3896f78dcda682fe91d"),
//		Maker: types.HexToAddress("0x9f612fcb422d1971c1be7416c37e3ebc77c0de19"),
//		Give: types.Trade{Token: types.HexToAddress("0x997919a608788621dd48b3896f78dcda682fe91d"), Amount: types.NewInt(300)},
//		Get:  types.Trade{Token: types.HexToAddress("0xd26114cd6EE289AccF82350c8d8487fedB8A0C07"), Amount: types.NewInt(3000)},
//	},
//
//	{
//		TransactionID: "0x89012a0d870d47c3c93526c05c4a2f494054c3f4dd8584e94af7d8dd90a535f8",
//		OrderHash: types.NewHash("0xEEAD6DBFC7340A56CAEDC044696A168870549A6A7F6F56961E84A54BD9970B8A"),
//		BlockNumber: 4862998,
//		Timestamp: types.NewInt(1515233753),
//		Taker: types.HexToAddress("0x997919a608788621dd48b3896f78dcda682fe91d"),
//		Maker: types.HexToAddress("0x9f612fcb422d1971c1be7416c37e3ebc77c0de19"),
//		Give: types.Trade{Token: types.HexToAddress("0x0000000000000000000000000000000000000000"), Amount: types.NewInt(40)},
//		Get:  types.Trade{Token: types.HexToAddress("0xd26114cd6EE289AccF82350c8d8487fedB8A0C07"), Amount: types.NewInt(3000)},
//	},
//
//	{
//		TransactionID: "0x98012a0d870d47c3c93526c05c4a2f494054c3f4dd8584e94af7d8dd90a535f8",
//		OrderHash: types.NewHash("0xEEAD6DBFC7340A56CAEDC044696A168870549A6A7F6F56961E84A54BD9970B8A"),
//		BlockNumber: 4862998,
//		Timestamp: types.NewInt(1515233754),
//		Taker: types.HexToAddress("0x997919a608788621dd48b3896f78dcda682fe91d"),
//		Maker: types.HexToAddress("0x9f612fcb422d1971c1be7416c37e3ebc77c0de19"),
//		Give: types.Trade{Token: types.HexToAddress("0x0000000000000000000000000000000000000000"), Amount: types.NewInt(760)},
//		Get:  types.Trade{Token: types.HexToAddress("0xd26114cd6EE289AccF82350c8d8487fedB8A0C07"), Amount: types.NewInt(1000)},
//	},
//
//	{
//		TransactionID: "0x98012a0d870d47c3c93526c05c4a2f494054c3f4dd8584e94af7d8dd90a535f8",
//		OrderHash: types.NewHash("0xEEAD6DBFC7340A56CAEDC044696A168870549A6A7F6F56961E84A54BD9970B8A"),
//		BlockNumber: 4862998,
//		Timestamp: types.NewInt(1515233755),
//		Taker: types.HexToAddress("0x997919a608788621dd48b3896f78dcda682fe91d"),
//		Maker: types.HexToAddress("0x9f612fcb422d1971c1be7416c37e3ebc77c0de19"),
//		Give: types.Trade{Token: types.HexToAddress("0x9f612fcb422d1971c1be7416c37e3ebc77c0de19"), Amount: types.NewInt(303)},
//		Get:  types.Trade{Token: types.HexToAddress("0xd26114cd6EE289AccF82350c8d8487fedB8A0C07"), Amount: types.NewInt(3003)},
//	},
//
//	{
//		TransactionID: "0x98012a0d870d47c3c93526c05c4a2f494054c3f4dd8584e94af7d8dd90a535f8",
//		OrderHash: types.NewHash("0xEEAD6DBFC7340A56CAEDC044696A168870549A6A7F6F56961E84A54BD9970B8A"),
//		BlockNumber: 4862998,
//		Timestamp: types.NewInt(1515233756),
//		Taker: types.HexToAddress("0x997919a608788621dd48b3896f78dcda682fe91d"),
//		Maker: types.HexToAddress("0x9f612fcb422d1971c1be7416c37e3ebc77c0de19"),
//		Give: types.Trade{Token: types.HexToAddress("0x9f612fcb422d1971c1be7416c37e3ebc77c0de19"), Amount: types.NewInt(1235)},
//		Get:  types.Trade{Token: types.HexToAddress("0xd26114cd6EE289AccF82350c8d8487fedB8A0C07"), Amount: types.NewInt(884532)},
//	},
//
//	{
//		TransactionID: "0x98012a0d870d47c3c93526c05c4a2f494054c3f4dd8584e94af7d8dd90a535f8",
//		OrderHash: types.NewHash("0xEEAD6DBFC7340A56CAEDC044696A168870549A6A7F6F56961E84A54BD9970B8A"),
//		BlockNumber: 4862998,
//		Timestamp: types.NewInt(1515233757),
//		Taker: types.HexToAddress("0x997919a608788621dd48b3896f78dcda682fe91d"),
//		Maker: types.HexToAddress("0x9f612fcb422d1971c1be7416c37e3ebc77c0de19"),
//		Give: types.Trade{Token: types.HexToAddress("0x0000000000000000000000000000000000000000"), Amount: types.NewInt(100)},
//		Get:  types.Trade{Token: types.HexToAddress("0xd26114cd6EE289AccF82350c8d8487fedB8A0C07"), Amount: types.NewInt(123491)},
//	},
//
//	{
//		TransactionID: "0x98012a0d870d47c3c93526c05c4a2f494054c3f4dd8584e94af7d8dd90a535f8",
//		OrderHash: types.NewHash("0xEEAD6DBFC7340A56CAEDC044696A168870549A6A7F6F56961E84A54BD9970B8A"),
//		BlockNumber: 4862999,
//		Timestamp: types.NewInt(1515233751),
//		Taker: types.HexToAddress("0x997919a608788621dd48b3896f78dcda682fe91d"),
//		Maker: types.HexToAddress("0x9f612fcb422d1971c1be7416c37e3ebc77c0de19"),
//		Give: types.Trade{Token: types.HexToAddress("0x0000000000000000000000000000000000000000"), Amount: types.NewInt(1324)},
//		Get:  types.Trade{Token: types.HexToAddress("0xd26114cd6EE289AccF82350c8d8487fedB8A0C07"), Amount: types.NewInt(123415)},
//	},
//
//
// }

func TestMongoHistory_AggregateTransactions(t *testing.T) {
	mgoConnection, err := NewMongoHistory(connection)

	response, err := (mgoConnection).AggregateTransactions(block)
	if response == nil {
		t.Errorf("Aggregation returned empty")
	}
	if err != nil {
		t.Errorf("Query not completed")
	}
}

