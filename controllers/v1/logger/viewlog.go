package logger

import (
	"fmt"
	"strings"

	"github.com/Reynadi531/sfhacks-backend-go/database"
	"github.com/Reynadi531/sfhacks-backend-go/utils"
	"github.com/gofiber/fiber/v2"
)

func ViewLogController(c *fiber.Ctx) error {
	if len(c.GetReqHeaders()["Authorization"]) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "authorization required",
		})
	}

	at := strings.Split(c.GetReqHeaders()["Authorization"], "Bearer ")

	jwtclaims, ok := utils.ExtractClaims(at[1])

	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "invalid token",
		})
	}

	userid := fmt.Sprintf("%s", jwtclaims["userid"])

	result, err := database.GetLogsByUserId(userid, 100)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "failed resolve data",
		})

	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "success retrive logs",
		"data":    result,
	})

}
