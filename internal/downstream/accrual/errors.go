package accrual

import (
	"errors"
)

var (
	ErrOrderNotFound     = errors.New("order not found")
	ErrResourceExhausted = errors.New("resource exhausted")
)
