package getwithdrawalshistory

import (
	"context"
	"github.com/aridae/gophermart-diploma/internal/model"
	"time"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Handle(_ context.Context) ([]model.Withdrawal, error) {
	return []model.Withdrawal{
		{
			OrderNumber: "111333222666",
			ProcessedAt: time.Now().UTC(),
			Sum:         12341,
		},
	}, nil
}
