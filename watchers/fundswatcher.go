package watchers

import "github.com/DexyProject/dexy-go/contracts"

type FundsWatcher struct {
	vault contracts.Vault
}

func (fw *FundsWatcher) Watch() {
}

func (fw *FundsWatcher) handleWithdraw() {

	withdrawn := 0

	// @todo call on orders

	onOrders := 0
	balance := 0

	if onOrders < balance {
		// @todo updated db
	}
}

func (fw *FundsWatcher) handleDeposit() {

	// @todo calculate the amount which is underfunded

	// @todo update mongo daatabase

}
