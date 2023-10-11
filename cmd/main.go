package main

import (
	"StockTrading/pkg/handler"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", handler.Index).Methods("GET")

	log.Println(http.ListenAndServe(":8080", router))
}
