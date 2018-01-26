package types

import (
	"encoding/json"
	"reflect"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"gopkg.in/mgo.v2/bson"
)

type Hash [32]byte

func NewHash(hash string) Hash {
	h := Hash{}
	h.SetBytes(common.FromHex(hash))
	return h
}

func (h Hash) String() string {
	return common.ToHex(h[:])
}

func (h Hash) GetBSON() (interface{}, error) {
	return h.String(), nil
}

func (h *Hash) SetBSON(raw bson.Raw) error {
	var val string
	err := raw.Unmarshal(&val)
	if err != nil {
		return err
	}

	h.SetBytes(common.FromHex(val))

	return nil
}

func (h Hash) MarshalJSON() ([]byte, error) {
	return json.Marshal(h.String())
}

func (h *Hash) UnmarshalJSON(input []byte) error {
	return hexutil.UnmarshalFixedJSON(reflect.TypeOf(Hash{}), input, h[:])
}

func (h *Hash) SetBytes(b []byte) {
	copy(h[:], b)
}
