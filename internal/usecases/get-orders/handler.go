package getorders

import (
	"context"
	"fmt"

	"github.com/aridae/gophermart-diploma/internal/auth/authctx"
	"github.com/aridae/gophermart-diploma/internal/model"
	domainerrors "github.com/aridae/gophermart-diploma/internal/model/domain-errors"
)

type ordersRepository interface {
	GetByOwner(ctx context.Context, owner model.User) ([]model.Order, error)
}

type Handler struct {
	ordersRepository ordersRepository
}

func NewHandler(ordersRepository ordersRepository) *Handler {
	return &Handler{
		ordersRepository: ordersRepository,
	}
}

func (h *Handler) Handle(ctx context.Context) ([]model.Order, error) {
	user, authorized := authctx.GetUserFromContext(ctx)
	if !authorized {
		return nil, domainerrors.UnauthorizedError()
	}

	orders, err := h.ordersRepository.GetByOwner(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("ordersRepository.GetByOwner: %w", err)
	}

	return orders, nil
}
