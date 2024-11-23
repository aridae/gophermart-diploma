package orderdb

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/aridae/gophermart-diploma/internal/model"
	"github.com/aridae/gophermart-diploma/pkg/slice"
	"github.com/georgysavva/scany/v2/pgxscan"
)

func (r *Repository) GetByOwner(ctx context.Context, owner model.User) ([]model.Order, error) {
	queryable := r.txGetter.DefaultTrOrDB(ctx, r.db)

	qb := baseSelectQuery.Where(squirrel.Eq{ownerLoginColumn: owner.Login})

	sql, args, err := qb.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %w", err)
	}

	var dtos []orderDTO
	if err = pgxscan.Select(ctx, queryable, &dtos, sql, args...); err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	orders := slice.MapBatchNoErr(dtos, mapDTOToDomainOrder)

	return orders, nil
}

func (r *Repository) GetByNumbers(ctx context.Context, orderNumbers []string) ([]model.Order, error) {
	queryable := r.txGetter.DefaultTrOrDB(ctx, r.db)

	qb := baseSelectQuery.Where(squirrel.Eq{orderNumberColumn: orderNumbers})

	sql, args, err := qb.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %w", err)
	}

	var dtos []orderDTO
	if err = pgxscan.Select(ctx, queryable, &dtos, sql, args...); err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	orders := slice.MapBatchNoErr(dtos, mapDTOToDomainOrder)

	return orders, nil
}
