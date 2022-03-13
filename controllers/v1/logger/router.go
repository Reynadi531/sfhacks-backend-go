package logger

import "github.com/gofiber/fiber/v2"

func LoggerRouter(r fiber.Router) {
	r.Get("/", ViewLogController)

	r.Post("/", CreateLogController)
}
