package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/DexyProject/dexy-go/handlers/history"
	txhistory "github.com/DexyProject/dexy-go/history"
)

func main() {

	r := mux.NewRouter()

	h, err := txhistory.NewMongoHistory("")
	if err != nil {
		log.Fatal("History:", err)
	}

	get := &history.GetTradeHistoryHandler{History: h}
	r.HandleFunc("/trades", get.Handle).Methods("GET").Queries("token", "")

	err = http.ListenAndServe(":12312", r)
	if err != nil {
		log.Fatal("Listen:", err)
	}
}
