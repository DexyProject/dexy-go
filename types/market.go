package types

type PairAmount struct {
	Quote string `json:"quote"`
	Base  string `json:"base"`
}

type Market struct {
	Token Address    `json:"token"`
	Bid   PairAmount `json:"bid"`
	Ask   PairAmount `json:"ask"`
}
