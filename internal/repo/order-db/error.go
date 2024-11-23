package orderdb

import "errors"

var (
	OrderNumberUniqueConstraintViolatedError = errors.New("order number unique constraint violation")
)
