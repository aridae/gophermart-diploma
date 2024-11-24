package accrual

import (
	"errors"
)

var (
	OrderNotFoundErr     = errors.New("order not found")
	ResourceExhaustedErr = errors.New("resource exhausted")
)
