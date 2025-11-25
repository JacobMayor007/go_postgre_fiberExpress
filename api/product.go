package api

import (
	"fmt"
	"go+postgre/repository"
	"go+postgre/types"

	"github.com/gofiber/fiber/v2"
)

type ProdReposit struct {
	ProdRepo repository.ProdRepo
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

func (pr *ProdReposit) UpdateProductById(f *fiber.Ctx) error {

	var body struct {
		ProductStock int16  `json:"product_stock"`
		Id           string `json:"id"`
		ProductName  string `json:"product_name"`
	}

	if err := f.BodyParser(&body); err != nil {
		return f.Status(404).JSON(fiber.Map{
			"title":   "Error occured",
			"message": "Error: " + err.Error(),
		})
	}

	err := pr.ProdRepo.UpdateProductById(body.Id, body.ProductName, body.ProductStock)
	if err != nil {
		return f.Status(404).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return f.Status(200).JSON(fiber.Map{
		"title":   "Product Updated Successfully",
		"message": "You have successfully updated your product",
	})
}
