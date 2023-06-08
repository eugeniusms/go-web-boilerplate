package common

import "github.com/pkg/errors"

const (
	ErrorUserAlreadyExist = "error user already exist"
	ErrorInvalidEmail     = "provided email is not valid"
)

var (
	ErrUserAlreadyExist = errors.New(ErrorUserAlreadyExist)
	ErrInvalidEmail     = errors.New(ErrorInvalidEmail)
)