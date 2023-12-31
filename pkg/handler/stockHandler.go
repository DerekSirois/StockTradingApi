package handler

import (
	"StockTrading/pkg/auth"
	"StockTrading/pkg/database"
	"StockTrading/pkg/fetching"
	"StockTrading/pkg/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

type Stock struct {
	Symbol string
	Amount int
}

func BuyStock(w http.ResponseWriter, r *http.Request) {
	s := &Stock{}
	err := json.NewDecoder(r.Body).Decode(s)
	if err != nil {
		utils.Respond(w, &utils.Response{Msg: fmt.Sprintf("couldn't decode json:%v", err)}, http.StatusBadRequest)
		return
	}

	ownerId, err := auth.GetAuthenticatedUserId(r)
	if err != nil {
		utils.Respond(w, &utils.Response{Msg: err.Error()}, http.StatusBadRequest)
		return
	}

	sFetch, err := fetching.GetStockInfo(s.Symbol)
	if err != nil {
		utils.Respond(w, &utils.Response{Msg: fmt.Sprintf("couldn't find the stock: %v", err)}, http.StatusBadRequest)
		return
	}

	sDb := MapStockToDatabase(s, ownerId, sFetch.Price)
	err = sDb.Create()
	if err != nil {
		utils.Respond(w, &utils.Response{Msg: fmt.Sprintf("couldn't create databse entry")}, http.StatusInternalServerError)
		return
	}

	utils.Respond(w, &utils.Response{Msg: fmt.Sprintf("Stock bought successfully at %v$", sDb.BuyPrice)}, http.StatusOK)
}

func SellStock(w http.ResponseWriter, r *http.Request) {
	s := &Stock{}
	err := json.NewDecoder(r.Body).Decode(s)
	if err != nil {
		utils.Respond(w, &utils.Response{Msg: fmt.Sprintf("couldn't decode json:%v", err)}, http.StatusBadRequest)
		return
	}

	ownerId, err := auth.GetAuthenticatedUserId(r)
	if err != nil {
		utils.Respond(w, &utils.Response{Msg: err.Error()}, http.StatusBadRequest)
		return
	}

	sDb, err := database.GetAmountOfStock(ownerId, s.Symbol)
	if err != nil {
		utils.Respond(w, &utils.Response{Msg: fmt.Sprintf("couldn't find the stock: %v", err)}, http.StatusBadRequest)
		return
	}
	//check if the user have more stocks then he is trying to sell
	if sDb.Amount < s.Amount {
		utils.Respond(w, &utils.Response{Msg: "You cannot sell more then you have"}, http.StatusBadRequest)
		return
	}

	s.Amount = -s.Amount
	sFetch, err := fetching.GetStockInfo(s.Symbol)
	if err != nil {
		utils.Respond(w, &utils.Response{Msg: fmt.Sprintf("couldn't find the stock: %v", err)}, http.StatusBadRequest)
		return
	}

	sDb = MapStockToDatabase(s, ownerId, sFetch.Price)
	err = sDb.Create()
	if err != nil {
		utils.Respond(w, &utils.Response{Msg: fmt.Sprintf("couldn't create databse entry")}, http.StatusInternalServerError)
		return
	}

	utils.Respond(w, &utils.Response{Msg: fmt.Sprintf("Stock sold successfully at %v$", sDb.BuyPrice)}, http.StatusOK)
}

func MapStockToDatabase(s *Stock, ownerId int, price float32) *database.Stock {
	return &database.Stock{
		Symbol:   s.Symbol,
		Amount:   s.Amount,
		BuyPrice: price,
		OwnerId:  ownerId,
	}
}
