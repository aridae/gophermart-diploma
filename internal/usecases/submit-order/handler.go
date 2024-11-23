package submitorder

import (
	"context"
	"errors"
	"fmt"
	"github.com/ShiraazMoollatjie/goluhn"
	"github.com/aridae/gophermart-diploma/internal/auth/authctx"
	"github.com/aridae/gophermart-diploma/internal/model"
	domainerrors "github.com/aridae/gophermart-diploma/internal/model/domain-errors"
	orderdb "github.com/aridae/gophermart-diploma/internal/repo/order-db"
	"time"
)

type ordersRepository interface {
	CreateOrder(ctx context.Context, orderSubmit model.OrderSubmit, now time.Time) error
}

type Handler struct {
	ordersRepository ordersRepository
	now              func() time.Time
}

func NewHandler(ordersRepository ordersRepository) *Handler {
	return &Handler{
		ordersRepository: ordersRepository,
		now:              time.Now().UTC,
	}
}

type Request struct {
	OrderNumber string
}

func (h *Handler) Handle(ctx context.Context, req Request) error {
	user, authorized := authctx.GetUserFromContext(ctx)
	if !authorized {
		return domainerrors.UnauthorizedError()
	}

	err := validateOrderNumber(req.OrderNumber)
	if err != nil {
		return domainerrors.InvalidOrderNumberError(req.OrderNumber, err.Error())
	}

	now := h.now()
	orderSubmit := model.OrderSubmit{
		Number: req.OrderNumber,
		Owner:  user,
		Status: model.OrderStatusNew,
	}

	err = h.ordersRepository.CreateOrder(ctx, orderSubmit, now)
	if err != nil {
		if errors.Is(err, orderdb.OrderNumberUniqueConstraintViolatedError) {
			return domainerrors.OrderNumberAlreadySubmittedError(req.OrderNumber)
		}

		return fmt.Errorf("ordersRepository.CreateOrder: %w", err)
	}

	return nil
}

func validateOrderNumber(number string) error {
	return goluhn.Validate(number)
}
