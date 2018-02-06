package types

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"gopkg.in/mgo.v2/bson"
)

type Bytes struct {
	hexutil.Bytes
}

func NewBytes(bytes string) (*Bytes, error) {

	b := &Bytes{}
	err := b.UnmarshalText([]byte(bytes))
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (b Bytes) GetBSON() (interface{}, error) {
	return b.String(), nil
}

func (b *Bytes) SetBSON(raw bson.Raw) error {
	var val string
	err := raw.Unmarshal(&val)
	if err != nil {
		return err
	}

	return b.UnmarshalText(common.FromHex(val))
}
