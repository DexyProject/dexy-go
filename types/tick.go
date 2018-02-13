package types

type Tick struct {
	Block     int64 `json:"block" bson:"block"`
	OpenTime  int64 `json:"opentime" bson:"opentime"`
	CloseTime int64 `json:"closetime" bson:"closetime"`
	Volume    Int   `json:"volume" bson:"volume"`
	Open      Int   `json:"open" bson:"open"`
	Close     Int   `json:"close" bson:"close"`
	High      Int   `json:"high" bson:"high"`
	Low       Int   `json:"low" bson:"low"`
}
