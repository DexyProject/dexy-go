package consumers

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/rpc"
)

type Block struct {
	Number *big.Int
}

type BlockConsumer struct {
	client rpc.Client

	channel chan Block

	sub *rpc.ClientSubscription
}

func NewBlockConsumer(client rpc.Client, channel chan Block) BlockConsumer {
	return BlockConsumer{client: client, channel: channel}
}

func (bc *BlockConsumer) StartConsuming() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	sub, err := bc.client.EthSubscribe(ctx, bc.channel, "newBlocks")
	if err != nil {
		return fmt.Errorf("subscribe error: %s", err.Error())
	}

	bc.sub = sub

	var lastBlock Block
	if err := bc.client.CallContext(ctx, &lastBlock, "eth_getBlockByNumber", "latest"); err != nil {
		fmt.Println("can't get latest block:", err)
		return fmt.Errorf("can't get latest block: %s", err.Error())
	}

	bc.channel <- lastBlock
	return fmt.Errorf("connection lost: %s", (<-bc.sub.Err()).Error())
}

func (bc *BlockConsumer) StopConsuming() {
	bc.sub.Unsubscribe()
}
