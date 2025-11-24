package main

type UserRepository interface {
	CreateUserAccount(*User) error
}

func (pb *PostgreDB) CreateUserAccount(user *User) error {
	query := `
        INSERT INTO account (first_name, last_name, email)
        VALUES ($1, $2, $3)
    `
	_, err := pb.db.Exec(query, user.FName, user.LName, user.Email)
	return err
}
