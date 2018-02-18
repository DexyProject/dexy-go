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

func (TradedMessage) Ack() {
	// @todo here we will want to send to a channel
}

func (TradedMessage) Reject() {
	// @todo here we will want to send to a channel
}
