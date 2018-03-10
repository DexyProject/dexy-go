package watchers

import (
	"github.com/DexyProject/dexy-go/consumers"
	"github.com/DexyProject/dexy-go/contracts"
	"github.com/DexyProject/dexy-go/balances"
)

type FundsWatcher struct {
	vault contracts.Vault

	balances balances.Balances

	withdraws, deposits <-chan consumers.BalanceChange
}

func (fw *FundsWatcher) Watch() {
	for {
		select {
		case withdraw := <-fw.withdraws:
			go fw.handleWithdraw(withdraw)
		case deposit := <-fw.deposits:
			go fw.handleDeposit(deposit)
		}
	}
}

func (fw *FundsWatcher) handleWithdraw(change consumers.BalanceChange) {

	withdrawn := change.Amount

	// @todo call on orders

	onOrders, err := fw.balances.OnOrders(change.User, change.Token)
	if err != nil {

	}

	balance := 0

	if onOrders < balance {
		// @todo updated db
	}
}

func (fw *FundsWatcher) handleDeposit(change consumers.BalanceChange) {

	// @todo calculate the amount which is underfunded

	// @todo update mongo daatabase

}
