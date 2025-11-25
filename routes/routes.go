package routes

import (
	"go+postgre/repository"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, user repository.UserRepository, prod repository.ProdRepo) {
	UserRoutes(app, user)
	ProductRoutes(app, prod)
}
