package main

import (
	"StockTrading/pkg/auth"
	"StockTrading/pkg/database"
	"StockTrading/pkg/handler"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", handler.Index).Methods("GET")
	router.HandleFunc("/login", handler.Login).Methods("POST")
	router.HandleFunc("/register", handler.Register).Methods("POST")
	router.HandleFunc("/fetch/{symbol}", auth.VerifyJWT(handler.FetchStock)).Methods("GET")

	err := database.InitDb()
	handleError(err)

	log.Println("Serving on port 8080")
	log.Println(http.ListenAndServe(":8080", router))
}

func handleError(err error) {
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
