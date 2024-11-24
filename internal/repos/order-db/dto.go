package orderdb

import (
	"time"

	"github.com/aridae/gophermart-diploma/internal/model"
	"github.com/aridae/gophermart-diploma/pkg/pointer"
)

type orderDTO struct {
	OrderNumber  string    `db:"order_number"`
	OrderStatus  string    `db:"order_status"`
	OwnerLogin   string    `db:"owner_login"`
	AccrualCents int64     `db:"accrual_cents"`
	CreatedAt    time.Time `db:"created_at"`
}

func mapDTOToDomainOrder(dto orderDTO) model.Order {
	var accrual *model.Money
	if dto.AccrualCents != 0 {
		accrual = pointer.To(model.NewMoneyFromCents(dto.AccrualCents))
	}

	return model.Order{
		OrderSubmit: model.OrderSubmit{
			Number: dto.OrderNumber,
			Owner:  model.User{Login: dto.OwnerLogin},
			Status: model.OrderStatus(dto.OrderStatus),
		},
		Accrual:    accrual,
		UploadedAt: dto.CreatedAt,
	}
}
