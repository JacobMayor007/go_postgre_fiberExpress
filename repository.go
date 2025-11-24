package main

type UserRepository interface {
	CreateUserAccount(*User) error
	CreateProduct(*Product) error
}

func (pb *PostgreDB) CreateUserAccount(user *User) error {
	query := `
        INSERT INTO account (first_name, last_name, email)
        VALUES ($1, $2, $3)
    `
	_, err := pb.db.Exec(query, user.FName, user.LName, user.Email)
	return err
}

func (pb *PostgreDB) CreateProduct(product *Product) error {
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
	_, err := pb.db.Exec(query,
		product.User_Id,
		product.Name,
		product.Description,
		product.Stock,
		product.Price,
		product.PaymentMethod,
	)

	return err
}
