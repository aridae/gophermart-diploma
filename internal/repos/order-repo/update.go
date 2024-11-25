package orderrepo

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/aridae/gophermart-diploma/internal/database"
	"github.com/aridae/gophermart-diploma/internal/model"
)

type Setter func(qb squirrel.UpdateBuilder) squirrel.UpdateBuilder

func SetOrderStatus(status model.OrderStatus) Setter {
	return func(qb squirrel.UpdateBuilder) squirrel.UpdateBuilder {
		return qb.Set(orderStatusColumn, status.String())
	}
}

func SetOrderAccrual(accrual model.Money) Setter {
	return func(qb squirrel.UpdateBuilder) squirrel.UpdateBuilder {
		return qb.Set(accrualCentsColumn, accrual.Cents())
	}
}

func (r *Repository) UpdateOrder(ctx context.Context, orderNumber string, setters ...Setter) error {
	queryable := r.txGetter.DefaultTrOrDB(ctx, r.db)

	qb := psql.Update(database.OrdersTable).Where(squirrel.Eq{orderNumberColumn: orderNumber})
	qb = applySetters(qb, setters)

	sql, args, err := qb.ToSql()
	if err != nil {
		return fmt.Errorf("failed to build query: %w", err)
	}

	if _, err = queryable.Exec(ctx, sql, args...); err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}

	return nil
}

func applySetters(qb squirrel.UpdateBuilder, setters []Setter) squirrel.UpdateBuilder {
	for _, setter := range setters {
		qb = setter(qb)
	}

	return qb
}
