package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/DexyProject/dexy-go/consumers"
	"github.com/DexyProject/dexy-go/contracts"
	"github.com/DexyProject/dexy-go/history"
	"github.com/DexyProject/dexy-go/log"
	"github.com/DexyProject/dexy-go/orderbook"
	"github.com/DexyProject/dexy-go/types"
	"github.com/DexyProject/dexy-go/watchers"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"
)

func main() {

	defer deferOnPanic()

	ethNode := flag.String("ethnode", "", "ethereum node address")
	mongo := flag.String("mongo", "", "mongodb connection string")
	addr := flag.String("addr", "", "contracts address")

	flag.Parse()

	if flag.NFlag() != 3 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	hist, err := history.NewMongoHistory(*mongo)
	if err != nil {
	}

	conn, err := ethclient.Dial(*ethNode)
	if err != nil {
		log.Fatal("", zap.Error(err))
	}

	ob, err := orderbook.NewMongoOrderBook(*mongo)
	if err != nil {
		log.Fatal("orderbook error", zap.Error(err))
	}

	ex, err := contracts.NewExchange(types.HexToAddress(*addr).Address, conn)

	channel := make(chan *consumers.TradedMessage)

	tc := consumers.NewTradedConsumer(ex, conn, channel)
	tf := watchers.NewTradeWatcher(hist, ex, ob, channel)

	err = tc.StartConsuming()
	if err != nil {
		log.Fatal("", zap.Error(err))
	}

	tf.Watch()
}

func deferOnPanic() {
	if err := recover(); err != nil {
		fmt.Println(err)
	}
}
