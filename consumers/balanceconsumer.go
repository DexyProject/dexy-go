package consumers

import (
	"github.com/DexyProject/dexy-go/contracts"
	"log"
	"fmt"
)

type BalanceConsumer struct {
	vault *contracts.Vault
}

func (bc *BalanceConsumer) StartConsuming() error {

	deposit, err := bc.vault.WatchDeposited(nil, nil, nil) // @todo
	if err != nil {
		return fmt.Errorf("failed to watch deposited: %s", err.Error())
	}

	withdraw, err := bc.vault.WatchWithdrawn(nil, nil, nil) // @todo
	if err != nil {
		return fmt.Errorf("failed to watch withdrawn: %s", err.Error())
	}

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
