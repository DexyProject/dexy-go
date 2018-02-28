package consumers

import "github.com/DexyProject/dexy-go/types"

type Message interface {
	Ack()
	Reject()
}

type TradedMessage struct {
	Transaction types.Transaction

	ack    chan<- types.Hash
	reject chan<- types.Hash
}

func NewTradedMessage(tx types.Transaction, ack, reject chan<- types.Hash) *TradedMessage {
	return &TradedMessage{Transaction: tx, ack: ack, reject: reject}
}

func (tm TradedMessage) Ack() {
	tm.ack <- tm.Transaction.TransactionID
}

func (tm TradedMessage) Reject() {
	tm.reject <- tm.Transaction.TransactionID
}

type CancelledMessage struct {
	Hash types.Hash

	ack    chan<- types.Hash
	reject chan<- types.Hash
}

func NewCancelledMessage(hash types.Hash, ack, reject chan<- types.Hash) *CancelledMessage {
	return &CancelledMessage{Hash: hash, ack: ack, reject: reject}
}

func (cm CancelledMessage) Ack() {
	cm.ack <- cm.Hash
}

func (cm CancelledMessage) Reject() {
	cm.reject <- cm.Hash
}
