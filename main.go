package main

import (
	"go+postgre/database"
	"go+postgre/repository"
	"go+postgre/routes"
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

	db, err := database.NewPostgreDB()
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Init(); err != nil {
		log.Fatal(err)
	}

	userRepo := repository.NewUserRepository(db)
	prodRepo := repository.ProdDbNew(db)

	routes.SetupRoutes(server, userRepo, prodRepo)

	log.Fatal(server.Listen(":3000"))
}
