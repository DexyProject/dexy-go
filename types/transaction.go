package types

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/DexyProject/dexy-go/types"
)

type Transaction struct {
	TransactionID string         `json:"tx" bson:"tx"`
	OrderHash     string         `json:"hash" bson:"hash"`
	BlockNumber   int            `json:"block" bson:"block"`
	Timestamp     string         `json:"timestamp" bson:"timestamp"`
	Taker         Address        `json:"taker" bson:"taker"`
	Maker         Address        `json:"maker" bson:"maker"`
	Give          Trade          `json:"give" bson:"give"`
	Get           Trade          `json:"get" bson:"get"`
}
