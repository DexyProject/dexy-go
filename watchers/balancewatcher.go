package watchers

import (
	"github.com/DexyProject/dexy-go/contracts"
	"github.com/DexyProject/dexy-go/types"
)

// @todo maybe a new name
type Balance struct {
	User types.Address
	Token types.Address
}

type BalanceWatcher struct {
	vault    *contracts.Vault

	in <-chan *Balance
}

func (bw *BalanceWatcher) Watch() {
	for {
		b := <- bw.in
		go bw.handle(b)
	}
}

func (bw *BalanceWatcher) handle(b *Balance) {

	if !hasOrders {
		return
	}

	checkBalance

	if balance > 0 {
		// set STATUS OPEN
		return
	}

	// set STATUS UNDERFUNDED

}
