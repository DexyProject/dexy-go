package watchers

import (
	"github.com/DexyProject/dexy-go/consumers"
	"github.com/DexyProject/dexy-go/contracts"
)

type FundsWatcher struct {
	vault contracts.Vault

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

	onOrders := 0
	balance := 0

	if onOrders < balance {
		// @todo updated db
	}
}

func (fw *FundsWatcher) handleDeposit(change consumers.BalanceChange) {

	// @todo calculate the amount which is underfunded

	// @todo update mongo daatabase

}
