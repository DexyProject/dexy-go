package main

import "github.com/DexyProject/dexy-go/watchers"

func main() {

	tf := watchers.TradeWatcher{}

	err := tf.Watch()
	if err != nil {
		// @todo
	}

}
