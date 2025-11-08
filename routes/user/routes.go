package user

import (
	"go-auth/controller/users"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserRoute struct {
	app *fiber.App
	db  *gorm.DB
}

func New(app *fiber.App, db *gorm.DB) *UserRoute {
	return &UserRoute{app: app, db: db}
}
func (r *UserRoute) InitRoutes() {
	r.app.Get("/user", func(c *fiber.Ctx) error {
		return users.New(r.db).Execute(c)
	})
}
