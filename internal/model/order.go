package model

import "time"

type OrderStatus string

const (
	OrderStatusInvalid    OrderStatus = "INVALID"
	OrderStatusNew        OrderStatus = "NEW"
	OrderStatusProcessed  OrderStatus = "PROCESSED"
	OrderStatusProcessing OrderStatus = "PROCESSING"
)

type OrderSubmit struct {
	Number string
}

type Order struct {
	OrderSubmit
	Accrual    *int
	Status     OrderStatus
	UploadedAt time.Time
}
