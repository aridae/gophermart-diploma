package domainerrors

import "fmt"

const (
	UnauthorizedErrorCode = iota + 1
	LoginAlreadyTakenErrorCode
	InvalidOrderNumberErrorCode
	OrderNumberAlreadySubmittedErrorCode
	InvalidUserCredentialsErrorCode
	InsufficientFundsBalance
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

func UnauthorizedError() error {
	return DomainError{
		msg:  "Unauthorized",
		Code: UnauthorizedErrorCode,
	}
}

func InvalidUserCredentialsError() error {
	return DomainError{
		msg:  "Invalid credentials",
		Code: InvalidUserCredentialsErrorCode,
	}
}

func LoginAlreadyTakenError(login string) error {
	return DomainError{
		msg:  fmt.Sprintf("Login %s already taken", login),
		Code: LoginAlreadyTakenErrorCode,
	}
}

func InvalidOrderNumberError(number string, reason string) error {
	return DomainError{
		msg:  fmt.Sprintf("Order number %s is invald: %s", number, reason),
		Code: InvalidOrderNumberErrorCode,
	}
}

func OrderNumberAlreadySubmittedError(number string) error {
	return DomainError{
		msg:  fmt.Sprintf("Order number %s has already been submitted by another user", number),
		Code: OrderNumberAlreadySubmittedErrorCode,
	}
}

func InsufficientFundsBalanceError(balance float32, withdrawal float32) error {
	return DomainError{
		msg:  fmt.Sprintf("Insufficient funds balance %.2f, withrawal request for the sum %.2f terminated", balance, withdrawal),
		Code: InsufficientFundsBalance,
	}
}

func NoAccessToOrderError(number string) error {
	return DomainError{
		msg:  fmt.Sprintf("No access: order with number %s is owned by another user", number),
		Code: NoAccessToOrder,
	}
}

func OrderNotFoundError(number string) error {
	return DomainError{
		msg:  fmt.Sprintf("Order with number %s not found", number),
		Code: OrderNotFoundErrorCode,
	}
}
