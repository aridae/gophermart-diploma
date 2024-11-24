package userbalancerepo

import "github.com/aridae/gophermart-diploma/internal/model"

type balanceDTO struct {
	UserLogin           string `db:"user_login"`
	CurrentBalanceCents int64  `db:"current_balance_cents"`
	WithdrawnCents      int64  `db:"withdrawn_cents"`
}

func mapDTOToDomainUserBalance(dto balanceDTO) model.Balance {
	return model.Balance{
		Owner: model.User{
			Login: dto.UserLogin,
		},
		Current:   model.NewMoneyFromCents(dto.CurrentBalanceCents),
		Withdrawn: model.NewMoneyFromCents(dto.WithdrawnCents),
	}
}
