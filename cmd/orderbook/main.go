package main

import (
	"log"
	"net/http"

	"github.com/DexyProject/dexy-go/handlers"
	"github.com/DexyProject/dexy-go/orderbook"
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	ob := &orderbook.MemoryOrderBook{}

	getorders := handlers.GetOrdersHandler{OrderBook: ob}
	getorder := handlers.GetOrderHandler{OrderBook: ob}
	createorder := handlers.CreateOrderHandler{OrderBook: ob}
	getticks := handlers.CreateHistoryHandler

	r.HandleFunc("/orders", getorders.Handle).Methods("GET").Queries("token", "")
	r.HandleFunc("/orders", createorder.Handle).Methods("POST")
	r.HandleFunc("/orders/{order}", getorder.Handle).Methods("GET")
	r.HandleFunc("/charts?token={token}", getticks.Handle).Methods("POST")

	//http.Handle("/", r)

	err := http.ListenAndServe(":12312", r)
	if err != nil {
		log.Fatal("Listen:", err)
	}
}
