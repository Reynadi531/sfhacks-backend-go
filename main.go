package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Reynadi531/sfhacks-backend-go/database"
	"github.com/Reynadi531/sfhacks-backend-go/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())
	app.Use(recover.New())

	database.InitDatabase()
	defer func() {
		database.DBDisconnect()
	}()

	api := app.Group("/api")
	routes.V1Router(api.Group("/v1"))

	var PORT string
	PORT = "3000"
	if len(os.Getenv("PORT")) != 0 {
		PORT = os.Getenv("PORT")
	}

	log.Fatal(app.Listen(fmt.Sprintf(":%s", PORT)))
}
