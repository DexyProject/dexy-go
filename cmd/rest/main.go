package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/DexyProject/dexy-go/balances"
	"github.com/DexyProject/dexy-go/endpoints"
	"github.com/DexyProject/dexy-go/history"
	"github.com/DexyProject/dexy-go/orderbook"
	"github.com/DexyProject/dexy-go/validators"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {

	defer deferOnPanic()

	r := mux.NewRouter()

	setupOrderBookEndpoints(r)
	setupHistoryEndpoints(r)

	http.Handle("/", r)

	headersOk := handlers.AllowedHeaders([]string{
		"Content-Type",
		"X-Requested-With",
		"Accept",
		"Accept-Language",
		"Accept-Encoding",
		"Content-Language",
		"Origin",
	})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	err := http.ListenAndServe(":12312", handlers.CORS(originsOk, headersOk, methodsOk)(r))
	if err != nil {
		log.Fatalf("Listen: %s", err.Error())
	}
}

func setupHistoryEndpoints(r *mux.Router) {
	h, err := history.NewMongoHistory(os.Args[1])
	if err != nil {
		log.Fatal("History:", err)
	}

	endpoint := endpoints.History{History: h}
	r.HandleFunc("/trades", endpoint.Handle).Methods("GET").Queries("token", "")
}

func setupOrderBookEndpoints(r *mux.Router) {
	ob, err := orderbook.NewMongoOrderBook(os.Args[1])
	if err != nil {
		log.Fatalf("Orderbook error: %v", err.Error())
	}

	validator, err := setupBalanceValidator()
	if err != nil {
		log.Fatalf("validator error: %v", err.Error())
	}

	orders := endpoints.Orders{OrderBook: ob, BalanceValidator: validator}

	r.HandleFunc("/orders", orders.GetOrders).Methods("GET", "HEAD").Queries("token", "")
	r.HandleFunc("/orders", orders.CreateOrder).Methods("POST")
	r.HandleFunc("/orders/{order}", orders.GetOrder).Methods("GET", "HEAD")
}

func setupBalanceValidator() (validators.BalanceValidator, error) {
	conn, err := ethclient.Dial(os.Args[2])
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the Ethereum client: %v", err)
	}

	b, err := balances.NewMongoBalances(os.Args[1])
	if err != nil {
		return nil, err
	}

	return &validators.RPCBalanceValidator{Conn: conn, Balances: b}, nil
}

func deferOnPanic() {
	if err := recover(); err != nil {
		fmt.Println(err)
	}
}
