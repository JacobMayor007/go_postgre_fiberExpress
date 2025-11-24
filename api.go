package main

import "github.com/gofiber/fiber/v2"

type UserReposit struct {
	UserRepo UserRepository
}

func (u *UserReposit) CreateUser(c *fiber.Ctx) error {
	var user User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := u.UserRepo.CreateUserAccount(&user); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "User created successfully",
	})
}
