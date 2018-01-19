package types

import (
	"github.com/ethereum/go-ethereum/common"
)

type Tick struct {
	Tx        string         `json:"tx" bson:"tx"`
	Hash      string         `json:"hash,omitempty" bson:"hash"`
	Block     string         `json:"block" bson:"block"`
	OpenTime  int64          `json:"opentime" bson:"opentime"`
	CloseTime int64          `json:"closetime" bson:"closetime"`
	Taker     common.Address `json:"taker" bson:"taker"`
	Maker     common.Address `json:"maker" bson:"maker"`
	Give      Trade          `json:"give" bson:"give,inline"`
	Get       Trade          `json:"get" bson:"get,inline"`
}










