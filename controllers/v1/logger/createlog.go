package logger

import (
	"fmt"
	"strings"
	"time"

	"github.com/Reynadi531/sfhacks-backend-go/database"
	"github.com/Reynadi531/sfhacks-backend-go/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type LoggerRequest struct {
	Heartrate int16  `validate:"required,number"`
	Emotion   string `validate:"required"`
	Result    string `validate:"required"`
}

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

var validate = validator.New()

func ValidateLoggerStruct(log LoggerRequest) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(log)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

func CreateLogController(c *fiber.Ctx) error {
	if len(c.GetReqHeaders()["Authorization"]) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "authorization required",
		})
	}

	log := new(LoggerRequest)

	if err := c.BodyParser(log); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
		})
	}

	errors := ValidateLoggerStruct(*log)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "validation failed",
			"data":    errors,
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

	resultid, err := database.InsertLog(database.LoggerBSON{
		UserId:    fmt.Sprintf("%s", jwtclaims["userid"]),
		Heartrate: log.Heartrate,
		Emotion:   log.Emotion,
		Result:    log.Result,
		Time:      time.Now(),
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "failed insert data to database",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "success saved log",
		"data": fiber.Map{
			"id": resultid,
		},
	})
}
