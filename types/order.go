package types

import (
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto/sha3"
)

var ORDER_HASH_SCHEME = NewHash("0xb9caf644225739cd2bda9073346357ae4a0c3d71809876978bd81cc702b7fdc7")

type Trade struct {
	Token  Address `json:"token" bson:"token"`
	Amount Int     `json:"amount" bson:"amount"`
}

type Orders struct {
	Asks []Order `json:"asks"`
	Bids []Order `json:"bids"`
}

type Order struct {
	Hash      Hash      `json:"hash,omitempty" bson:"_id"`
	Price     float64   `json:"-" bson:"price"`
	Make      Trade     `json:"make" bson:"make"`
	Take      Trade     `json:"take" bson:"take"`
	Expires   Timestamp `json:"expires" bson:"expires"`
	Nonce     int64     `json:"nonce" bson:"nonce"`
	User      Address   `json:"user" bson:"user"`
	Exchange  Address   `json:"exchange" bson:"exchange"`
	Signature EC        `json:"signature" bson:"signature"`
	Filled    Int       `json:"filled,omitempty" bson:"filled"`
}

func (o *Order) OrderHash() Hash {
	if o.Hash.String() == (Hash{}).String() {
		o.generateHash()
	}

	return o.Hash
}

func (o *Order) generateHash() {

	hash := sha3.NewKeccak256()
	hash.Write(o.Take.Token.Address[:])
	hash.Write(o.Take.Amount.U256()[:])
	hash.Write(o.Make.Token.Address[:])
	hash.Write(o.Make.Amount.U256()[:])
	hash.Write(NewInt(o.Expires.Unix()).U256()[:])
	hash.Write(NewInt(o.Nonce).U256()[:])
	hash.Write(o.User.Address[:])
	hash.Write(o.Exchange.Address[:])

	sha := sha3.NewKeccak256()
	sha.Write(ORDER_HASH_SCHEME.Hash[:])
	sha.Write(hash.Sum(nil))

	o.Hash.SetBytes(sha.Sum(nil))
}

func (o *Order) Validate() error {
	if !common.IsHexAddress(o.Take.Token.String()) || !common.IsHexAddress(o.Make.Token.String()) {
		return fmt.Errorf("address is not valid hex")
	}

	if strings.ToLower(o.Make.Token.String()) == strings.ToLower(o.Take.Token.String()) {
		return fmt.Errorf("token addresses are identical")
	}

	if !o.Expires.After(time.Now()) {
		return fmt.Errorf("invalid expires time: %v", o.Expires)
	}

	zero := new(big.Int).SetInt64(0)
	if o.Take.Amount.Cmp(zero) == 0 || o.Make.Amount.Cmp(zero) == 0 {
		return fmt.Errorf("amounts are not allowed to be 0")
	}

	return nil
}
