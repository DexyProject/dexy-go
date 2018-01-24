package types

import (
	"encoding/json"
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
	var s string
	err := raw.Unmarshal(&s)
	if err != nil {
		return err
	}

	a.Address = common.HexToAddress(s)

	return nil
}

func (a Address) IsZero() bool {
	return a.String() == "0x0000000000000000000000000000000000000000"
}

func (a *Address) MarshalJSON() ([]byte, error) {
	return json.Marshal(strings.ToLower(a.Address.String()))
}
