package httpapi

import (
	"encoding/json"
	"errors"
	"net/http"

	domainerrors "github.com/aridae/gophermart-diploma/internal/model/domain-errors"
)

func mapDomainErrorToHTTPStatusCode(err error) (int, string) {
	if domerror := new(domainerrors.DomainError); errors.As(err, domerror) {
		switch domerror.Code {
		case domainerrors.UnauthorizedErrorCode:
			return http.StatusUnauthorized, domerror.Error()
		case domainerrors.LoginAlreadyTakenErrorCode:
			return http.StatusConflict, domerror.Error()
		case domainerrors.InvalidOrderNumberErrorCode:
			return http.StatusUnprocessableEntity, domerror.Error()
		case domainerrors.OrderNumberAlreadySubmittedErrorCode:
			return http.StatusConflict, domerror.Error()
		case domainerrors.InvalidUserCredentialsErrorCode:
			return http.StatusUnauthorized, domerror.Error()
		case domainerrors.InsufficientOrderAccrual:
			return http.StatusPaymentRequired, domerror.Error()
		case domainerrors.NoAccessToOrder:
			return http.StatusUnprocessableEntity, domerror.Error()
		case domainerrors.OrderNotFoundErrorCode:
			return http.StatusUnprocessableEntity, domerror.Error()
		}
	}

	return http.StatusInternalServerError, err.Error()
}

func mustPresentJSONError(err error, w http.ResponseWriter) {
	code, msg := mapDomainErrorToHTTPStatusCode(err)

	jsonErr := struct {
		Message string `json:"message"`
	}{
		Message: msg,
	}

	payload, _ := json.Marshal(jsonErr)

	_, _ = w.Write(payload)
	w.WriteHeader(code)
}

func mustPresentJSONErrorWithCode(err error, w http.ResponseWriter, code int) {
	jsonErr := struct {
		Message string `json:"message"`
	}{
		Message: err.Error(),
	}

	payload, _ := json.Marshal(jsonErr)

	_, _ = w.Write(payload)
	w.WriteHeader(code)
}
