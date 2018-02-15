package types

type Tick struct {
	//Pair      Pair      `json:"pair" bson:"pair"`
	Block     int64     `json:"block" bson:"block"`
	Volume    Int       `json:"volume" bson:"volume"`
	Open      float64   `json:"open" bson:"open"`
	Close     float64   `json:"close" bson:"close"`
	High      float64   `json:"high" bson:"high"`
	Low       float64   `json:"low" bson:"low"`
}

type Pair struct {
	Base      Address   `json:"base" bson:"base"`
	Quote     Address   `json:"quote" bson:"quote"`

}
