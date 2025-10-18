package main

import (
	webapp "go-auth/app"
	"go-auth/routes"
	"log"
	"os"
)

func getEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func main() {
	app := webapp.Instance()
	routes.InitRoutes(app)
	port := getEnv("PORT", "3000")
	if err := app.Listen(":" + port); err != nil {
		log.Fatal(err)
	}
}
