package watchers

import (
	"github.com/DexyProject/dexy-go/contracts"
	"github.com/DexyProject/dexy-go/orderbook"
	"github.com/DexyProject/dexy-go/types"
	"math/big"
	"github.com/DexyProject/dexy-go/log"
	"go.uber.org/zap"
)

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
		log.Error("order check failed", zap.Error(err))
		return
	}

	if !hasOrders {
		return
	}

	balance, err := bw.vault.BalanceOf(nil, b.Token.Address, b.User.Address)
	if err != nil {
		log.Error("balance call failed", zap.Error(err))
		return
	}

	status := types.UNDER_FUNDED
	if balance.Cmp(big.NewInt(0)) > 0 {
		status = types.OPEN
	}

	err = bw.ob.SetOrderStatuses(b.Token, b.User, status)
	if err != nil {
		log.Error("updating order status failed", zap.Error(err))
	}
}
