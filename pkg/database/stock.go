package database

import "fmt"

type Stock struct {
	Id       int
	Symbol   string
	Amount   int
	BuyPrice float32
	OwnerId  int
}

func (s *Stock) Buy() error {
	_, err := db.Exec("INSERT INTO stock (symbol, amount, buy_price, owner_id) VALUES ($1, $2, $3, $4)", s.Symbol, s.Amount, s.BuyPrice, s.OwnerId)
	if err != nil {
		return fmt.Errorf("couldn't insert a stock: %v", err)
	}

	return nil
}

func GetBuyOwner(ownerId int) ([]*Stock, error) {
	s := make([]*Stock, 0)
	err := db.Select(s, "SELECT * from stock where owner_id = $1", ownerId)
	if err != nil {
		return nil, err
	}

	return s, nil
}
