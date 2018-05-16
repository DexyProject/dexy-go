package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/DexyProject/dexy-go/contracts"
	"github.com/DexyProject/dexy-go/log"
	"github.com/DexyProject/dexy-go/orderbook"
	"github.com/DexyProject/dexy-go/types"
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
		log.Fatal("vault error", zap.Error(err))
	}

	ob, err := orderbook.NewMongoOrderBook(*mongo)
	if err != nil {
		log.Fatal("orderbook error", zap.Error(err))
	}

	c := make(chan watchers.Balance)
	bw := watchers.NewBalanceWatcher(ob, v, c)

	go bw.Watch()

	ds := make(chan *contracts.VaultDeposited)
	_, err = v.WatchDeposited(nil, ds, nil)
	if err != nil {
		log.Fatal("deposit watch error", zap.Error(err))
	}

	ws := make(chan *contracts.VaultWithdrawn)
	_, err = v.WatchWithdrawn(nil, ws, nil)
	if err != nil {
		log.Fatal("withdraw watch error", zap.Error(err))
	}

	go handleWithdraw(c, ws)
	handleDeposit(c, ds)
}

func handleWithdraw(b chan<- watchers.Balance, c <-chan *contracts.VaultWithdrawn) {
	for {
		data := <-c
		b <- watchers.Balance{
			User:  types.Address{data.User},
			Token: types.Address{data.Token},
		}
	}
}

func handleDeposit(b chan<- watchers.Balance, c <-chan *contracts.VaultDeposited) {
	for {
		data := <-c
		b <- watchers.Balance{
			User:  types.Address{data.User},
			Token: types.Address{data.Token},
		}
	}
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
