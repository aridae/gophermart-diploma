package model

import "time"

type Withdrawal struct {
	OrderNumber string
	ProcessedAt time.Time
	Sum         float32
}
