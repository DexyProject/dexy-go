package main

import (
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

	hist, err := history.NewMongoHistory(os.Args[1])
	if err != nil {
	}

	conn, err := ethclient.Dial(os.Args[2])
	if err != nil {
		log.Fatal(err.Error())
	}

	ob, err := orderbook.NewMongoOrderBook(os.Args[1])
	if err != nil {
		log.Fatalf("Orderbook error: %v", err.Error())
	}

	ex, err := exchange.NewExchangeInterface(types.HexToAddress(os.Args[3]).Address, conn)

	tf := watchers.TradeWatcher{
		History:   hist,
		Exchange:  *ex,
		Orderbook: ob,
		Ethereum:  conn,
	}

	err = tf.Watch()
	if err != nil {
		log.Fatal(err.Error())
	}
}

func deferOnPanic() {
	if err := recover(); err != nil {
		fmt.Println(err)
	}
}
