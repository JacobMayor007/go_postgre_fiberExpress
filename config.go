package main

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

type PostgreDB struct {
	db *sql.DB
}

func NewPostgreDB() (*PostgreDB, error) {
	password := os.Getenv("PASSWORD")
	connStr := "user=postgres dbname=postgres password=" + password + " sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgreDB{
		db: db,
	}, nil
}

func (pb *PostgreDB) Init() error {
	pb.createUserTable()
	pb.createProductTable()
	return nil
}

func (pb *PostgreDB) createUserTable() error {
	query := `create table if not exists account (
		id serial primary key,
		email text unique,
		first_name varchar(50),
		last_name varchar(50),
    	created_at timestamp default now()
	)`

	_, err := pb.db.Exec(query)
	if err != nil {
		return err
	}

	funcTrig := `
		create or replace function setCreatedAt()
		returns trigger as $$
		begin
			new.created_at = NOW();
			return new;
		end;
		$$ language plpgsql;

		create or replace function setUpdatedAt()
		returns trigger AS $$
		begin
			new.updated_at := NOW();
			return new;
		end;
		$$ language plpgsql;
	`

	_, err = pb.db.Exec(funcTrig)
	if err != nil {
		return err
	}

	trigger := `create trigger created_at_trigger
		before insert on account
		for each row

	`

	_, err = pb.db.Exec(trigger)
	return err
}

func (pb *PostgreDB) createProductTable() error {

	query := `create table if not exists products (
		product_id serial primary key,
		user_id int references account (id),
		product_name varchar(50),
		product_description text,
		product_stock int4,
		price serial,
		paymentMethod text,
   		created_at timestamp default now()
	)`

	_, err := pb.db.Exec(query)
	return err
}
