package handler

import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintln(w, "Welcome to the stock trading app")
}
