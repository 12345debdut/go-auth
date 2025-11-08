package users

import (
	"go-auth/models/users"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type GetUserController struct {
	db *gorm.DB
}

func New(db *gorm.DB) *GetUserController {
	return &GetUserController{db: db}
}

func (u *GetUserController) Execute(c *fiber.Ctx) error {
	model := users.UserResponseModel{
		Name:  "Debdut Saha",
		Email: "debdut.saha.1@gmail.com",
		Role:  "admin",
	}
	return c.Status(200).JSON(model)
}
