package getbalance

import (
	"context"
	"github.com/aridae/gophermart-diploma/internal/auth/authctx"
	"github.com/aridae/gophermart-diploma/internal/model"
	domainerrors "github.com/aridae/gophermart-diploma/internal/model/domain-errors"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Handle(ctx context.Context) (model.Balance, error) {
	user, authorized := authctx.GetUserFromContext(ctx)
	if !authorized {
		return model.Balance{}, domainerrors.UnauthorizedError()
	}

	_ = user

	// TODO(calculate current balance, return)
	return model.Balance{
		Current:   666.6,
		Withdrawn: 777.7,
	}, nil
}
