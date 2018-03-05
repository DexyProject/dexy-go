package consumers

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Block struct {
	Number *big.Int
}

type BlockConsumer struct {
	client *ethclient.Client

	channel chan *types.Header

	sub ethereum.Subscription
}

func NewBlockConsumer(client *ethclient.Client, channel chan *types.Header) BlockConsumer {
	return BlockConsumer{client: client, channel: channel}
}

func (bc *BlockConsumer) StartConsuming() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	sub, err := bc.client.SubscribeNewHead(ctx, bc.channel)
	if err != nil {
		return fmt.Errorf("subscribe error: %s", err.Error())
	}

	bc.sub = sub

	return nil
}

func (bc *BlockConsumer) StopConsuming() {
	bc.sub.Unsubscribe()
}
