package watchers

import (
	"log"

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
			log.Printf("test: %+v", withdraw)
		case deposit := <-fw.deposits:
			log.Printf("test: %+v", deposit)
		}
	}
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
