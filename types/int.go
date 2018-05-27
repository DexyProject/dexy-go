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

func (x Int) Add(y Int) (Int) {
	return Int{*x.Int.Add(&x.Int, &y.Int)}
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

	return x.setFromString(val)
}

func (x Int) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}

func (x *Int) UnmarshalJSON(input []byte) error {
	var val string
	err := json.Unmarshal(input, &val)
	if err != nil {
		return err
	}

	return x.setFromString(val)
}

func (x Int) U256() []byte {
	return abi.U256(&x.Int)
}

func (x *Int) setFromString(val string) error {
	num, ok := new(big.Int).SetString(val, 10)
	if !ok {
		return fmt.Errorf("could not set string")
	}

	x.Int = *num

	return nil
}
