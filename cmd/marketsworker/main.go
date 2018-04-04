package main

import (
	"flag"
	"os"
	"time"

	"github.com/DexyProject/dexy-go/builders"
	"github.com/DexyProject/dexy-go/consumers"
	"github.com/DexyProject/dexy-go/log"
	"github.com/DexyProject/dexy-go/markets"
	"github.com/DexyProject/dexy-go/orderbook"
	"github.com/DexyProject/dexy-go/ticks"
	"github.com/DexyProject/dexy-go/types"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"
)

func main() {

	ethNode := flag.String("ethnode", "", "ethereum node address")
	mongo := flag.String("mongo", "", "mongodb connection string")

	flag.Parse()

	if flag.NFlag() != 2 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	channel := make(chan *ethtypes.Header)

	tokens := make([]types.Address, 0)
	tokens = append(tokens, types.HexToAddress("0xbebb2325ef529e4622761498f1f796d262100768"))

	mb := builders.MarketsBuilder{}
	m := markets.MongoMarkets{}
	ob, err := orderbook.NewMongoOrderBook(*mongo)
	if err != nil {
		log.Fatal("", zap.Error(err))
	}

	t, err := ticks.NewMongoTicks(*mongo)
	if err != nil {
		log.Fatal("", zap.Error(err))
	}

	conn, err := ethclient.Dial(*ethNode)
	if err != nil {
		log.Fatal("", zap.Error(err))
	}

	bc := consumers.NewBlockConsumer(conn, channel)

	err = bc.StartConsuming()
	if err != nil {
		log.Fatal("failed to start consuming", zap.Error(err))
	}

	for {

		head := <-channel

		//	 we sleep here in case transactions are still inserting, 5 seconds should probably be enough
		time.Sleep(5 * time.Second)

		ts, err := t.FetchLatestTickForTokens(tokens)
		if err != nil {
			log.Error("", zap.Error(err))
			continue
		}

		asks, err := ob.GetLowestAsks(tokens)
		if err != nil {
			log.Error("", zap.Error(err))
			continue
		}

		bids, err := ob.GetHighestBids(tokens)
		if err != nil {
			log.Error("", zap.Error(err))
			continue
		}

		ms := mb.Build(tokens, ts, asks, bids)

		err = m.InsertMarkets(ms)
		if err != nil {
			log.Fatal("failed to insert markets", zap.Error(err))
		}

		log.Debug("inserted markets", zap.String("block", head.Number.String()))
	}
}
