package healthcheck

import (
	"go-web-boilerplate/interfaces"
	"go-web-boilerplate/shared"

	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	Interfaces interfaces.Holder
	Shared     shared.Holder
}

func (c *Controller) Routes(app *fiber.App) {
	app.Get("/healthcheck", c.healthcheck)
}

// All godoc
// @Tags Healthcheck
// @Summary Check system status
// @Description Put all mandatory parameter
// @Accept  json
// @Produce  json
// @Success 200 {array} dto.Status
// @Failure 200 {array} dto.Status
// @Router /healthcheck [get]
func (c *Controller) healthcheck(ctx *fiber.Ctx) error {
	c.Shared.Logger.Println("checking server status")
	data, _ := c.Interfaces.HealthcheckViewService.SystemHealthcheck()
	return ctx.Status(fiber.StatusOK).JSON(data)
}

func NewController(interfaces interfaces.Holder, shared shared.Holder) Controller {
	return Controller{
		Interfaces: interfaces,
		Shared:     shared,
	}
}