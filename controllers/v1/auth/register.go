package auth

import (
	"github.com/Reynadi531/sfhacks-backend-go/database"
	"github.com/Reynadi531/sfhacks-backend-go/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Name     string `validate:"required,min=3,max=32"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=6,max=32"`
	Phone    string `validate:"required,min=4"`
	Age      int8   `validate:"required"`
	Gender   string `validate:"required"`
}

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

var validate = validator.New()

func ValidateStruct(user User) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(user)
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

func RegisterController(c *fiber.Ctx) error {
	user := new(User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	errors := ValidateStruct(*user)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "validation failed",
			"data":    errors,
		})

	}

	isExist, err := utils.CheckUser(user.Email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "failed check if user exist",
		})
	}

	if isExist {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "user already exist",
		})
	}

	hashedpassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "failed hash password",
		})
	}

	uuidid := uuid.New()

	inputuser := database.UserBSON{
		Id:       uuidid.String(),
		Name:     user.Name,
		Email:    user.Email,
		Password: string(hashedpassword),
		Phone:    user.Phone,
		Age:      user.Age,
		Gender:   user.Gender,
	}

	id, err := database.InsertUser(inputuser)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  fiber.StatusInternalServerError,
			"message": "failed to insert the user",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "success registered the user",
		"data": fiber.Map{
			"userid": id,
		},
	})
}
