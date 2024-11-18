package getbalance

import (
	"context"
	"github.com/aridae/gophermart-diploma/internal/model"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Handle(_ context.Context) (model.Balance, error) {
	return model.Balance{
		Current:   666.6,
		Withdrawn: 777.7,
	}, nil
}
