package userrepo

import "errors"

var (
	ErrLoginUniqueConstraintViolated = errors.New("user login unique constraint violation")
)
