package types

type Tick struct {
	Pair      Pair      `json:"pair" bson:"pair"`
	Block     uint64    `json:"block" bson:"block"`
	Volume    Int       `json:"volume" bson:"volume"`
	Open      float64   `json:"open" bson:"open"`
	Close     float64   `json:"close" bson:"close"`
	High      float64   `json:"high" bson:"high"`
	Low       float64   `json:"low" bson:"low"`
	Timestamp Timestamp `json:"timestamp" bson:"timestamp"`
}

type Pair struct {
	Quote Address `json:"quote" bson:"quote"`
	Base  Address `json:"base" bson:"base"`
}
