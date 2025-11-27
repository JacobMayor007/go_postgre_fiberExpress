package types

type User struct {
	FName string `json:"fName"`
	LName string `json:"lName"`
	Email string `json:"email"`
}

type Product struct {
	Name          string `json:"product_name"`
	Description   string `json:"product_description"`
	Stock         int16  `json:"product_stock"`
	Price         int32  `json:"product_price"`
	PaymentMethod string `json:"product_paymentMethod"`
	User_Id       string `json:"user_id"`
}

type ProductUser struct {
	Product Product
	User    User
}
