package routes

import (
	"go+postgre/api"
	"go+postgre/repository"

	"github.com/gofiber/fiber/v2"
)

// Pass the repository to the route function
func ProductRoutes(app *fiber.App, repo repository.ProdRepo) {
	prodApi := api.ProdReposit{
		ProdRepo: repo,
	}
	app.Post("/product", prodApi.ProductHandle)
	app.Get("/product", prodApi.ProductHandle)
	app.Put("/product", prodApi.ProductHandle)
	app.Delete("/product", prodApi.ProductHandle)
	app.Delete("/products", prodApi.MultipleProductsHandle)
}
