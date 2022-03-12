package auth

import (
	"github.com/gofiber/fiber/v2"
)

func AuthRouter(r fiber.Router) {
	r.Post("/register", RegisterController)
	r.Post("/login", LoginController)
}
