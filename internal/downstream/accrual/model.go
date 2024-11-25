package accrual

import "github.com/aridae/gophermart-diploma/internal/model"

type orderDTO struct {
	Number  string  `json:"order"`
	Status  string  `json:"status"`
	Accrual float32 `json:"accrual"`
}

type Order struct {
	Number  string
	Status  model.OrderStatus
	Accrual model.Money
}

func parseOrderDTO(dto orderDTO) Order {
	return Order{
		Number:  dto.Number,
		Status:  model.OrderStatus(dto.Status),
		Accrual: model.NewMoney(dto.Accrual),
	}
}
