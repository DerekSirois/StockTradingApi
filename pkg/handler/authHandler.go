package handler

import (
	"StockTrading/pkg/auth"
	"StockTrading/pkg/database"
	"StockTrading/pkg/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Username string
	Password string
}

type responseToken struct {
	Token string
}

func Register(w http.ResponseWriter, r *http.Request) {
	u := &User{}
	err := json.NewDecoder(r.Body).Decode(u)
	if err != nil {
		utils.Respond(w, utils.Response{Msg: err.Error()}, http.StatusBadRequest)
		return
	}
	ud, err := MapUserToDatabase(u)
	if err != nil {
		utils.Respond(w, utils.Response{Msg: err.Error()}, http.StatusInternalServerError)
		return
	}

	err = ud.Create()
	if err != nil {
		utils.Respond(w, utils.Response{Msg: err.Error()}, http.StatusInternalServerError)
		return
	}

	utils.Respond(w, utils.Response{Msg: "User created successfully"}, http.StatusOK)
}

func Login(w http.ResponseWriter, r *http.Request) {
	u := &User{}
	err := json.NewDecoder(r.Body).Decode(u)
	if err != nil {
		utils.Respond(w, utils.Response{Msg: err.Error()}, http.StatusBadRequest)
		return
	}

	ud, err := database.GetByUsername(u.Username)
	if err != nil {
		utils.Respond(w, utils.Response{Msg: err.Error()}, http.StatusInternalServerError)
		return
	}

	if !auth.CheckPasswordHash(u.Password, ud.Password) {
		utils.Respond(w, utils.Response{Msg: "Wrong username or password"}, http.StatusUnauthorized)
		return
	}

	token, err := auth.CreateJWTToken(ud.Id, ud.Username)
	if err != nil {
		utils.Respond(w, utils.Response{Msg: fmt.Sprintf("couln't generate the token: %v", err)}, http.StatusInternalServerError)
	}

	utils.Respond(w, responseToken{Token: token}, http.StatusOK)
}

func MapUserToDatabase(u *User) (*database.Users, error) {
	hash, err := auth.HashPassword(u.Password)
	if err != nil {
		return nil, err
	}
	return &database.Users{
		Username: u.Username,
		Password: hash,
	}, nil
}
