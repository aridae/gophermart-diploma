package httpapi

import (
	getbalance "github.com/aridae/gophermart-diploma/internal/usecases/get-balance"
	getorders "github.com/aridae/gophermart-diploma/internal/usecases/get-orders"
	getwithdrawalshistory "github.com/aridae/gophermart-diploma/internal/usecases/get-withdrawals-history"
	loginuser "github.com/aridae/gophermart-diploma/internal/usecases/login-user"
	registeruser "github.com/aridae/gophermart-diploma/internal/usecases/register-user"
	requestwithdrawal "github.com/aridae/gophermart-diploma/internal/usecases/request-withdrawal"
	submitorder "github.com/aridae/gophermart-diploma/internal/usecases/submit-order"
)

type ApiService struct {
	getBalanceHandler            *getbalance.Handler
	getOrdersHandler             *getorders.Handler
	getWithdrawalsHistoryHandler *getwithdrawalshistory.Handler
	loginUserHandler             *loginuser.Handler
	registerUserHandler          *registeruser.Handler
	requestWithdrawalHandler     *requestwithdrawal.Handler
	submitOrderHandler           *submitorder.Handler
}

func NewAPIService(
	getBalanceHandler *getbalance.Handler,
	getOrdersHandler *getorders.Handler,
	getWithdrawalsHistoryHandler *getwithdrawalshistory.Handler,
	loginUserHandler *loginuser.Handler,
	registerUserHandler *registeruser.Handler,
	requestWithdrawalHandler *requestwithdrawal.Handler,
	submitOrderHandler *submitorder.Handler,
) *ApiService {
	api := &ApiService{
		getBalanceHandler:            getBalanceHandler,
		getOrdersHandler:             getOrdersHandler,
		getWithdrawalsHistoryHandler: getWithdrawalsHistoryHandler,
		loginUserHandler:             loginUserHandler,
		registerUserHandler:          registerUserHandler,
		requestWithdrawalHandler:     requestWithdrawalHandler,
		submitOrderHandler:           submitOrderHandler,
	}

	return api
}
