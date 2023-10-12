package database

import "fmt"

type Stock struct {
	Id       int
	Symbol   string
	Amount   int
	BuyPrice float32
	OwnerId  int
}

func (s *Stock) Buy() (int, error) {
	res, err := db.Exec("INSERT INTO stock (symbol, amount, buy_price, owner_id) VALUES ($1, $2, $3, $4)", s.Symbol, s.Amount, s.BuyPrice, s.OwnerId)
	if err != nil {
		return 0, fmt.Errorf("couldn't insert a stock: %v", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("couldn't get the id: %v", err)
	}

	return int(id), nil
}

func GetBuyOwner(ownerId int) ([]*Stock, error) {
	s := make([]*Stock, 0)
	err := db.Select(s, "SELECT * from stock where owner_id = $1", ownerId)
	if err != nil {
		return nil, err
	}

	return s, nil
}
