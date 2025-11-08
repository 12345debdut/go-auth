package main

import (
	"fmt"
	webapp "go-auth/app"
	"go-auth/db"
	"go-auth/routes"
	"go-auth/utils"
	"log"
)

func main() {
	fmt.Println(">>> Starting main()")

	dbGorm, _, disconnect := db.New().InitDb()
	defer disconnect()
	log.Println(">>> Database connected OK")

	app := webapp.Instance()
	log.Println(">>> Webapp instance created")

	routes.InitRoutes(app, dbGorm)
	log.Println(">>> Routes initialized")

	port := utils.GetEnv("APP_PORT", "3000")
	log.Println(">>> Using port:", port)

	err := app.Listen(":" + port)
	if err != nil {
		log.Fatalf(">>> App failed to start: %v", err)
	} else {
		log.Println(">>> App started successfully")
	}
}
