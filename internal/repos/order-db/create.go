package orderdb

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/aridae/gophermart-diploma/internal/database"
	"github.com/aridae/gophermart-diploma/internal/model"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

func (r *Repository) CreateOrder(ctx context.Context, orderSubmit model.OrderSubmit, now time.Time) error {
	queryable := r.txGetter.DefaultTrOrDB(ctx, r.db)

	qb := psql.Insert(database.OrdersTable).
		Columns(
			orderNumberColumn,
			orderStatusColumn,
			ownerLoginColumn,
			createdAtColumn,
		).
		Values(
			orderSubmit.Number,
			orderSubmit.Status.String(),
			orderSubmit.Owner.Login,
			now,
		)

	sql, args, err := qb.ToSql()
	if err != nil {
		return fmt.Errorf("failed to build query: %w", err)
	}

	if _, err = queryable.Exec(ctx, sql, args...); err != nil {
		if pgerr := new(pgconn.PgError); errors.As(err, &pgerr) && isOrderNumberUniqueConstraintViolated(pgerr) {
			return OrderNumberUniqueConstraintViolatedError
		}

		return fmt.Errorf("failed to execute query: %w", err)
	}

	return nil
}

func isOrderNumberUniqueConstraintViolated(pgerr *pgconn.PgError) bool {
	return pgerr.Code == pgerrcode.UniqueViolation && pgerr.ConstraintName == uniqueOrderNumberContraintName
}
