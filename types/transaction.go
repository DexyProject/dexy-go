package types

import "github.com/ethereum/go-ethereum/common"

type Transaction struct {
	TransactionID string         `json:"tx"`
	OrderHash     string         `json:"hash"`
	BlockNumber   int            `json:"block"`
	Timestamp     string         `json:"timestamp"`
	Taker         common.Address `json:"taker"`
	Maker         common.Address `json:"maker"`
	Give          Trade          `json:"give"`
	Get           Trade          `json:"get"`
}
