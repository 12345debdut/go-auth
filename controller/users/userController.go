package users

import (
	"go-auth/models/users"

	"github.com/gofiber/fiber/v2"
)

type GetUserController struct{}

func New() *GetUserController {
	return &GetUserController{}
}

func (u *GetUserController) Execute(c *fiber.Ctx) error {
	model := users.UserResponseModel{
		Name:  "Debdut Saha",
		Email: "debdut.saha.1@gmail.com",
		Role:  "admin",
	}
	return c.Status(200).JSON(model)
}
