package handler

import (
	"StockTrading/pkg/fetching"
	"StockTrading/pkg/utils"
	"fmt"
	"github.com/gorilla/mux"
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
		utils.Respond(w, &utils.Response{Msg: err.Error()}, http.StatusInternalServerError)
		return
	}

	utils.Respond(w, s, http.StatusOK)
}
