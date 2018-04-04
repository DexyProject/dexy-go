package types

type Market struct {
	Token  Address `json:"token" bson:"token"`
	Bid    float64 `json:"bid" bson:"bid"`
	Ask    float64 `json:"ask" bson:"ask"`
	Volume float64 `json:"volume" bson:"volume"`
	Last   float64 `json:"last" bson:"last"`
}

type Price struct {
	Base  string
	Quote string
}

type Prices map[Address]Price