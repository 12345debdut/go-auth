package user

import (
	"go-auth/controller/users"

	"github.com/gofiber/fiber/v2"
)

type UserRoute struct {
	app *fiber.App
}

func New(app *fiber.App) *UserRoute {
	return &UserRoute{app: app}
}
func (r *UserRoute) InitRoutes() {
	r.app.Get("/user", func(c *fiber.Ctx) error {
		return users.New().Execute(c)
	})
}
