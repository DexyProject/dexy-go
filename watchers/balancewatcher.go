package watchers

import (
	"github.com/DexyProject/dexy-go/contracts"
	"github.com/DexyProject/dexy-go/orderbook"
	"github.com/DexyProject/dexy-go/types"
	"math/big"
)

// @todo maybe a new name
type Balance struct {
	User  types.Address
	Token types.Address
}

type BalanceWatcher struct {
	ob    orderbook.OrderBook
	vault *contracts.Vault

	in <-chan *Balance
}

func (bw *BalanceWatcher) Watch() {
	for {
		b := <-bw.in
		go bw.handle(b)
	}
}

func (bw *BalanceWatcher) handle(b *Balance) {
	hasOrders, err := bw.ob.HasOrders(b.Token, b.User)
	if err != nil {
		// @todo log
		return
	}

	if !hasOrders {
		return
	}

	balance, err := bw.vault.BalanceOf(nil, b.Token.Address, b.User.Address)
	if err != nil {
		// @todo log
		return
	}

	if balance.Cmp(big.NewInt(0)) > 0 {
		// set STATUS OPEN
		return
	}

	// set STATUS UNDERFUNDED
}
