package fetching

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const APIKEY = "API_KEYf6GFU3ZXXNIZMR7FAMR5B5V522F64NB3"

func GetStockInfo(symbol string) (*Stock, error) {
	link := fmt.Sprintf("https://api.finage.co.uk/last/trade/stock/%s?apikey=%s", strings.ToUpper(symbol), APIKEY)
	res, err := http.Get(link)
	if err != nil {
		return nil, err
	}

	s := &Stock{}
	err = json.NewDecoder(res.Body).Decode(s)
	if err != nil {
		return nil, err
	}

	return s, nil
}
