package routes

import (
	"github.com/Reynadi531/sfhacks-backend-go/controllers/v1/auth"
	"github.com/gofiber/fiber/v2"
)

func V1Router(r fiber.Router) {
	auth.AuthRouter(r.Group("/auth"))
}
