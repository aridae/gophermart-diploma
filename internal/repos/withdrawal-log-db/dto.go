package withdrawallogdb

import (
	"time"

	"github.com/aridae/gophermart-diploma/internal/model"
)

type withdrawalLogDTO struct {
	SumCents    int64     `db:"sum_cents"`
	OrderNumber string    `db:"order_number"`
	ActorLogin  string    `db:"actor_login"`
	CreatedAt   time.Time `db:"created_at"`
}

func mapDTOToDomainWithdrawalLog(dto withdrawalLogDTO) model.WithdrawalLog {
	return model.WithdrawalLog{
		Sum:         model.NewMoneyFromCents(dto.SumCents),
		OrderNumber: dto.OrderNumber,
		CreatedAt:   dto.CreatedAt,
		Actor:       model.User{Login: dto.ActorLogin},
	}
}
