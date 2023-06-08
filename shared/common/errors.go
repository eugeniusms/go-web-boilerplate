package common

import "github.com/pkg/errors"

const (
	ErrorUserAlreadyExist         = "error user already exist"
	ErrorInvalidEmail             = "provided email is not valid"
	ErrorUnregisteredEmail        = "email is not registered"
	ErrorIncorrectEmailOrPassword = "incorrect email or password"
)

var (
	ErrUserAlreadyExist         = errors.New(ErrorUserAlreadyExist)
	ErrInvalidEmail             = errors.New(ErrorInvalidEmail)
	ErrUnregisteredEmail        = errors.New(ErrorUnregisteredEmail)
	ErrIncorrectEmailOrPassword = errors.New(ErrorIncorrectEmailOrPassword)
)