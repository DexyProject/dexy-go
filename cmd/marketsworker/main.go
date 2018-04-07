package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"os"
	//"time"

	//"github.com/DexyProject/dexy-go/builders"
	"github.com/DexyProject/dexy-go/consumers"
	"github.com/DexyProject/dexy-go/log"
	//"github.com/DexyProject/dexy-go/markets"
	//"github.com/DexyProject/dexy-go/orderbook"
	//"github.com/DexyProject/dexy-go/repositories"
	"github.com/DexyProject/dexy-go/ticks"
	"github.com/DexyProject/dexy-go/types"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"
	"fmt"
)

func main() {

	ethNode := flag.String("ethnode", "", "ethereum node address")
	mongo := flag.String("mongo", "", "mongodb connection string")
	path := flag.String("path", "", "path to tokens file")

	flag.Parse()

	if flag.NFlag() != 3 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	channel := make(chan *ethtypes.Header)

	tokens := loadTokens(*path)

	conn, err := ethclient.Dial(*ethNode)
	if err != nil {
		log.Fatal("", zap.Error(err))
	}

	//tr, err := repositories.NewCacheTokensRepository(*ethNode)
	//if err != nil {
	//	log.Fatal("", zap.Error(err))
	//}
	//
	//mb := builders.NewMarketsBuilder(tr)
	//m, err := markets.NewMongoMarkets(*mongo)
	//if err != nil {
	//	log.Fatal("", zap.Error(err))
	//}
	//
	//ob, err := orderbook.NewMongoOrderBook(*mongo)
	//if err != nil {
	//	log.Fatal("", zap.Error(err))
	//}

	t, err := ticks.NewMongoTicks(*mongo)
	if err != nil {
		log.Fatal("", zap.Error(err))
	}

	bc := consumers.NewBlockConsumer(conn, channel)

	err = bc.StartConsuming()
	if err != nil {
		log.Fatal("failed to start consuming", zap.Error(err))
	}

	ts, err := t.FetchLatestCloseForTokens(tokens)
	fmt.Printf("%+v", ts)
	//for {
	//
	//	head := <-channel
	//
	//	//	 we sleep here in case transactions are still inserting, 5 seconds should probably be enough
	//	time.Sleep(5 * time.Second)
	//
	//	ts, err := t.FetchLatestTickForTokens(tokens)
	//	if err != nil {
	//		log.Error("", zap.Error(err))
	//		continue
	//	}
	//
	//	asks, err := ob.GetLowestAsks(tokens)
	//	if err != nil {
	//		log.Error("", zap.Error(err))
	//		continue
	//	}
	//
	//	bids, err := ob.GetHighestBids(tokens)
	//	if err != nil {
	//		log.Error("", zap.Error(err))
	//		continue
	//	}
	//
	//	ms := mb.Build(tokens, ts, asks, bids)
	//
	//	err = m.InsertMarkets(ms)
	//	if err != nil {
	//		log.Fatal("failed to insert markets", zap.Error(err))
	//	}
	//
	//	log.Debug("inserted markets", zap.String("block", head.Number.String()))
	//}
}

func loadTokens(path string) []types.Address {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("", zap.Error(err))
	}

	tokens := make([]types.Address, 0)

	err = json.Unmarshal(file, &tokens)
	if err != nil {
		log.Fatal("", zap.Error(err))
	}

	if len(tokens) == 0 {
		log.Fatal("no tokens in file", zap.String("file", path))
	}

	return tokens
}
