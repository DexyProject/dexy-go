package types

import (
	"fmt"
	"math"
	"math/big"
	"strings"
	"time"

	"github.com/DexyProject/dexy-go/contracts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto/sha3"
)

var ORDER_HASH_SCHEME = NewHash("0xa8da5e6ea8c46a0516b3a2e3b010f264e8334214f4b37ff5f2bc8a2dd3f32be1")

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
	Price     string    `json:"-" bson:"price"`
	Give      Trade     `json:"give" bson:"give"`
	Get       Trade     `json:"get" bson:"get"`
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
	hash.Write(o.Get.Token.Address[:])
	hash.Write(o.Get.Amount.U256()[:])
	hash.Write(o.Give.Token.Address[:])
	hash.Write(o.Give.Amount.U256()[:])
	hash.Write(NewInt(o.Expires.Unix()).U256()[:])
	hash.Write(NewInt(o.Nonce).U256()[:])
	hash.Write(o.User.Address[:])
	hash.Write(o.Exchange.Address[:])

	sha := sha3.NewKeccak256()
	sha.Write(ORDER_HASH_SCHEME[:])
	sha.Write(hash.Sum(nil))

	o.Hash.SetBytes(sha.Sum(nil))
}

func (o *Order) Validate() error {
	if !common.IsHexAddress(o.Get.Token.String()) || !common.IsHexAddress(o.Give.Token.String()) {
		return fmt.Errorf("address is not valid hex")
	}

	if strings.ToLower(o.Give.Token.String()) == strings.ToLower(o.Get.Token.String()) {
		return fmt.Errorf("token addresses are identical")
	}

	if !o.Expires.After(time.Now()) {
		return fmt.Errorf("invalid expires time: %v", o.Expires)
	}

	zero := new(big.Int).SetInt64(0)
	if o.Get.Amount.Cmp(zero) == 0 || o.Give.Amount.Cmp(zero) == 0 {
		return fmt.Errorf("amounts are not allowed to be 0")
	}

	return nil
}

func (t *Trade) CalcPrice(g Trade, base Address) (float64, error) {
	var erc20 contracts.ERC20
	decimals, _ := erc20.Decimals(nil)
	if t.Amount.Sign() <= 0 || g.Amount.Sign() <= 0 {
		return 0.0, fmt.Errorf("can not divide by zero")
	}

	tFloat := new(big.Float).SetInt(&t.Amount.Int)
	gFloat := new(big.Float).SetInt(&g.Amount.Int)

	decimalsFloat := float64(decimals)
	if t.Token == base {
		price, _ := new(big.Float).Quo(tFloat, gFloat).Float64()
		return (price / math.Pow(10.0, decimalsFloat)), nil
	}
	price, _ := new(big.Float).Quo(gFloat, tFloat).Float64()
	return (price / math.Pow(10.0, decimalsFloat)), nil

}
