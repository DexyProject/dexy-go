package types

import (
	"gopkg.in/mgo.v2/bson"
	"math/big"
	"github.com/ethereum/go-ethereum/accounts/abi"
)

type Int struct {
	big.Int
}

func NewInt(x int64) Int {
	return Int{*new(big.Int).SetInt64(x)}
}

func (x Int) GetBSON() (interface{}, error) {
	return x.Int.String(), nil
}

func (x *Int) SetBSON(raw bson.Raw) error {

	//var addr string
	//err := raw.Unmarshal(addr)
	//if err != nil {
	//	return err
	//}

	return nil
}

func (x Int) U256() []byte {
	return abi.U256(&x.Int)
}