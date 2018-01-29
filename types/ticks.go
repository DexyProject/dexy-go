package types

import "github.com/DexyProject/dexy-go/types"

type Ticks struct {
	Block                int64          `json:"block" bson:"block"`
	OpenTime             int64          `json:"opentime" bson:"opentime"`
	CloseTime            int64          `json:"closetime" bson:"closetime"`
//	Volume               int64          `json:"volume" bson:"volume"`
	Open                 types.Int      `json:"open" bson:"open"`
	Close                types.Int      `json:"close" bson:"close"`
	High                 types.Int      `json:"high" bson:"high"`
	Low                  types.Int      `json:"low" bson:"low"`
}
