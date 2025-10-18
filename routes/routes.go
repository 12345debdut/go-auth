package routes

import (
	"go-auth/routes/auth"
	"go-auth/routes/user"

	"github.com/gofiber/fiber/v2"
)

type Route interface {
	InitRoutes()
}

func InitRoutes(app *fiber.App) {
	var routeSlice []Route
	routes := append(routeSlice, auth.New(app), user.New(app))
	for _, route := range routes {
		route.InitRoutes()
	}
}
