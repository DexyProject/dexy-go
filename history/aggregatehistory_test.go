package history

import (
	"fmt"
	"testing"

	"github.com/DexyProject/dexy-go/repositories"
	"github.com/DexyProject/dexy-go/types"
)

const (
	connection = "mongodb://127.0.0.1:27017"
	block      = 4862998
)

var multiToken = []types.Transaction{
	{
		TransactionID:    types.NewHash("0x87012a0d870d47c3c93526c05c4a2f494054c3f4dd8584e94af7d8dd90a535f4"),
		TransactionIndex: 1,
		OrderHash:        types.NewHash("0xEEAD6DBFC7340A56CAEDC044696A168870549A6A7F6F56961E84A54BD9970B8A"),
		BlockNumber:      4862998,
		Timestamp:        types.NewInt(1515233752),
		Taker:            types.HexToAddress("0x997919a608788621dd48b3896f78dcda682fe91d"),
		Maker:            types.HexToAddress("0x9f612fcb422d1971c1be7416c37e3ebc77c0de19"),
		Make:             types.Trade{Token: types.HexToAddress("0x0000000000000000000000000000000000000000"), Amount: types.NewInt(300)},
		Take:             types.Trade{Token: types.HexToAddress("09dfd26114cd6EE289AccF82350c8d8487fedB8A0C"), Amount: types.NewInt(3000)},
	},
	{
		TransactionID:    types.NewHash("0x87012a0d870d47c3c93526c05c4a2f494054c3f4dd8584e94af7d8dd90a535f3"),
		TransactionIndex: 2,
		OrderHash:        types.NewHash("0xEEAD6DBFC7340A56CAEDC044696A168870549A6A7F6F56961E84A54BD9970B8A"),
		BlockNumber:      4862998,
		Timestamp:        types.NewInt(1515233752),
		Taker:            types.HexToAddress("0x997919a608788621dd48b3896f78dcda682fe91d"),
		Maker:            types.HexToAddress("0x9f612fcb422d1971c1be7416c37e3ebc77c0de19"),
		Make:             types.Trade{Token: types.HexToAddress("0x0000000000000000000000000000000000000000"), Amount: types.NewInt(300)},
		Take:             types.Trade{Token: types.HexToAddress("0xd26114cd6EE289AccF82350c8d8487fedB8A0C07"), Amount: types.NewInt(3000)},
	},
	{
		TransactionID:    types.NewHash("0x87012a0d870d47c3c93526c05c4a2f494054c3f4dd8584e94af7d8dd90a535f2"),
		TransactionIndex: 3,
		OrderHash:        types.NewHash("0xEEAD6DBFC7340A56CAEDC044696A168870549A6A7F6F56961E84A54BD9970B8A"),
		BlockNumber:      4862998,
		Timestamp:        types.NewInt(1515233752),
		Taker:            types.HexToAddress("0x997919a608788621dd48b3896f78dcda682fe91d"),
		Maker:            types.HexToAddress("0x9f612fcb422d1971c1be7416c37e3ebc77c0de19"),
		Make:             types.Trade{Token: types.HexToAddress("0x997919a608788621dd48b3896f78dcda682fe91d"), Amount: types.NewInt(300)},
		Take:             types.Trade{Token: types.HexToAddress("0x0000000000000000000000000000000000000000"), Amount: types.NewInt(3000)},
	},
	{
		TransactionID:    types.NewHash("0x87012a0d870d47c3c93526c05c4a2f494054c3f4dd8584e94af7d8dd90a535f8"),
		TransactionIndex: 5,
		OrderHash:        types.NewHash("0xEEAD6DBFC7340A56CAEDC044696A168870549A6A7F6F56961E84A54BD9970B8A"),
		BlockNumber:      4862998,
		Timestamp:        types.NewInt(1515233752),
		Taker:            types.HexToAddress("0x997919a608788621dd48b3896f78dcda682fe91d"),
		Maker:            types.HexToAddress("0x9f612fcb422d1971c1be7416c37e3ebc77c0de19"),
		Make:             types.Trade{Token: types.HexToAddress("0x0000000000000000000000000000000000000000"), Amount: types.NewInt(300)},
		Take:             types.Trade{Token: types.HexToAddress("0xd26114cd6EE289AccF82350c8d8487fedB8A0C07"), Amount: types.NewInt(3000)},
	},
}

func TestHistoryAggregation_AggregateTransactions(t *testing.T) {
	repository := repositories.NewMockCacheTokensRepository()
	repository.AddToken(types.HexToAddress("0xd26114cd6EE289AccF82350c8d8487fedB8A0C07"), 15)
	repository.AddToken(types.HexToAddress("0x997919a608788621dd48b3896f78dcda682fe91d"), 12)
	repository.AddToken(types.HexToAddress("09dfd26114cd6EE289AccF82350c8d8487fedB8A0C"), 11)

	mgoConnection, err := NewHistoryAggregation(connection, repository)
	if err != nil {
		t.Errorf("could not establish new connection")
	}
	mgoConnection.session.Clone()
	defer mgoConnection.session.Close()

	insertData(multiToken)
	ticks, err := mgoConnection.AggregateTransactions(block)
	if ticks == nil {
		t.Errorf("could not aggregate transactions")
	}

	if err != nil {
		t.Error(err)
	}

	// @todo tests need to be fixed
}

func insertData(transactions []types.Transaction) error {
	mgoConnection, err := NewHistoryAggregation(connection, nil)
	if err != nil {
		return fmt.Errorf("could not establish new connection")
	}
	mgoConnection.session.Clone()
	defer mgoConnection.session.Close()
	c := mgoConnection.session.DB(DBName).C(FileName)

	for _, t := range transactions {
		err := c.Insert(t)
		if err != nil {
			return fmt.Errorf("could not insert transaction")
		}
	}

	return nil
}
