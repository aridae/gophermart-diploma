package domainerrors

import (
	"fmt"

	"github.com/aridae/gophermart-diploma/internal/model"
)

const (
	UnauthorizedErrorCode = iota + 1
	LoginAlreadyTakenErrorCode
	InvalidOrderNumberErrorCode
	OrderNumberAlreadySubmittedErrorCode
	InvalidUserCredentialsErrorCode
	InsufficientOrderAccrual
	NoAccessToOrder
	OrderNotFoundErrorCode
)

type DomainError struct {
	msg  string
	Code int64
}

func (de DomainError) Error() string {
	return de.msg
}

func ErrUnauthorized() error {
	return DomainError{
		msg:  "Unauthorized",
		Code: UnauthorizedErrorCode,
	}
}

func ErrInvalidUserCredentials() error {
	return DomainError{
		msg:  "Invalid credentials",
		Code: InvalidUserCredentialsErrorCode,
	}
}

func ErrLoginAlreadyTaken(login string) error {
	return DomainError{
		msg:  fmt.Sprintf("Login %s already taken", login),
		Code: LoginAlreadyTakenErrorCode,
	}
}

func ErrInvalidOrderNumber(number string, reason string) error {
	return DomainError{
		msg:  fmt.Sprintf("Order number %s is invald: %s", number, reason),
		Code: InvalidOrderNumberErrorCode,
	}
}

func ErrOrderNumberAlreadySubmitted(number string) error {
	return DomainError{
		msg:  fmt.Sprintf("Order number %s has already been submitted by another user", number),
		Code: OrderNumberAlreadySubmittedErrorCode,
	}
}

func ErrInsufficientOrderAccrual(number string, accrual model.Money, withdrawal model.Money) error {
	return DomainError{
		msg:  fmt.Sprintf("Insufficient accrual %.2f for order number %s, withrawal request for the sum %.2f terminated", accrual.Float32(), number, withdrawal.Float32()),
		Code: InsufficientOrderAccrual,
	}
}
