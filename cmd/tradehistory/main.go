package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/DexyProject/dexy-go/handlers/history"
)

func main() {

	r := mux.NewRouter()

	get := &history.GetTradeHistoryHandler{}
	r.HandleFunc("/trades/{token}", get.Handle).Methods("GET")

	err := http.ListenAndServe(":12312", r)
	if err != nil {
		log.Fatal("Listen:", err)
	}
}
