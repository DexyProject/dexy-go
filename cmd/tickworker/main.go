package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/DexyProject/dexy-go/consumers"
	"github.com/DexyProject/dexy-go/gateways/history"
	"github.com/DexyProject/dexy-go/repositories"
	"github.com/DexyProject/dexy-go/gateways/ticks"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
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
		log.Fatal(err.Error())
	}

	bc := consumers.NewBlockConsumer(conn, channel)

	repo, err := repositories.NewCacheTokensRepository(*ethNode)
	if err != nil {
		log.Fatalf("failed to create tokens repository: %s", err.Error())
	}

	aggregation, err := history.NewHistoryAggregation(*mongo, repo)
	if err != nil {
		log.Fatalf("failed to create aggregation: %s", err.Error())
	}

	tickdb, err := ticks.NewMongoTicks(*mongo)
	if err != nil {
		log.Fatalf("failed to create mongo ticks: %s", err.Error())
	}

	err = bc.StartConsuming()
	if err != nil {
		log.Fatalf("failed to start consuming: %s", err.Error())
	}

	for {

		head := <-channel

		// we sleep here in case transactions are still inserting, 5 seconds should probably be enough
		time.Sleep(5 * time.Second)

		t, err := aggregation.AggregateTransactions(head.Number.Int64())
		if err != nil {
			log.Printf("failed aggregating transactions: %s", err.Error())
		}

		if len(t) == 0 {
			log.Printf("no ticks for block: %s", head.Number.String())
			continue
		}

		log.Printf("inserting ticks for block: %s", head.Number.String())

		err = tickdb.InsertTicks(t)
		if err != nil {
			log.Printf("failed to insert ticks: %s", err.Error())
		}
	}
}

func deferOnPanic() {
	if err := recover(); err != nil {
		fmt.Println(err)
	}
}
