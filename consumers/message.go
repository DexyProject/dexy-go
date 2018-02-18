package consumers

import "github.com/DexyProject/dexy-go/types"

type Message interface {
	Ack()
	Reject()
}

type TradedMessage struct {
	Transaction types.Transaction

	ack    chan<- types.Bytes
	reject chan<- types.Bytes
}

func NewTradedMessage(tx types.Transaction, ack, reject chan<- types.Bytes) *TradedMessage {
	return &TradedMessage{Transaction: tx, ack: ack, reject: reject}
}

func (tm TradedMessage) Ack() {
	tm.ack <- tm.Transaction.TransactionID
}

func (tm TradedMessage) Reject() {
	tm.reject <- tm.Transaction.TransactionID
}
