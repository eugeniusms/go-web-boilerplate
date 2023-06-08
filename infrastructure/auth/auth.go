package auth

import (
	"go-web-boilerplate/interfaces"
	"go-web-boilerplate/shared"
	"go-web-boilerplate/shared/dto"

	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	Interfaces interfaces.Holder
	Shared     shared.Holder
}

func (c *Controller) register(ctx *fiber.Ctx) error {
	var req dto.CreateUserRequest
	if err := ctx.BodyParser(&req); err != nil {
		return err
	}

	data, err := c.Interfaces.AuthViewService.RegisterUser(req)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}
	return ctx.Status(fiber.StatusOK).JSON(data)
}

func (c *Controller) login(ctx *fiber.Ctx) error {
	var req dto.LoginRequest
	if err := ctx.BodyParser(&req); err != nil {
		return err
	}

	data, err := c.Interfaces.AuthViewService.Login(req)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(data)
}

func (c *Controller) Routes(app *fiber.App) {
	auth := app.Group("/auth")
	auth.Post("/register", c.register)
	auth.Post("/login", c.login)
}

func NewController(interfaces interfaces.Holder, shared shared.Holder) Controller {
	return Controller{
		Interfaces: interfaces,
		Shared:     shared,
	}
}