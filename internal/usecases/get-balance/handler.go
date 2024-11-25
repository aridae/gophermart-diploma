package getbalance

import (
	"context"
	"fmt"

	"github.com/aridae/gophermart-diploma/internal/auth/authctx"
	"github.com/aridae/gophermart-diploma/internal/model"
	domainerrors "github.com/aridae/gophermart-diploma/internal/model/domain-errors"
)

type balanceRepository interface {
	GetUserBalance(ctx context.Context, user model.User) (model.Balance, error)
}

type Handler struct {
	balanceRepository balanceRepository
}

func NewHandler(balanceRepository balanceRepository) *Handler {
	return &Handler{
		balanceRepository: balanceRepository,
	}
}

func (h *Handler) Handle(ctx context.Context) (model.Balance, error) {
	user, authorized := authctx.GetUserFromContext(ctx)
	if !authorized {
		return model.Balance{}, domainerrors.ErrUnauthorized()
	}

	balance, err := h.balanceRepository.GetUserBalance(ctx, user)
	if err != nil {
		return model.Balance{}, fmt.Errorf("balanceRepository.GetUserBalance: %w", err)

	}

	return balance, nil
}
