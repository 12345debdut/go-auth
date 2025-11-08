package auth

import (
	login "go-auth/controller/auth"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AuthRoute struct {
	app *fiber.App
	db  *gorm.DB
}

func New(app *fiber.App, goOrm *gorm.DB) *AuthRoute {
	return &AuthRoute{app: app, db: goOrm}
}
func (r *AuthRoute) InitRoutes() {
	r.app.Get("/login", func(c *fiber.Ctx) error {
		return login.New(r.db).Execute(c)
	})
}
