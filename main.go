package main

import (
	"go-web-boilerplate/di"
	"go-web-boilerplate/infrastructure"
	"log"

	"go-web-boilerplate/shared/config"

	"github.com/gofiber/fiber/v2"
)

// @title Go Web Boilerplate
// @version 1.0
// @description API definition for Go Web Boilerplate
// @host localhost:8000
// @BasePath /
func main() {
	container := di.Container

	err := container.Invoke(func(http *fiber.App, env *config.EnvConfig, holder infrastructure.Holder) error {	
		infrastructure.Routes(http, holder)

		err := http.Listen(":" + env.PORT)
		if err != nil {
			return err
		}
		
		return nil
	})

	if err != nil {
		log.Fatalf("error when starting http server: %s", err.Error())
	}
}