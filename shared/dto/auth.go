package dto

import (
	"time"

	"gorm.io/gorm"
)

type (
	User struct {
		ID             uint   `gorm:"primaryKey;autoIncrement"`
		Email          string `gorm:"column:email;unique;not null"`
		Fullname       string `gorm:"column:fullname"`
		HashedPassword string `gorm:"column:hashed_password"`
		CreatedAt      time.Time
		UpdatedAt      time.Time
		DeletedAt      gorm.DeletedAt
	}

	PasswordReset struct {
		ID        uint `gorm:"primaryKey;autoIncrement"`
		UserID    uint
		User      User   `gorm:"onDelete:CASCADE"`
		Token     string `gorm:"column:token"`
		CreatedAt time.Time
		Valid     time.Time
	}

	// CreateUserRequest CreateUserRequest
	CreateUserRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
		Fullname string `json:"fullname" validate:"required"`
	}

	// CreateUserResponse CreateUserResponse
	CreateUserResponse struct {
		Email    string `json:"email"`
		Fullname string `json:"fullname"`
	}

	// LoginRequest LoginRequest
	LoginRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}

	// LoginResponse LoginResponse
	LoginResponse struct {
		Token string `json:"token"`
	}

	// EditUserRequest EditUserRequest

	EditUserRequest struct {
		Fullname string `json:"fullname" validate:"required"`
	}

	EditUserPayload struct {
		Fullname string
		ID       float64
		Email    string
	}

	// EditUserResponse EditUserResponse
	EditUserResponse struct {
		Email    string `json:"email"`
		Fullname string `json:"fullname"`
	}

	// ForgotPasswordRequest ForgotPasswordRequest
	ForgotPasswordRequest struct {
		Email string `json:"email" validate:"required,email"`
	}

	// ResetPasswordRequest ResetPasswordRequest
	ResetPasswordRequest struct {
		Password string `json:"password" validate:"required"`
		Token    string `json:"token" validate:"required"`
	}
)

func (r *CreateUserRequest) TransformToUserModel(hp string) User {
	return User{
		Email:          r.Email,
		Fullname:       r.Fullname,
		HashedPassword: hp,
	}
}