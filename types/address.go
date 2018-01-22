package types

import (
	"github.com/ethereum/go-ethereum/common"
	"gopkg.in/mgo.v2/bson"
)

type Address struct {
	common.Address
}

func (a Address) GetBSON() (interface{}, error) {
	return a.String(), nil
}

func (a *Address) SetBSON(raw bson.Raw) error {

	var addr string
	err := raw.Unmarshal(addr)
	if err != nil {
		return err
	}

	return a.UnmarshalJSON([]byte(addr))
}