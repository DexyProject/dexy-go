package main

import (
	"flag"
	"os"
	"time"

	"github.com/DexyProject/dexy-go/builders"
	"github.com/DexyProject/dexy-go/consumers"
	"github.com/DexyProject/dexy-go/log"
	"github.com/DexyProject/dexy-go/markets"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"
)

func main() {

	ethNode := flag.String("ethnode", "", "ethereum node address")
	//mongo := flag.String("mongo", "", "mongodb connection string")

	flag.Parse()

	if flag.NFlag() != 2 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	channel := make(chan *types.Header)

	mb := builders.MarketsBuilder{}
	m := markets.MongoMarkets{}

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

		// @todo

		err := m.InsertMarkets()
		if err != nil {
			// @todo log
		}

		// we sleep here in case transactions are still inserting, 5 seconds should probably be enough
		time.Sleep(5 * time.Second)

	}
}
