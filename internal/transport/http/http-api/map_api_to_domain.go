package httpapi

import (
	"github.com/aridae/gophermart-diploma/internal/model"
	oapispec "github.com/aridae/gophermart-diploma/internal/transport/http/http-api/oapi-spec"
	loginuser "github.com/aridae/gophermart-diploma/internal/usecases/login-user"
	registeruser "github.com/aridae/gophermart-diploma/internal/usecases/register-user"
	requestwithdrawal "github.com/aridae/gophermart-diploma/internal/usecases/request-withdrawal"
)

func mapAPIToDomainLoginRequest(r oapispec.PostUserLoginJSONBody) loginuser.Request {
	return loginuser.Request{
		Login:    r.Login,
		Password: r.Password,
	}
}

func mapAPIToDomainRegisterRequest(r oapispec.PostUserRegisterJSONBody) registeruser.Request {
	return registeruser.Request{
		Login:    r.Login,
		Password: r.Password,
	}
}

func mapAPIToDomainWithdrawalRequest(r oapispec.PostUserBalanceJSONBody) requestwithdrawal.Request {
	return requestwithdrawal.Request{
		OrderNumber: r.Order,
		Sum:         model.NewMoney(r.Sum),
	}
}
