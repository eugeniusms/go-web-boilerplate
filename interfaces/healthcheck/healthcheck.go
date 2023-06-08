package healthcheck

import (
	"go-web-boilerplate/application"
	"go-web-boilerplate/shared"
	"go-web-boilerplate/shared/dto"
)

type (
	ViewService interface {
		SystemHealthcheck() (dto.HCStatus, error)
	}

	viewService struct {
		application application.Holder
		shared      shared.Holder
	}
)


func (v *viewService) SystemHealthcheck() (dto.HCStatus, error) {
	status := make([]dto.Status, 0)

	httpStatus := v.application.HealthcheckService.HttpHealthcheck(v.shared.Http)
	status = append(status, httpStatus)

	dbStatus := v.application.HealthcheckService.DatabaseHealthcheck(v.shared.DB)
	status = append(status, dbStatus)

	return dto.HCStatus{
		Status: status,
	}, nil
}

func NewViewService(application application.Holder, shared shared.Holder) ViewService {
	return &viewService{
		application: application,
		shared:      shared,
	}
}