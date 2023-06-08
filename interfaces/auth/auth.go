package auth

import (
	"errors"
	"fmt"
	"go-web-boilerplate/application"
	"go-web-boilerplate/shared"
	"go-web-boilerplate/shared/common"
	"go-web-boilerplate/shared/dto"

	"github.com/golang-module/carbon"
	"github.com/twharmon/gouid"
	"golang.org/x/crypto/bcrypt"
)

type (
	ViewService interface {
		RegisterUser(req dto.CreateUserRequest) (dto.CreateUserResponse, error)
		Login(req dto.LoginRequest) (dto.LoginResponse, error)
		EditUser(req dto.EditUserPayload) (dto.EditUserResponse, error)
		ForgotPassword(req dto.ForgotPasswordRequest) error
		ResetPassword(req dto.ResetPasswordRequest) error
	}

	viewService struct {
		application application.Holder
		shared      shared.Holder
	}
)

func (v *viewService) RegisterUser(req dto.CreateUserRequest) (dto.CreateUserResponse, error) {
	var (
		res dto.CreateUserResponse
	)

	isUserExist, _ := v.application.AuthService.CheckUserExist(req.Email)
	if isUserExist {
		return res, errors.New("user already exist")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 14)
	if err != nil {
		return res, err
	}

	err = v.application.AuthService.CreateUser(req.TransformToUserModel(string(hashedPassword)))
	if err != nil {
		return res, err
	}

	res = dto.CreateUserResponse{
		Email:    req.Email,
		Fullname: req.Fullname,
	}

	return res, nil
}

func (v *viewService) Login(req dto.LoginRequest) (dto.LoginResponse, error) {
	var (
		res dto.LoginResponse
	)

	isUserExist, user := v.application.AuthService.CheckUserExist(req.Email)
	if !isUserExist {
		return res, errors.New("no user found for given email")
	}

	err := bcrypt.CompareHashAndPassword(
		[]byte(user.HashedPassword),
		[]byte(req.Password),
	)
	if err != nil {
		return res, errors.New("incorrect password")
	}

	token, err := common.GenerateToken(v.shared.Env.SecretKey, user)
	if err != nil {
		return res, err
	}

	res = dto.LoginResponse{
		Token: token,
	}

	return res, nil
}

func (v *viewService) EditUser(req dto.EditUserPayload) (dto.EditUserResponse, error) {
	var (
		res dto.EditUserResponse
	)

	isUserExist, user := v.application.AuthService.CheckUserExist(req.Email)
	if !isUserExist {
		return res, errors.New("no user found for given email")
	}

	user.Fullname = req.Fullname

	err := v.application.AuthService.EditUser(user)
	if err != nil {
		return res, err
	}

	res = dto.EditUserResponse{
		Email:    req.Email,
		Fullname: req.Fullname,
	}

	return res, nil
}

func (v *viewService) ForgotPassword(req dto.ForgotPasswordRequest) error {
	isUserExist, user := v.application.AuthService.CheckUserExist(req.Email)
	if !isUserExist {
		return errors.New("no user found for given email")
	}

	fpwEntry := dto.PasswordReset{
		UserID: user.ID,
		Token:  gouid.String(6, gouid.LowerCaseAlphaNum),
		Valid:  carbon.Now().AddMinutes(5).ToStdTime(),
	}

	err := v.application.AuthService.CreatePasswordReset(fpwEntry)
	if err != nil {
		return err
	}

	// ToDo: send email to user

	return nil
}

func (v *viewService) ResetPassword(req dto.ResetPasswordRequest) error {
	var (
		pw dto.PasswordReset
	)

	err := v.application.AuthService.GetResetToken(req.Token, &pw)
	if err != nil {
		return errors.New("token is invalid")
	}

	if carbon.Now().ToStdTime().After(pw.Valid) {
		return errors.New("token is expired")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 14)
	if err != nil {
		return err
	}

	fmt.Println(pw.User)
	pw.User.HashedPassword = string(hashedPassword)

	err = v.application.AuthService.EditUser(pw.User)
	if err != nil {
		return err
	}

	go v.application.AuthService.RemovePreviousPasswordResetToken(pw.UserID)

	return nil
}

func NewViewService(application application.Holder, shared shared.Holder) ViewService {
	return &viewService{
		application: application,
		shared:      shared,
	}
}