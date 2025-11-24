package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type UserReposit struct {
	UserRepo UserRepository
}

func (u *UserReposit) CreateUser(f *fiber.Ctx) error {
	var user User

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

func (u *UserReposit) CreateProduct(f *fiber.Ctx) error {
	var product Product

	if err := f.BodyParser(&product); err != nil {
		return err
	}

	if err := u.UserRepo.CreateProduct(&product); err != nil {
		fmt.Printf("Error Product: %s", err.Error())

		return f.Status(404).JSON(fiber.Map{
			"message": "Product creation failed",
		})
	}

	return f.JSON(fiber.Map{
		"message": "Product created successfully",
	})
}
