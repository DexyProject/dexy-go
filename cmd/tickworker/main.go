package main

import (
	"fmt"
	"log"
	"time"

	"github.com/DexyProject/dexy-go/consumers"
	"github.com/DexyProject/dexy-go/history"
	"github.com/DexyProject/dexy-go/ticks"
	"github.com/ethereum/go-ethereum/core/types"
)

func main() {

	defer deferOnPanic()

	channel := make(chan *types.Header)

	bc := consumers.NewBlockConsumer(nil, channel)

	aggregation, err := history.NewHistoryAggregation()
	if err != nil {
		log.Fatalf("failed to create aggregation: %s", err.Error())
	}

	tickdb, err := ticks.NewMongoTicks()
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
