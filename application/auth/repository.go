package auth

import (
	"go-web-boilerplate/shared"
	"go-web-boilerplate/shared/dto"

	"gorm.io/gorm"
)

type Repository interface {
	CheckUserExist(email string) (bool, error)
	CreateUser(newUser dto.UserModel) error
}

type repository struct {
	db *gorm.DB
}

func NewAuthRepostiory(holder shared.Holder) (Repository, error) {
	return &repository{
		db: holder.DB,
	}, nil
}

func (r *repository) CheckUserExist(email string) (bool, error) {
	var user dto.UserModel
	err := r.db.First(&user, "email = ?", email).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}

	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *repository) CreateUser(newUser dto.UserModel) error {
	err := r.db.Create(&newUser).Error
	if err != nil {
		return err
	}
	return nil
}