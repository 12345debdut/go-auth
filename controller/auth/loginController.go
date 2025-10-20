package auth

import (
	"go-auth/models/auth"

	"github.com/12345debdut/go-config"

	"github.com/gofiber/fiber/v2"
)

type LoginController struct{}

func New() *LoginController {
	return &LoginController{}
}

func (l *LoginController) Execute(c *fiber.Ctx) error {
	data := auth.LoginResponse{
		Name:    "Debdut Saha",
		Token:   "cbewuocnwcfrwcbrwfjrwoverwvnrevirweinv",
		Email:   "debdut.saha.1@gmail.com",
		Role:    "admin",
		Message: "",
	}
	configData := config.GreetFromConfig()
	data.Message = configData
	return c.Status(200).JSON(data)
}
