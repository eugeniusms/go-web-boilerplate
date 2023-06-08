package healthcheck

import (
	"go-web-boilerplate/shared"
	"go-web-boilerplate/shared/dto"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type (
	Service interface {
		HttpHealthcheck(app *fiber.App) dto.Status
		DatabaseHealthcheck(db *gorm.DB) dto.Status
	}

	service struct {
		shared shared.Holder
	}
)

func (h *service) HttpHealthcheck(app *fiber.App) dto.Status {
	data := dto.HCData{
		HandlerCount: app.HandlersCount(),
	}
	return dto.Status{
		Name:   dto.HTTP,
		Status: dto.OK,
		Data:   data,
	}
}

func (h *service) DatabaseHealthcheck(db *gorm.DB) dto.Status {
	var (
		status = dto.Status{Name: dto.DB, Status: dto.OK}
	)

	pgDB, _ := db.DB()
	err := pgDB.Ping()

	if err != nil {
		h.shared.Logger.Errorf("error on db ping, %s", err.Error())
		status.Status = dto.Error
		status.Data = err.Error()
		return status
	}

	status.Data = pgDB.Stats()

	return status
}

func NewHealthcheckService(shared shared.Holder) Service {
	return &service{
		shared: shared,
	}
}