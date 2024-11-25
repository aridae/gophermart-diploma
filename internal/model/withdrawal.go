package model

import "time"

type WithdrawalLog struct {
	Sum         Money
	OrderNumber string
	CreatedAt   time.Time
	Actor       User
}
