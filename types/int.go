package types

import (
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"gopkg.in/mgo.v2/bson"
)

type Int struct {
	big.Int
}

func NewInt(x int64) Int {
	return Int{*new(big.Int).SetInt64(x)}
}

func IntFromString(x string) (*Int, bool) {
	i, ok := new(Int).SetString(x, 0)
	return &Int{*i}, ok
}

func (x Int) GetBSON() (interface{}, error) {
	return x.Int.String(), nil
}

func (x *Int) SetBSON(raw bson.Raw) error {
	var val string
	err := raw.Unmarshal(&val)
	if err != nil {
		return err
	}

	num, ok := new(big.Int).SetString(val, 10)
	if !ok {
		return fmt.Errorf("could not set string")
	}

	x.Int = *num

	return nil
}

func (x Int) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}

func (x Int) U256() []byte {
	return abi.U256(&x.Int)
}
