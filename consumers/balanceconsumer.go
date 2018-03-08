package consumers

import (
	"fmt"
	"log"

	"github.com/DexyProject/dexy-go/contracts"
	"github.com/ethereum/go-ethereum/event"
)

type BalanceConsumer struct {
	vault *contracts.Vault

	stop chan struct{}

	withdraw, deposit event.Subscription
}

func (bc *BalanceConsumer) StartConsuming() error {

	deposited := make(chan *contracts.VaultDeposited)
	deposit, err := bc.vault.WatchDeposited(nil, deposited, nil)
	if err != nil {
		return fmt.Errorf("failed to watch deposited: %s", err.Error())
	}

	withdrawn := make(chan *contracts.VaultWithdrawn)
	withdraw, err := bc.vault.WatchWithdrawn(nil, withdrawn, nil)
	if err != nil {
		return fmt.Errorf("failed to watch withdrawn: %s", err.Error())
	}

	bc.withdraw = withdraw
	bc.deposit = deposit

	go bc.consume(deposited, withdrawn)

	return nil
}

func (bc *BalanceConsumer) StopConsuming() {
	bc.withdraw.Unsubscribe()
	bc.deposit.Unsubscribe()
	close(bc.stop)
}

func (bc *BalanceConsumer) consume(deposited chan *contracts.VaultDeposited, withdrawn chan *contracts.VaultWithdrawn) {
	for {
		select {
		case deposit := <-deposited:
			log.Print(deposit)
		case withdraw := <-withdrawn:
			log.Print(withdraw)
		case <-bc.stop:
			return
		}
	}
}
