package types

import (
	"gopkg.in/mgo.v2/bson"
	"encoding/json"
)

type Transaction struct {
	TransactionID Bytes   `json:"tx" bson:"tx"`
	OrderHash     Hash    `json:"hash" bson:"hash"`
	BlockNumber   int64   `json:"block" bson:"block"`
	Timestamp     int64   `json:"timestamp" bson:"timestamp"`
	Taker         Address `json:"taker" bson:"taker"`
	Maker         Address `json:"maker" bson:"maker"`
	Give          Trade   `json:"give" bson:"give"`
	Get           Trade   `json:"get" bson:"get"`
}

func (t *Transaction) MarshalBson() ([]byte, error) {
	return json.Marshal(t)
}

func (t *Transaction) unMarshalBson(input []byte) error {
	return bson.Unmarshal(input, t)
}
