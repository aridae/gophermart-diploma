package userdb

import "errors"

var (
	LoginUniqueConstraintViolatedError = errors.New("user login unique constraint violation")
)
