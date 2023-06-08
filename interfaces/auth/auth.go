package auth

import (
	"go-web-boilerplate/application"
	"go-web-boilerplate/shared"
	"go-web-boilerplate/shared/dto"
)

type ViewService interface {
	RegisterUser(req dto.CreateUserRequest) (dto.CreateUserResponse, error)
}

type viewService struct {
	application application.Holder
	shared      shared.Holder
}

func NewViewService(application application.Holder, shared shared.Holder) ViewService {
	return &viewService{
		application: application,
		shared:      shared,
	}
}

func (v *viewService) RegisterUser(req dto.CreateUserRequest) (dto.CreateUserResponse, error) {
	resp, err := v.application.AuthService.CreateUser(req)
	if err != nil {
		return dto.CreateUserResponse{}, err
	}

	token, err := v.application.AuthService.GenerateToken(v.shared.Env.SecretKey, resp.Email)
	if err != nil {
		return dto.CreateUserResponse{}, err
	}

	resp.Token = token
	return resp, nil
}