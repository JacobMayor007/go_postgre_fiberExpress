package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type PostgreDB struct {
	Db *sql.DB
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
		Db: db,
	}, nil
}

func (pb *PostgreDB) Init() error {
	_, err := pb.Db.Exec(`
		create extension if not exists "uuid-ossp"
	`)

	if err != nil {
		return err
	}

	pb.createUserTable()
	pb.createProductTable()
	return nil
}

func (pb *PostgreDB) createUserTable() error {
	query := `create table if not exists account (
		user_id uuid primary key default uuid_generate_v4(),
		email text unique,
		first_name varchar(50),
		last_name varchar(50),
    	created_at timestamp default now()
	)`

	_, err := pb.Db.Exec(query)
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

	_, err = pb.Db.Exec(funcTrig)
	if err != nil {
		return err
	}

	trigger := `create trigger created_at_trigger
		before insert on account
		for each row

	`

	_, err = pb.Db.Exec(trigger)
	return err
}

func (pb *PostgreDB) createProductTable() error {

	query := `
		create table if not exists products (
		product_id uuid primary key default uuid_generate_v4(),
		user_id uuid references account (user_id),
		product_name varchar(50),
		product_description text,
		product_stock int2,
		product_price int4,
		product_paymentMethod text,
   		created_at timestamp default now()
	)`

	_, err := pb.Db.Exec(query)
	if err != nil {
		fmt.Printf("error: %v", err.Error())
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

	_, err = pb.Db.Exec(funcTrig)
	if err != nil {
		return err
	}

	trigger := `create trigger created_at_trigger
		before insert on products
		for each row
	`

	_, err = pb.Db.Exec(trigger)

	return err
}
