package consumers

import "github.com/DexyProject/dexy-go/types"

type Message interface {
	Ack()
	Reject()
}

type TradedMessage struct {
	Transaction types.Transaction
}

func NewTradedMessage(tx types.Transaction) *TradedMessage {
	return &TradedMessage{Transaction: tx}
}

// @todo
func (TradedMessage) Ack() {
	panic("implement me")
}

func (TradedMessage) Reject() {
	panic("implement me")
}
