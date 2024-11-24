package submitorder

import (
	"context"
	"fmt"
	"time"

	"github.com/aridae/gophermart-diploma/internal/auth/authctx"
	"github.com/aridae/gophermart-diploma/internal/model"
	domainerrors "github.com/aridae/gophermart-diploma/internal/model/domain-errors"
)

type transactionManager interface {
	Do(ctx context.Context, fn func(ctx context.Context) error) (err error)
}

type ordersRepository interface {
	CreateOrder(ctx context.Context, orderSubmit model.OrderSubmit, now time.Time) error
	GetByNumbers(ctx context.Context, orderNumbers []string) ([]model.Order, error)
}

type Handler struct {
	transactionManager transactionManager
	ordersRepository   ordersRepository
	now                func() time.Time
}

func NewHandler(
	transactionManager transactionManager,
	ordersRepository ordersRepository,
) *Handler {
	return &Handler{
		transactionManager: transactionManager,
		ordersRepository:   ordersRepository,
		now:                time.Now().UTC,
	}
}

type Request struct {
	OrderNumber string
}

type ResponseCode int

const (
	OrderNumberAlreadyLoadedByThisOwner ResponseCode = iota + 1
	OrderNumberAccepted
)

type Response struct {
	Code ResponseCode
}

func (h *Handler) Handle(ctx context.Context, req Request) (Response, error) {
	user, authorized := authctx.GetUserFromContext(ctx)
	if !authorized {
		return Response{}, domainerrors.ErrUnauthorized()
	}

	err := validateOrderNumber(req.OrderNumber)
	if err != nil {
		return Response{}, domainerrors.ErrInvalidOrderNumber(req.OrderNumber, err.Error())
	}

	now := h.now()

	var resp Response
	err = h.transactionManager.Do(ctx, func(ctx context.Context) error {
		orders, txErr := h.ordersRepository.GetByNumbers(ctx, []string{req.OrderNumber})
		if txErr != nil {
			return fmt.Errorf("ordersRepository.GetByNumbers: %w", txErr)
		}

		if len(orders) > 0 {
			if orders[0].Owner.Login != user.Login {
				return domainerrors.ErrOrderNumberAlreadySubmitted(req.OrderNumber)
			}
			resp.Code = OrderNumberAlreadyLoadedByThisOwner
			return nil
		}

		txErr = h.ordersRepository.CreateOrder(ctx, model.OrderSubmit{
			Number: req.OrderNumber,
			Owner:  user,
			Status: model.OrderStatusNew,
		}, now)
		if txErr != nil {
			return fmt.Errorf("ordersRepository.CreateOrder: %w", txErr)
		}

		resp.Code = OrderNumberAccepted
		return nil
	})
	if err != nil {
		return Response{}, fmt.Errorf("transaction failed with err: %w", err)
	}

	return resp, nil
}

func validateOrderNumber(number string) error {
	if !isLuhnValid(number) {
		return fmt.Errorf("invalid order number: %s", number)
	}

	return nil
}

func isLuhnValid(number string) bool {
	sum, length := 0, len(number)
	if length < 2 {
		return false
	}
	for index, num := range number {
		dig := int(num - '0')
		if length%2 == index%2 {
			dig *= 2
			if dig > 9 {
				dig = dig%10 + dig/10
			}
		}
		sum += dig
	}
	return sum%10 == 0
}
