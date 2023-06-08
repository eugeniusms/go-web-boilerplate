package di

import (
	"go-web-boilerplate/application"
	"go-web-boilerplate/infrastructure"
	"go-web-boilerplate/interfaces"
	"go-web-boilerplate/shared"
	"log"

	"go.uber.org/dig"
)

var Container = dig.New()

func init() {
	if err := shared.Register(Container); err != nil {
		log.Fatal(err.Error())
	}

	if err := application.Register(Container); err != nil {
		log.Fatal(err.Error())
	}

	if err := interfaces.Register(Container); err != nil {
		log.Fatal(err.Error())
	}

	if err := infrastructure.Register(Container); err != nil {
		log.Fatal(err.Error())
	}
}