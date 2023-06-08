package auth

import (
	"go-web-boilerplate/application"
	"go-web-boilerplate/shared"
	"go-web-boilerplate/shared/dto"
)

type ViewService interface {
	RegisterUser(req dto.CreateUserRequest) (dto.CreateUserResponse, error)
	Login(req dto.LoginRequest) (dto.LoginResponse, error)
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

	token, err := v.application.AuthService.GenerateToken(v.shared.Env.SecretKey, resp.ID, resp.Email)
	if err != nil {
		return dto.CreateUserResponse{}, err
	}

	resp.Token = token
	return resp, nil
}

func (v *viewService) Login(req dto.LoginRequest) (dto.LoginResponse, error) {
	user, err := v.application.AuthService.Login(req)
	if err != nil {
		return dto.LoginResponse{}, err
	}

	token, err := v.application.AuthService.GenerateToken(v.shared.Env.SecretKey, user.ID, user.Email)
	if err != nil {
		return dto.LoginResponse{}, err
	}

	return dto.LoginResponse{
		Token: token,
	}, nil
}