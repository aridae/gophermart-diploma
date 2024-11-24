package orderrepo

import "errors"

var (
	ErrOrderNumberUniqueConstraintViolated = errors.New("order number unique constraint violation")
)
