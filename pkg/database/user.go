package database

import "fmt"

type Users struct {
	Id       int
	Username string
	Password []byte
}

func (u *Users) Create() (int, error) {
	res, err := db.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", u.Username, u.Password)
	if err != nil {
		return 0, fmt.Errorf("couldn't insert user: %v", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("couldn't get the id: %v", err)
	}

	return int(id), nil
}

func GetByUsername(name string) (*Users, error) {
	u := &Users{}
	err := db.Get(u, "SELECT * FROM users WHERE username = $1", name)
	if err != nil {
		return nil, fmt.Errorf("couldn't get the user: %v", err)
	}

	return u, nil
}
