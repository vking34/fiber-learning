package auth

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/vking34/fiber-messenger/models"
	"github.com/vking34/fiber-messenger/services/user_service"
	"github.com/vking34/fiber-messenger/utils"
)

type tokenReq struct {
	Token string `json:"token" validate:"required,min=1"`
}

type facebookResp struct {
	ID      string                 `json:"id"`
	Name    string                 `json:"name"`
	Email   string                 `json:"email"`
	Picture map[string]interface{} `json:"picture"`
}

// FacebookCallback login and register with fb
func FacebookCallback(c *fiber.Ctx) error {
	var req tokenReq
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

	resp, err := http.Get("https://graph.facebook.com/v10.0/me?access_token=" + req.Token + "&fields=id,name,email,picture.type(large)")
	if err != nil {
		c.Status(500).JSON(fiber.Map{
			"status":  false,
			"message": "Can not call Facebook API",
		})
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	var fbResp facebookResp
	err = json.Unmarshal(bodyBytes, &fbResp)
	if err != nil {
		c.Status(500).JSON(fiber.Map{
			"status":  false,
			"message": "Can not parse fb response",
		})
	}

	userRecord, err := user_service.FindUserByEmail(fbResp.Email)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  false,
			"message": "Server error",
			"error":   err,
		})
	}

	if userRecord.ID != 0 {
		return GenerateToken(c, userRecord)
	}

	var user models.User
	user.Username = fbResp.Email
	user.Email = fbResp.Email
	user.Name = fbResp.Name
	user.Picture = fbResp.Picture["data"].(map[string]interface{})["url"].(string)
	user.Password, err = utils.HashPassword(fbResp.Email + fbResp.ID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": false, "message": err})
	}

	err = user_service.CreateUser(&user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": false, "message": err})
	}

	return GenerateToken(c, &user)
}
