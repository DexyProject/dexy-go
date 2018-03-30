package main

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/DexyProject/dexy-go/consumers"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
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

	conn, err := ethclient.Dial(*ethNode)
	if err != nil {
		log.Fatal(err.Error())
	}

	bc := consumers.NewBlockConsumer(conn, channel)

	err = bc.StartConsuming()
	if err != nil {
		log.Fatalf("failed to start consuming: %s", err.Error())
	}

	for {

		head := <-channel

		// we sleep here in case transactions are still inserting, 5 seconds should probably be enough
		time.Sleep(5 * time.Second)

	}
}
