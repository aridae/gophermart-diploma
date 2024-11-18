package requestwithdrawal

import (
	"context"
)

type Request struct {
	OrderNumber string
	Sum         float32
}

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Handle(_ context.Context, _ Request) error {
	return nil
}
