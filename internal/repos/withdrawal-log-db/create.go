package withdrawallogdb

import (
	"context"
	"fmt"
	"time"

	"github.com/aridae/gophermart-diploma/internal/database"
	"github.com/aridae/gophermart-diploma/internal/model"
)

func (r *Repository) CreateWithdrawalLog(ctx context.Context, withdrawal model.WithdrawalLog, now time.Time) error {
	queryable := r.txGetter.DefaultTrOrDB(ctx, r.db)

	qb := psql.Insert(database.WithdrawalsTable).
		Columns(
			orderNumberColumn,
			actorLoginColumn,
			sumCentsColumn,
			requestedAtColumn,
		).
		Values(
			withdrawal.OrderNumber,
			withdrawal.Actor.Login,
			withdrawal.Sum.Cents(),
			now,
		)

	sql, args, err := qb.ToSql()
	if err != nil {
		return fmt.Errorf("failed to build query: %w", err)
	}

	if _, err = queryable.Exec(ctx, sql, args...); err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}

	return nil
}
