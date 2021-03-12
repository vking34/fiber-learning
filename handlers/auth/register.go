package auth

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/vking34/fiber-messenger/constants"
	"github.com/vking34/fiber-messenger/models"
	"github.com/vking34/fiber-messenger/services/user_service"
	"github.com/vking34/fiber-messenger/utils"
	"gorm.io/gorm"
)

type registerReq struct {
	Username string `json:"username" validate:"required,min=6,max=100"`
	Password string `json:"password" validate:"required,min=6"`
	Email    string `json:"email" validate:"required,email"`
	Name     string `json:"name" validate:"required,min=3,max=200"`
}

type registerResp struct {
	Username string `json:"username"`
	Name     string `json:"name"`
}

// CreateUser create user
func CreateUser(c *fiber.Ctx) error {
	var req registerReq
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": err,
		})
	}

	validationErr := utils.ValidateStruct(req)
	if validationErr != nil {
		c.Status(400).JSON(validationErr)
		return nil
	}

	result, users := user_service.FindUsersByUsernameOrEmail(req.Username, req.Email)
	if result.RowsAffected > 0 {
		if users[0].Username == req.Username {
			return c.Status(400).JSON(constants.ExistingUsername)
		}

		return c.Status(400).JSON(constants.ExistingEmail)
	}

	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return c.Status(500).JSON(fiber.Map{
			"status":  false,
			"message": "Can not create user",
			"error":   result.Error,
		})
	}

	hashedPass, err := utils.HashPassword(req.Password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": false, "message": err})
	}

	var user models.User
	user.Username = req.Username
	user.Password = hashedPass
	user.Name = req.Name
	user.Email = req.Email
	if err := user_service.CreateUser(&user); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  false,
			"message": "Can not create user",
			"error":   err,
		})
	}

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Created user",
		"data":    registerResp{user.Username, user.Name},
	})
}
