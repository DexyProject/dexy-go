package consumers

import (
	"github.com/DexyProject/dexy-go/contracts"
	"log"
)

type BalanceConsumer struct {
	vault *contracts.Vault
}

func (bc *BalanceConsumer) StartConsuming() error {

	sub, err := bc.vault.WatchDeposited(nil, nil, nil) // @todo
	sub, err := bc.vault.WatchWithdrawn(nil, nil, nil) // @todo

	return nil
}

func (bc *BalanceConsumer) StopConsuming() {
	panic("implement me")
}

func (bc *BalanceConsumer) consume(deposited chan *contracts.VaultDeposited, withdrawn chan *contracts.VaultWithdrawn) {
	for {
		select {
		case deposit := <-deposited:
			log.Print(deposit)
		case withdraw := <-withdrawn:
			log.Print(withdraw)
		}
	}
}
