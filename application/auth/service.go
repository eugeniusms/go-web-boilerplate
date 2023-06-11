package auth

import (
	"go-web-boilerplate/shared"
	"go-web-boilerplate/shared/dto"
)

type (
	Service interface {
		CheckUserExist(email string) (bool, dto.User)
		CreateUser(user dto.User) error
		EditUser(user dto.User) error
		CreatePasswordReset(pw dto.PasswordReset) error
		GetResetToken(token string, pw *dto.PasswordReset) error
		RemovePreviousPasswordResetToken(id uint)
		GetUserContext(id uint) dto.User
		ListUser(preload string) dto.UserSlice
	}

	service struct {
		shared shared.Holder
	}
)

func (s *service) CheckUserExist(email string) (bool, dto.User) {
	var user dto.User

	err := s.shared.DB.First(&user, "email = ?", email).Error

	return err == nil, user
}

func (s *service) CreateUser(user dto.User) error {
	err := s.shared.DB.Create(&user).Error
	return err
}

func (s *service) EditUser(user dto.User) error {
	err := s.shared.DB.Save(&user).Error
	return err
}

func (s *service) CreatePasswordReset(pw dto.PasswordReset) error {
	err := s.shared.DB.Create(&pw).Error
	return err
}

func (s *service) GetResetToken(token string, pw *dto.PasswordReset) error {
	err := s.shared.DB.Preload("User").First(pw, "token = ?", token).Error
	return err
}

func (s *service) RemovePreviousPasswordResetToken(id uint) {
	var pw dto.PasswordReset
	s.shared.DB.Where("user_id = ?", id).Delete(&pw)
}

func (s *service) GetUserContext(id uint) dto.User {
	var user dto.User
	s.shared.DB.Where("id = ?", id).First(&user)
	return user
}

func (s *service) ListUser(preload string) dto.UserSlice {
	var users []dto.User

	s.shared.DB.Preload(preload).Find(&users)

	return users
}

func NewAuthService(holder shared.Holder) (Service, error) {
	return &service{
		shared: holder,
	}, nil
}