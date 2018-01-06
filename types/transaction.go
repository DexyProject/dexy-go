package types

type Transaction struct {
	TransactionID string `json:"tx"`
	OrderHash     string `json:"hash"`
	BlockNumber   int    `json:"hash"`
	Timestamp     string `json:"timestamp"`
}
