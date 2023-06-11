package dto

import (
	"time"

	"gorm.io/gorm"
)

const (
	Google     UserLoginType = "Google"
	Credential UserLoginType = "Credential"
)

type (
	UserLoginType string
	User          struct {
		ID                uint          `gorm:"primaryKey;autoIncrement"`
		Email             string        `gorm:"column:email;unique;not null"`
		Fullname          string        `gorm:"column:fullname"`
		HashedPassword    string        `gorm:"column:hashed_password"`
		Type              UserLoginType `gorm:"column:type;default:Credential"`
		CreatedAt         time.Time
		UpdatedAt         time.Time
		DeletedAt         gorm.DeletedAt
	}

	UserSlice []User

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

	// GoogleLoginRequest GoogleLoginRequest
	GoogleLoginRequest struct {
		Token string `json:"token" validate:"required"`
	}

	GoogleData struct {
		Sub           string `json:"sub"`
		Name          string `json:"name"`
		GivenName     string `json:"given_name"`
		FamilyName    string `json:"family_name"`
		Picture       string `json:"picture"`
		Email         string `json:"email"`
		EmailVerified string `json:"email_verified"`
		Locale        string `json:"locale"`
	}

	// EditUserRequest EditUserRequest

	EditUserRequest struct {
		Fullname string `json:"fullname" validate:"omitempty"`
		Password string `json:"password" validate:"omitempty"`
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

func (g GoogleData) ToUser() User {
	return User{
		Fullname:          g.Name,
		Email:             g.Email,
		Type:              Google,
	}
}

func (r *CreateUserRequest) TransformToUserModel(hp string) User {
	return User{
		Email:             r.Email,
		Fullname:          r.Fullname,
		HashedPassword:    hp,
	}
}