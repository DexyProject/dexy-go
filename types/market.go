package types

type Market struct {
	Token Address `json:"token"`
	Bid   string  `json:"bid"`
	Ask   string  `json:"ask"`
}
