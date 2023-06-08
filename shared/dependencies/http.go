package depedencies

import (
	"github.com/goccy/go-json"

	_ "go-web-boilerplate/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/swagger"
	"github.com/sirupsen/logrus"
)

func NewHttp(log *logrus.Logger) *fiber.App {
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	app.Use(logger.New(logger.Config{
		Format: "[${time}]:[${ip}] ${status} - ${latency} ${method} ${path}\n",
	}))

	app.Use(cors.New(cors.ConfigDefault))

	app.Get("/metrics", monitor.New(monitor.Config{Title: "Go Web Boilerplate Server Metrics"}))

	app.Get("/swagger/*", swagger.HandlerDefault)

	log.Infoln("fiber http receiver initialized")

	return app
}