package getwithdrawalshistory

import (
	"context"
	"fmt"

	"github.com/aridae/gophermart-diploma/internal/auth/authctx"
	"github.com/aridae/gophermart-diploma/internal/model"
	domainerrors "github.com/aridae/gophermart-diploma/internal/model/domain-errors"
)

type withdrawalsRepository interface {
	GetByActor(ctx context.Context, actor model.User) ([]model.WithdrawalLog, error)
}

type Handler struct {
	withdrawalsRepository withdrawalsRepository
}

func NewHandler(withdrawalsRepository withdrawalsRepository) *Handler {
	return &Handler{withdrawalsRepository: withdrawalsRepository}
}

func (h *Handler) Handle(ctx context.Context) ([]model.WithdrawalLog, error) {
	user, authorized := authctx.GetUserFromContext(ctx)
	if !authorized {
		return nil, domainerrors.ErrUnauthorized()
	}

	withdrawalsLogs, err := h.withdrawalsRepository.GetByActor(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("withdrawalsRepository.GetByActor: %w", err)
	}

	return withdrawalsLogs, nil
}
