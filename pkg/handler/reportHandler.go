package handler

import (
	"StockTrading/pkg/auth"
	"StockTrading/pkg/database"
	"StockTrading/pkg/fetching"
	"StockTrading/pkg/utils"
	"fmt"
	"net/http"
)

type StockReport struct {
	Symbol            string
	BuyValueTotal     float32
	CurrentValueTotal float32
	Percentage        float32
}

type Report struct {
	Stocks []*StockReport
}

func GetReport(w http.ResponseWriter, r *http.Request) {
	id, err := auth.GetAuthenticatedUserId(r)
	if err != nil {
		utils.Respond(w, &utils.Response{Msg: err.Error()}, http.StatusBadRequest)
		return
	}

	buy, err := database.GetBuyByOwner(id)
	if err != nil {
		utils.Respond(w, &utils.Response{Msg: fmt.Sprintf("couldn't get the stocks informations: %v", err)}, http.StatusInternalServerError)
		return
	}

	report, err := MapDatabaseToReport(buy)
	if err != nil {
		utils.Respond(w, &utils.Response{Msg: fmt.Sprintf("couldn't map the report: %v", err)}, http.StatusInternalServerError)
		return
	}

	utils.Respond(w, report, http.StatusOK)
}

func MapDatabaseToReport(buy []*database.Stock) (*Report, error) {
	r := &Report{}
	for _, stock := range buy {
		sFetch, err := fetching.GetStockInfo(stock.Symbol)
		if err != nil {
			return nil, err
		}
		curVal := float32(stock.Amount) * sFetch.Price
		r.Stocks = append(r.Stocks, &StockReport{
			Symbol:            stock.Symbol,
			BuyValueTotal:     stock.BuyPrice,
			CurrentValueTotal: curVal,
			Percentage:        ((curVal / stock.BuyPrice) - 1) * 100,
		})
	}
	return r, nil
}
