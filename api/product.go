package api

import (
	"fmt"
	"go+postgre/repository"
	"go+postgre/types"

	"github.com/gofiber/fiber/v2"
)

type UserReposit struct {
	UserRepo repository.UserRepository
}

type ProdReposit struct {
	ProdRepo repository.ProdRepo
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

func (pr *ProdReposit) CreateProduct(f *fiber.Ctx) error {
	var product types.Product

	if err := f.BodyParser(&product); err != nil {
		return err
	}

	if err := pr.ProdRepo.CreateProduct(&product); err != nil {
		fmt.Printf("Error Product: %s", err.Error())

		return f.Status(404).JSON(fiber.Map{
			"message": "Product creation failed",
		})
	}

	return f.JSON(fiber.Map{
		"message": "Product created successfully",
	})
}

func (pr *ProdReposit) GetProductById(f *fiber.Ctx) error {
	var body struct {
		Id string `json:"id"`
	}

	if err := f.BodyParser(&body); err != nil {
		return f.Status(404).JSON(fiber.Map{
			"title":   "Getting product is not successful",
			"message": "" + err.Error(),
		})
	}

	prod, err := pr.ProdRepo.GetProductById(body.Id)

	if err != nil {
		return f.Status(404).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return f.Status(200).JSON(prod)

}
