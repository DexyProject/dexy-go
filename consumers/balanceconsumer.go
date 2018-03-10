package consumers

import (
	"fmt"

	"github.com/DexyProject/dexy-go/contracts"
	"github.com/DexyProject/dexy-go/types"
	"github.com/ethereum/go-ethereum/event"
)

type BalanceChange struct {
	User   types.Address
	Amount types.Int
}

type BalanceConsumer struct {
	vault *contracts.Vault

	stop chan struct{}

	withdrawSink, depositSink chan<- BalanceChange
	withdraw, deposit         event.Subscription
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
			bc.depositSink <- balanceChangeForDeposit(*deposit)
		case withdraw := <-withdrawn:
			bc.withdrawSink <- balanceChangeForWithdraw(*withdraw)
		case <-bc.stop:
			return
		}
	}
}

func balanceChangeForDeposit(deposit contracts.VaultDeposited) BalanceChange {
	return BalanceChange{User: types.Address{Address: deposit.User}, Amount: types.Int{Int: *deposit.Amount}}
}

func balanceChangeForWithdraw(withdraw contracts.VaultWithdrawn) BalanceChange {
	return BalanceChange{User: types.Address{Address: withdraw.User}, Amount: types.Int{Int: *withdraw.Amount}}
}
