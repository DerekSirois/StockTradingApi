package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func Respond(w http.ResponseWriter, data any, status int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Println(err)
	}
}

type Response struct {
	Msg string
}
