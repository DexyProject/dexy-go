package types

type Tick struct {
	Block     int64     `json:"block" bson:"block"`
	OpenTime  int64     `json:"opentime" bson:"opentime"`
	CloseTime int64     `json:"closetime" bson:"closetime"`
	Volume    Int       `json:"volume" bson:"volume"`
	Open      float64   `json:"open" bson:"open"`
	Close     float64   `json:"close" bson:"close"`
	High      float64   `json:"high" bson:"high"`
	Low       float64   `json:"low" bson:"low"`
}
