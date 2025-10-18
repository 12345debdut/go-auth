package controller

import "github.com/gofiber/fiber/v2"

type Controller interface {
	Execute(c *fiber.Ctx) error
}
