package types

import (
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"gopkg.in/mgo.v2/bson"
)

type Address struct {
	common.Address
}

func HexToAddress(hex string) Address {
	return Address{Address: common.HexToAddress(hex)}
}

func (a Address) GetBSON() (interface{}, error) {
	return strings.ToLower(a.String()), nil
}

func (a *Address) SetBSON(raw bson.Raw) error {

	var addr string
	err := raw.Unmarshal(addr)
	if err != nil {
		return err
	}

	return a.UnmarshalJSON([]byte(addr))
}

func (a Address) IsZero() bool {
	return a.String() == "0x0000000000000000000000000000000000000000"
}
