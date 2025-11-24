package main

import (
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {

	server := fiber.New()
	server.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Accept, Content-Type, Origins",
		AllowMethods: strings.Join([]string{
			fiber.MethodGet,
			fiber.MethodPost,
			fiber.MethodHead,
			fiber.MethodPut,
			fiber.MethodDelete,
			fiber.MethodPatch,
		}, ","),
	}))

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error on loading env data")
	}

	user, err := NewPostgreDB()
	if err != nil {
		log.Fatal(err)
	}

	if err := user.Init(); err != nil {
		log.Fatal(err)
	}

	userApi := &UserReposit{
		UserRepo: user,
	}

	server.Get("/", func(res *fiber.Ctx) error {
		return res.SendStatus(fiber.StatusOK)
	})

	server.Post("/user", userApi.CreateUser)
	server.Post("/product", userApi.CreateProduct)
	server.Get("/user", userApi.GetUserById)

	log.Fatal(server.Listen(":3000"))
}
