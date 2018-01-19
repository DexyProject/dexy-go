package main

import (
	"log"
	"net/http"
	"os"

	"github.com/DexyProject/dexy-go/handlers"
	"github.com/DexyProject/dexy-go/orderbook"
	"github.com/gorilla/mux"
	"fmt"
)

func main() {

	defer deferOnPanic()

	ob, err := orderbook.NewMongoOrderBook(os.Args[1]) // @todo
	if err != nil {
		log.Fatalf("Orderbook error: %v", err.Error())
	}

	getorders := handlers.GetOrdersHandler{OrderBook: ob}
	getorder := handlers.GetOrderHandler{OrderBook: ob}
	createorder := handlers.CreateOrderHandler{OrderBook: ob}

	r := mux.NewRouter()
	r.HandleFunc("/orders", getorders.Handle).Methods("GET").Queries("token", "")
	r.HandleFunc("/orders", createorder.Handle).Methods("POST")
	r.HandleFunc("/orders/{order}", getorder.Handle).Methods("GET")
	http.Handle("/", r)

	err = http.ListenAndServe(":12312", r)
	if err != nil {
		log.Fatalf("Listen: %s", err.Error())
	}
}

func deferOnPanic() {
	if err := recover(); err != nil {
		fmt.Println(err)
	}
}
