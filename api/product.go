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

func (pr *ProdReposit) ProductHandle(f *fiber.Ctx) error {
	switch f.Method() {
	case "GET":
		return pr.GetProductById(f)
	case "DELETE":
		return pr.DeleteProductById(f)
	case "PUT":
		return pr.UpdateProductById(f)
	case "POST":
		return pr.CreateProduct(f)
	default:
		return fmt.Errorf("method not allowed %v", f.Method())
	}
}

func (pr *ProdReposit) MultipleProductsHandle(f *fiber.Ctx) error {
	switch f.Method() {
	case "DELETE":
		return pr.MultipleDeletion(f)
	default:
		return fmt.Errorf("method now allowed ")
	}
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

func (pr *ProdReposit) DeleteProductById(f *fiber.Ctx) error {
	var body struct {
		Id string `json:"id"`
	}

	if err := f.BodyParser(&body); err != nil {
		return f.Status(404).JSON(fiber.Map{
			"title":   "Error occured",
			"message": "Error: " + err.Error(),
		})
	}

	if err := pr.ProdRepo.DeleteProductById(body.Id); err != nil {
		return f.Status(404).JSON(fiber.Map{
			"title":   "Error occured",
			"message": "Error: " + err.Error(),
		})
	}

	return f.Status(200).JSON(fiber.Map{
		"title":      "Successful",
		"message":    "You have successfully deleted product",
		"successful": true,
	})
}

func (pr *ProdReposit) MultipleDeletion(f *fiber.Ctx) error {

	var body struct {
		UserId string `json:"user_id"`
	}

	if err := f.BodyParser(&body); err != nil {
		return f.Status(404).JSON(fiber.Map{
			"title":   "Error occured",
			"message": "Error: " + err.Error(),
		})
	}

	if err := pr.ProdRepo.MultipleDeletion(body.UserId); err != nil {
		return f.Status(404).JSON(fiber.Map{
			"title":   "Error occured",
			"message": "Error message: " + err.Error(),
		})
	}

	return f.Status(200).JSON(fiber.Map{
		"title":   "Successful",
		"message": "You have successfully deleted the product.",
	})
}
