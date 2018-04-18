package types

import "github.com/DexyProject/dexy-go/contracts"

type Transaction struct {
	TransactionID    Hash      `json:"tx" bson:"_id"`
	TransactionIndex uint      `json:"index" bson:"index"`
	OrderHash        Hash      `json:"hash" bson:"hash"`
	BlockNumber      uint64    `json:"block" bson:"block"`
	Timestamp        Timestamp `json:"timestamp" bson:"timestamp"`
	Taker            Address   `json:"taker" bson:"taker"`
	Maker            Address   `json:"maker" bson:"maker"`
	Make             Trade     `json:"make" bson:"make"`
	Take             Trade     `json:"take" bson:"take"`
}

func NewTransaction(trade contracts.ExchangeTraded, timestamp Int) Transaction {
	return Transaction{
		TransactionID:    Hash{trade.Raw.TxHash},
		TransactionIndex: trade.Raw.Index,
		OrderHash:        Hash{trade.Hash},
		BlockNumber:      trade.Raw.BlockNumber,
		Timestamp:        NewTimestampFromInt(timestamp),
		Taker:            Address{Address: trade.Taker},
		Maker:            Address{Address: trade.Maker},
		Make: Trade{
			Token:  Address{Address: trade.MakerToken},
			Amount: Int{Int: *trade.MakerTokenAmount},
		},
		Take: Trade{
			Token:  Address{Address: trade.TakerToken},
			Amount: Int{Int: *trade.TakerTokenAmount},
		},
	}
}
