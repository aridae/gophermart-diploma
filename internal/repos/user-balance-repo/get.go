package userbalancerepo

import (
	"context"
	"errors"
	"fmt"
	"github.com/aridae/gophermart-diploma/internal/database"

	"github.com/Masterminds/squirrel"
	"github.com/aridae/gophermart-diploma/internal/model"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
)

func (r *Repository) GetUserBalance(ctx context.Context, user model.User) (model.Balance, error) {
	queryable := r.txGetter.DefaultTrOrDB(ctx, r.db)

	qb := baseSelectQuery.Where(squirrel.Eq{fmt.Sprintf("%s.login", database.UsersTable): user.Login})

	sql, args, err := qb.ToSql()
	if err != nil {
		return model.Balance{}, fmt.Errorf("failed to build query: %w", err)
	}

	var dto balanceDTO
	if err = pgxscan.Get(ctx, queryable, &dto, sql, args...); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.Balance{}, nil
		}
		return model.Balance{}, fmt.Errorf("failed to execute query: %w", err)
	}

	balance := mapDTOToDomainUserBalance(dto)

	return balance, nil
}
