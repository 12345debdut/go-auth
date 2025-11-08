package routes

import (
	"go-auth/routes/auth"
	"go-auth/routes/user"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Route interface {
	InitRoutes()
}

func InitRoutes(app *fiber.App, goOrm *gorm.DB) {
	var routeSlice []Route
	routes := append(routeSlice, auth.New(app, goOrm), user.New(app, goOrm))
	for _, route := range routes {
		route.InitRoutes()
	}
}
