package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/DexyProject/dexy-go/balances"
	"github.com/DexyProject/dexy-go/contracts"
	"github.com/DexyProject/dexy-go/endpoints"
	"github.com/DexyProject/dexy-go/history"
	"github.com/DexyProject/dexy-go/orderbook"
	"github.com/DexyProject/dexy-go/ticks"
	"github.com/DexyProject/dexy-go/validators"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {

	defer deferOnPanic()

	ethNode := flag.String("ethnode", "", "ethereum node address")
	mongo := flag.String("mongo", "", "mongodb connection string")
	vaultaddr := flag.String("vault", "", "vault address")

	flag.Parse()

	if flag.NFlag() != 3 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	v, err := setupVault(*ethNode, common.HexToAddress(*vaultaddr))
	if err != nil {
		log.Fatal(err)
	}

	bv, err := setupBalanceValidator(v, *mongo)
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	setupOrderBookEndpoints(*mongo, bv, v, r)
	setupMarketsEndpoints(*mongo, r)
	setupHistoryEndpoints(*mongo, r)
	setupTickEndpoint(*mongo, r)
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

	err = http.ListenAndServe(":9000", handlers.CORS(originsOk, headersOk, methodsOk)(r))
	if err != nil {
		log.Fatalf("Listen: %s", err.Error())
	}
}

func setupHistoryEndpoints(mongo string, r *mux.Router) {
	h, err := history.NewMongoHistory(mongo)
	if err != nil {
		log.Fatal("History:", err)
	}

	endpoint := endpoints.History{History: h}
	r.HandleFunc("/trades", endpoint.Handle).Methods("GET").Queries("token", "")
}

func setupOrderBookEndpoints(mongo string, bv validators.BalanceValidator, v *contracts.Vault, r *mux.Router) {
	ob, err := orderbook.NewMongoOrderBook(mongo)
	if err != nil {
		log.Fatalf("Orderbook error: %v", err.Error())
	}

	orders := endpoints.Orders{OrderBook: ob, BalanceValidator: bv, Vault: v}

	r.HandleFunc("/orderbook", orders.GetOrderBook).Methods("GET", "HEAD").Queries("token", "")
	r.HandleFunc("/orders", orders.GetOrders).Methods("GET", "HEAD").Queries("token", "")
	r.HandleFunc("/orders", orders.CreateOrder).Methods("POST")
	r.HandleFunc("/orders/{order}", orders.GetOrder).Methods("GET", "HEAD")
}

func setupMarketsEndpoints(mongo string, r *mux.Router) {
	ob, err := orderbook.NewMongoOrderBook(mongo)
	if err != nil {
		log.Fatalf("Orderbook error: %v", err.Error())
	}

	markets := endpoints.Markets{OrderBook: ob}

	r.HandleFunc("/markets", markets.GetMarkets).Methods("GET", "HEAD")
}

func setupTickEndpoint(mongo string, r *mux.Router) {
	tickdb, err := ticks.NewMongoTicks(mongo)
	if err != nil {
		log.Fatalf("tickdb error: %s", err.Error())
	}

	t := endpoints.Ticks{Ticks: tickdb}

	r.HandleFunc("/ticks", t.GetTicks).Methods("GET", "HEAD").Queries("token", "")
}

func setupBalanceValidator(v *contracts.Vault, mongo string) (validators.BalanceValidator, error) {

	b, err := balances.NewMongoBalances(mongo)
	if err != nil {
		return nil, err
	}

	return validators.NewRPCBalanceValidator(v, b), nil
}

func setupVault(ethereum string, addr common.Address) (*contracts.Vault, error) {
	conn, err := ethclient.Dial(ethereum)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the Ethereum client: %v", err)
	}

	return contracts.NewVault(addr, conn)
}

func deferOnPanic() {
	if err := recover(); err != nil {
		fmt.Println(err)
	}
}
