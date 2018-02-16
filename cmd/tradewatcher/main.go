package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/DexyProject/dexy-go/exchange"
	"github.com/DexyProject/dexy-go/history"
	"github.com/DexyProject/dexy-go/orderbook"
	"github.com/DexyProject/dexy-go/types"
	"github.com/DexyProject/dexy-go/watchers"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {

	defer deferOnPanic()

	ethNode := flag.String("ethnode", "", "ethereum node address")
	mongo := flag.String("mongo", "", "mongodb connection string")
	addr := flag.String("addr", "", "exchange address")

	flag.Parse()

	if *ethNode == "" || *mongo == "" || *addr == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	hist, err := history.NewMongoHistory(*mongo)
	if err != nil {
	}

	conn, err := ethclient.Dial(*ethNode)
	if err != nil {
		log.Fatal(err.Error())
	}

	ob, err := orderbook.NewMongoOrderBook(*mongo)
	if err != nil {
		log.Fatalf("Orderbook error: %v", err.Error())
	}

	ex, err := exchange.NewExchangeInterface(types.HexToAddress(*addr).Address, conn)

	channel := make(chan types.Transaction)

	tf := watchers.NewTradeWatcher(hist, *ex, ob, channel)

	tf.Watch()
}

func deferOnPanic() {
	if err := recover(); err != nil {
		fmt.Println(err)
	}
}
