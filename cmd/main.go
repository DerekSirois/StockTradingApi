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
	router.HandleFunc("/fetch/{symbol}", handler.FetchStock).Methods("GET")

	log.Println("Serving on port 8080")
	log.Println(http.ListenAndServe(":8080", router))
}
