package submitorder

import (
	"context"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Handle(ctx context.Context) error {
	return nil
}
