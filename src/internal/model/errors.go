package model

import "github.com/pkg/errors"

var (
	ErrBadRequest    = errors.New("bad request")
	ErrUserNotExists = errors.New("user not exists")
	ErrForbidden     = errors.New("access denied")
	ErrNotFound      = errors.New("not found")
)
