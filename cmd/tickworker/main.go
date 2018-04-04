package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/DexyProject/dexy-go/consumers"
	"github.com/DexyProject/dexy-go/history"
	"github.com/DexyProject/dexy-go/log"
	"github.com/DexyProject/dexy-go/repositories"
	"github.com/DexyProject/dexy-go/ticks"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"
)

func main() {

	defer deferOnPanic()

	ethNode := flag.String("ethnode", "", "ethereum node address")
	mongo := flag.String("mongo", "", "mongodb connection string")

	flag.Parse()

	if flag.NFlag() != 2 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	channel := make(chan *types.Header)

	conn, err := ethclient.Dial(*ethNode)
	if err != nil {
		log.Fatal("", zap.Error(err))
	}

	bc := consumers.NewBlockConsumer(conn, channel)

	repo, err := repositories.NewCacheTokensRepository(*ethNode)
	if err != nil {
		log.Fatal("failed to create tokens repository", zap.Error(err))
	}

	aggregation, err := history.NewHistoryAggregation(*mongo, repo)
	if err != nil {
		log.Fatal("failed to create aggregation", zap.Error(err))
	}

	tickdb, err := ticks.NewMongoTicks(*mongo)
	if err != nil {
		log.Fatal("failed to create mongo ticks", zap.Error(err))
	}

	err = bc.StartConsuming()
	if err != nil {
		log.Fatal("failed to start consuming", zap.Error(err))
	}

	for {

		head := <-channel

		// we sleep here in case transactions are still inserting, 5 seconds should probably be enough
		time.Sleep(5 * time.Second)

		t, err := aggregation.AggregateTransactions(head.Number.Int64())
		if err != nil {
			log.Error("failed aggregating transactions", zap.Error(err))
		}

		if len(t) == 0 {
			log.Debug("no ticks for block", zap.String("block", head.Number.String()))
			continue
		}

		log.Debug("inserting ticks for block", zap.String("block", head.Number.String()))

		err = tickdb.InsertTicks(t)
		if err != nil {
			log.Error("failed to insert ticks", zap.Error(err))
		}
	}
}

func deferOnPanic() {
	if err := recover(); err != nil {
		fmt.Println(err)
	}
}
