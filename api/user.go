package api

import (
	"go+postgre/repository"
	"go+postgre/types"

	"github.com/gofiber/fiber/v2"
)

type UserReposit struct {
	UserRepo repository.UserRepository
}

func (u *UserReposit) CreateUser(f *fiber.Ctx) error {
	var user types.User

	if err := f.BodyParser(&user); err != nil {
		return f.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := u.UserRepo.CreateUserAccount(&user); err != nil {
		return f.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return f.JSON(fiber.Map{
		"message": "User created successfully",
	})
}

func (u *UserReposit) GetUserById(f *fiber.Ctx) error {

	var body struct {
		ID string `json:"id"`
	}

	if err := f.BodyParser(&body); err != nil {
		return f.Status(404).JSON(fiber.Map{
			"title":   "Getting user is not successful",
			"message": "" + err.Error(),
		})
	}

	user, err := u.UserRepo.GetUserById(body.ID)
	if err != nil {
		return f.Status(404).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return f.Status(200).JSON(user)
}
