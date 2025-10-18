package auth

import (
	login "go-auth/controller/auth"

	"github.com/gofiber/fiber/v2"
)

type AuthRoute struct {
	app *fiber.App
}

func New(app *fiber.App) *AuthRoute {
	return &AuthRoute{app: app}
}
func (r *AuthRoute) InitRoutes() {
	r.app.Get("/login", func(c *fiber.Ctx) error {
		return login.New().Execute(c)
	})
}
