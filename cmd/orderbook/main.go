package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/DexyProject/dexy-go/endpoints"
	"github.com/DexyProject/dexy-go/orderbook"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
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

	headersOk := handlers.AllowedHeaders([]string{"Content-Type", "X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	err = http.ListenAndServe(":12312", handlers.CORS(originsOk, headersOk, methodsOk)(r))
	if err != nil {
		log.Fatalf("Listen: %s", err.Error())
	}
}

func deferOnPanic() {
	if err := recover(); err != nil {
		fmt.Println(err)
	}
}
