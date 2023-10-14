package database

var schema = `
    CREATE TABLE if not exists users(
        id int PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
        username TEXT,
        password bytea
    );
    
    CREATE TABLE if not exists stock(
        id int PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
        symbol text,
        amount int,
        buy_price DECIMAL(10, 2),
        owner_id int,
        CONSTRAINT fk_users FOREIGN KEY(owner_id) 
	  		REFERENCES users(id)
    );
`
