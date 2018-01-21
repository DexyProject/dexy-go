package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/DexyProject/dexy-go/handlers"
	"github.com/DexyProject/dexy-go/orderbook"
	"github.com/gorilla/mux"
	muxhandlers "github.com/gorilla/handlers"
	"github.com/DexyProject/dexy-go/endpoints"
)

func main() {

	defer deferOnPanic()

	ob, err := orderbook.NewMongoOrderBook(os.Args[1]) // @todo
	if err != nil {
		log.Fatalf("Orderbook error: %v", err.Error())
	}

	orders := endpoints.Orders{OrderBook: ob} // @todo balance validator

	r := mux.NewRouter()
	r.HandleFunc("/orders", orders.GetOrders).Methods("GET", "HEAD").Queries("token", "")
	r.HandleFunc("/orders", orders.CreateOrder).Methods("POST")
	r.HandleFunc("/orders/{order}", orders.GetOrder).Methods("GET", "HEAD")
	http.Handle("/", r)

	headersOk := muxhandlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := muxhandlers.AllowedOrigins([]string{"*"})
	methodsOk := muxhandlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	err = http.ListenAndServe(":12312", muxhandlers.CORS(originsOk, headersOk, methodsOk)(r))
	if err != nil {
		log.Fatalf("Listen: %s", err.Error())
	}
}

func deferOnPanic() {
	if err := recover(); err != nil {
		fmt.Println(err)
	}
}
