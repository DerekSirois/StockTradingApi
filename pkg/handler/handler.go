package handler

import (
	"StockTrading/pkg/fetching"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprintln(w, "Welcome to the stock trading app")
}

func FetchStock(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	symbol := vars["symbol"]

	s, err := fetching.GetStockInfo(symbol)
	if err != nil {
		respond(w, &response{msg: err.Error()}, http.StatusInternalServerError)
		return
	}

	respond(w, s, http.StatusOK)
}

func respond(w http.ResponseWriter, data any, status int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Println(err)
	}
}

type response struct {
	msg string
}
