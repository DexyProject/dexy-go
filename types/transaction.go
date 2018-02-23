package types

import "github.com/DexyProject/dexy-go/exchange"

type Transaction struct {
	TransactionID    Bytes   `json:"tx" bson:"_id"`
  TransactionIndex uint    `json:"index" bson:"index"`
	OrderHash        Hash    `json:"hash" bson:"hash"`
	BlockNumber      uint64  `json:"block" bson:"block"`
	Timestamp        Int     `json:"timestamp" bson:"timestamp"`
	Taker            Address `json:"taker" bson:"taker"`
	Maker            Address `json:"maker" bson:"maker"`
	Give             Trade   `json:"give" bson:"give"`
	Get              Trade   `json:"get" bson:"get"`
}

func NewTransaction(trade exchange.ExchangeInterfaceTraded, timestamp Int) Transaction {
	return Transaction{
		TransactionID:    Bytes{Bytes: trade.Raw.TxHash.Bytes()},
		TransactionIndex: trade.Raw.Index,
		OrderHash:        trade.Hash,
		BlockNumber:      trade.Raw.BlockNumber,
		Timestamp:        timestamp,
		Taker:            Address{Address: trade.Taker},
		Maker:            Address{Address: trade.Maker},
		Give: Trade{
			Token:  Address{Address: trade.TokenGive},
			Amount: Int{Int: *trade.AmountGive},
		},
		Get: Trade{
			Token:  Address{Address: trade.TokenGet},
			Amount: Int{Int: *trade.AmountGet},
		},
}}
