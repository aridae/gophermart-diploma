package domainerrors

import "fmt"

const (
	UnauthorizedErrorCode = iota + 1
	LoginAlreadyTakenErrorCode
	InvalidOrderNumberErrorCode
	OrderNumberAlreadySubmittedErrorCode
	InvalidUserCredentialsErrorCode
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

func InvalidOrderNumberError(number string) error {
	return DomainError{
		msg:  fmt.Sprintf("Order number %s is invald", number),
		Code: InvalidOrderNumberErrorCode,
	}
}

func OrderNumberAlreadySubmittedError(number string) error {
	return DomainError{
		msg:  fmt.Sprintf("Order number %s has already been submitted by another user", number),
		Code: OrderNumberAlreadySubmittedErrorCode,
	}
}
