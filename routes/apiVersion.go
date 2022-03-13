package routes

import (
	"github.com/Reynadi531/sfhacks-backend-go/controllers/v1/auth"
	"github.com/Reynadi531/sfhacks-backend-go/controllers/v1/logger"
	"github.com/gofiber/fiber/v2"
)

func V1Router(r fiber.Router) {
	auth.AuthRouter(r.Group("/auth"))
	logger.LoggerRouter(r.Group("/logger"))
}
