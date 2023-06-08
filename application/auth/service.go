package auth

import (
	"go-web-boilerplate/shared"
	"go-web-boilerplate/shared/common"
	"go-web-boilerplate/shared/dto"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type Service interface {
	CreateUser(newUser dto.CreateUserRequest) (dto.CreateUserResponse, error)
	GenerateToken(secretKey, id, email string) (string, error)
	Login(user dto.LoginRequest) (bool, error)
}

type service struct {
	repo   Repository
	shared shared.Holder
}

func NewAuthService(repo Repository, holder shared.Holder) (Service, error) {
	return &service{
		repo:   repo,
		shared: holder,
	}, nil
}

func (s *service) CreateUser(user dto.CreateUserRequest) (dto.CreateUserResponse, error) {
	ok := common.IsValidEmail(user.Email)
	if !ok {
		return dto.CreateUserResponse{}, errors.Wrap(common.ErrUserAlreadyExist, "email is invalid")
	}

	exist, err := s.repo.CheckUserExist(user.Email)
	if err != nil {
		return dto.CreateUserResponse{}, errors.Wrap(err, "error checking exist email")
	}

	if exist {
		return dto.CreateUserResponse{}, errors.Wrap(common.ErrUserAlreadyExist, "email already exists")
	}

	hashed, err := common.HashPassword(user.Password)
	if err != nil {
		return dto.CreateUserResponse{}, errors.Wrap(err, "error hashing password")
	}

	newUser := dto.UserModel{
		ID:             uuid.New().String(),
		Email:          user.Email,
		Fullname:       user.Fullname,
		HashedPassword: hashed,
	}

	err = s.repo.CreateUser(newUser)
	if err != nil {
		return dto.CreateUserResponse{}, errors.Wrap(err, "failed to create new user to db")
	}

	return dto.CreateUserResponse{
		ID:       newUser.ID,
		Email:    newUser.Email,
		Fullname: newUser.Fullname,
	}, nil
}

func (s *service) GenerateToken(secretKey, id, email string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", errors.Wrap(err, "failed to generate jwt token")
	}

	return signedToken, nil
}

func (s *service) Login(user dto.LoginRequest) (bool, error) {
	u, err := s.repo.GetUserByEmail(user.Email)
	if errors.Cause(err) == common.ErrUnregisteredEmail {
		return false, err
	}

	if err != nil {
		return false, errors.Wrap(err, "failed to find user by email")
	}

	ok := common.CheckPasswordHash(user.Password, u.HashedPassword)
	if !ok {
		return false, common.ErrIncorrectEmailOrPassword
	}

	return ok, nil
}