package repository

import (
	"go+postgre/database"
	"go+postgre/types"
)

type ProdRepo interface {
	CreateProduct(*types.Product) error
	GetProductById(id string) (*types.ProductUser, error)
	UpdateProductById(id, product_name string, product_stock int16) error
	DeleteProductById(id string) error
	MultipleDeletion(user_id string) error
}

type ProdDb struct {
	DB *database.PostgreDB
}

func ProdDbNew(db *database.PostgreDB) *ProdDb {
	return &ProdDb{
		DB: db,
	}
}

func (pb *ProdDb) CreateProduct(product *types.Product) error {
	query := `
		INSERT INTO products 
		(
		user_id, 
		product_name, 
		product_description, 
		product_stock,
		product_price,
		product_paymentMethod
		)
        VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := pb.DB.Db.Exec(query,
		product.User_Id,
		product.Name,
		product.Description,
		product.Stock,
		product.Price,
		product.PaymentMethod,
	)

	return err
}

func (pb *ProdDb) GetProductById(id string) (*types.ProductUser, error) {

	rows := pb.DB.Db.QueryRow(`
	 SELECT 
            p.product_name, 
            p.product_description,
            p.product_price,
            p.product_stock,
			p.user_id,
			a.email,
			a.first_name,
			a.last_name
        FROM products p
        INNER JOIN account a ON p.user_id = a.user_id
        WHERE p.product_id = $1`, id)

	product := &types.ProductUser{}

	err := rows.Scan(
		&product.Product.Name,
		&product.Product.Description,
		&product.Product.Price,
		&product.Product.Stock,
		&product.Product.User_Id,
		&product.User.Email,
		&product.User.FName,
		&product.User.LName,
	)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (pb *ProdDb) UpdateProductById(id, product_name string, product_stock int16) error {
	statement := `
		update products
		set product_name = $2, product_stock = $3
		where product_id = $1;
	`
	_, err := pb.DB.Db.Exec(statement, id, product_name, product_stock)

	return err
}

func (pb *ProdDb) DeleteProductById(id string) error {
	statement := `
		delete from products
		where product_id = $1;
	`

	_, err := pb.DB.Db.Exec(statement, id)

	return err
}

func (pb *ProdDb) MultipleDeletion(user_id string) error {
	statement := `
		delete from products
		where user_id = $1
	`

	_, err := pb.DB.Db.Exec(statement, user_id)

	return err
}
