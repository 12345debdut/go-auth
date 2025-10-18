package app

import (
	"sync"

	"github.com/gofiber/fiber/v2"
)

var (
	instance *fiber.App
	once     sync.Once
)

// Instance returns the singleton Fiber app for the whole project.
// It is initialized on first use. Any package importing and calling
// Instance() will receive the same *fiber.App pointer.
func Instance() *fiber.App {
	once.Do(func() {
		instance = fiber.New()
	})
	return instance
}
