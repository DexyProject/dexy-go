package types

type Transaction struct {
	TransactionID string  `json:"tx" bson:"tx"`
	OrderHash     string  `json:"hash" bson:"hash"`
	BlockNumber   int64   `json:"block" bson:"block"`
	Timestamp     int64   `json:"timestamp" bson:"timestamp"`
	Taker         Address `json:"taker" bson:"taker"`
	Maker         Address `json:"maker" bson:"maker"`
	Give          Trade   `json:"give" bson:"give"`
	Get           Trade   `json:"get" bson:"get"`
}
