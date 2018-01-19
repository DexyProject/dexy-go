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

	ob, err := orderbook.NewMongoOrderBook("localhost") // @todo
	if err != nil {
		log.Fatal("Mongo error", err)
	}

	getorders := handlers.GetOrdersHandler{OrderBook: ob}
	getorder := handlers.GetOrderHandler{OrderBook: ob}
	createorder := handlers.CreateOrderHandler{OrderBook: ob}

	r.HandleFunc("/orders", getorders.Handle).Methods("GET").Queries("token", "")
	r.HandleFunc("/orders", createorder.Handle).Methods("POST")
	r.HandleFunc("/orders/{order}", getorder.Handle).Methods("GET")

	//http.Handle("/", r)

	err := http.ListenAndServe(":12312", r)
	if err != nil {
		log.Fatal("Listen:", err)
	}
}
