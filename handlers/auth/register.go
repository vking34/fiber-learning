package auth

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/vking34/fiber-messenger/db"
	"github.com/vking34/fiber-messenger/models"
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

	var user models.User
	result := db.DB.Where(&models.User{Username: req.Username}).First(&user)

	if result.RowsAffected > 0 {
		return c.Status(400).JSON(fiber.Map{
			"status":  false,
			"message": "Existing username",
		})
	}

	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
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

	user.Username = req.Username
	user.Password = hashedPass
	user.Name = req.Name

	if err := db.DB.Create(&user).Error; err != nil {
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
