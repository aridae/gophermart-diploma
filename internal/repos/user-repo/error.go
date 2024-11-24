package userrepo

import "errors"

var (
	LoginUniqueConstraintViolatedError = errors.New("user login unique constraint violation")
)
