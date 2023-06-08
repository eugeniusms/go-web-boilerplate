package dto

import (
	"time"

	"gorm.io/gorm"
)

type CreateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Fullname string `json:"fullname"`
}

type CreateUserResponse struct {
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
	Token    string `json:"token"`
}

type UserModel struct {
	ID             string    `gorm:"column:id;unique;not null"`
	Email          string    `gorm:"column:email;unique;not null"`
	Fullname       string    `gorm:"column:fullname"`
	HashedPassword string    `gorm:"column:hashed_password"`
	CreatedAt      time.Time `gorm:"column:created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at"`
	DeletedAt      gorm.DeletedAt
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}