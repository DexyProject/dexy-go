package types

type Market struct {
	Token  Address `json:"token,omitempty" bson:"token,omitempty"`
	Bid    float64 `json:"bid,omitempty" bson:"bid,omitempty"`
	Ask    float64 `json:"ask,omitempty" bson:"ask,omitempty"`
	Volume Int     `json:"volume,omitempty" bson:"volume,omitempty"`
	Last   float64 `json:"last,omitempty" bson:"last,omitempty"`
}

type Price struct {
	Base  string
	Quote string
}

type Prices map[Address]Price
