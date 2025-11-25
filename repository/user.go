package repository

import (
	"go+postgre/database"
	"go+postgre/types"
)

type UserRepository interface {
	CreateUserAccount(*types.User) error
	GetUserById(id string) (*types.User, error)
}

type UserReposit struct {
	DB *database.PostgreDB
}

func NewUserRepository(db *database.PostgreDB) *UserReposit {
	return &UserReposit{
		DB: db,
	}
}

func (pb *UserReposit) CreateUserAccount(user *types.User) error {
	query := `
        INSERT INTO account (first_name, last_name, email)
        VALUES ($1, $2, $3)
    `

	_, err := pb.DB.Db.Exec(query, user.FName, user.LName, user.Email)
	return err
}

func (pb *UserReposit) GetUserById(id string) (*types.User, error) {

	rows := pb.DB.Db.QueryRow(`
        SELECT email, first_name, last_name
        FROM account
        WHERE user_id = $1
    `, id)

	user := new(types.User)

	err := rows.Scan(
		&user.Email,
		&user.FName,
		&user.LName,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}
