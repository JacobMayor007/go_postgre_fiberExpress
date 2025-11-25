package routes

import (
	"go+postgre/api"
	"go+postgre/repository"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App, repo repository.UserRepository) {
	userApi := api.UserReposit{
		UserRepo: repo,
	}

	app.Post("/user", userApi.CreateUser)
	app.Get("/user", userApi.GetUserById)
}
