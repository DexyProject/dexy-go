package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/DexyProject/dexy-go/contracts"
	"github.com/DexyProject/dexy-go/log"
	"github.com/DexyProject/dexy-go/watchers"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"
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

	v, err := setupVault(*ethNode, common.HexToAddress(*addr))
	if err != nil {
		log.Fatal("", zap.Error(err))
	}

	bw := watchers.NewBalanceWatcher()

}

func setupVault(ethereum string, addr common.Address) (*contracts.Vault, error) {
	conn, err := ethclient.Dial(ethereum)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the ethereum client: %v", err)
	}

	return contracts.NewVault(addr, conn)
}

func deferOnPanic() {
	if err := recover(); err != nil {
		fmt.Println(err)
	}
}
