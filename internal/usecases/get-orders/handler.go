package getorders

import (
	"context"
	"github.com/aridae/gophermart-diploma/internal/model"
	"time"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Handle(_ context.Context) ([]model.Order, error) {
	accural := 777
	return []model.Order{
		{
			OrderSubmit: model.OrderSubmit{
				Number: "111333222666",
			},
			Accrual:    &accural,
			Status:     model.OrderStatusNew,
			UploadedAt: time.Now().UTC(),
		},
	}, nil
}
