package main

import (
	"log"
	"net/http"

	"github.com/DexyProject/dexy-go/endpoints"
	"github.com/DexyProject/dexy-go/history"
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	h, err := history.NewMongoHistory("")
	if err != nil {
		log.Fatal("History:", err)
	}

	get := &endpoints.History{History: h}
	r.HandleFunc("/trades", get.Handle).Methods("GET").Queries("token", "")

	err = http.ListenAndServe(":12312", r)
	if err != nil {
		log.Fatal("Listen:", err)
	}
}
