package types

type Transaction struct {
	TransactionID Bytes   `json:"tx" bson:"tx"`
	OrderHash     Hash    `json:"hash" bson:"hash"`
	BlockNumber   uint64   `json:"block" bson:"block"`
	Timestamp     int64   `json:"timestamp" bson:"timestamp"`
	Taker         Address `json:"taker" bson:"taker"`
	Maker         Address `json:"maker" bson:"maker"`
	Give          Trade   `json:"give" bson:"give"`
	Get           Trade   `json:"get" bson:"get"`
}
