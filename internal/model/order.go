package model

import "time"

type OrderStatus string

func (s OrderStatus) String() string {
	return string(s)
}

const (
	OrderStatusInvalid    OrderStatus = "INVALID"
	OrderStatusNew        OrderStatus = "NEW"
	OrderStatusProcessed  OrderStatus = "PROCESSED"
	OrderStatusProcessing OrderStatus = "PROCESSING"
)

type OrderSubmit struct {
	Number string
	Owner  User
	Status OrderStatus
}

type Order struct {
	OrderSubmit
	Accrual    *Money
	UploadedAt time.Time
}
