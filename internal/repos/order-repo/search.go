package orderrepo

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/aridae/gophermart-diploma/internal/model"
	"github.com/aridae/gophermart-diploma/pkg/slice"
	"github.com/georgysavva/scany/v2/pgxscan"
)

type Pagination struct {
	Page  int
	Limit int
}

type Filter struct {
	Statuses []model.OrderStatus
}

func (r *Repository) Search(ctx context.Context, filter Filter, pagination Pagination) ([]model.Order, error) {
	queryable := r.txGetter.DefaultTrOrDB(ctx, r.db)

	qb := baseSelectQuery
	qb = applyFilter(qb, filter)
	qb = applyPagination(qb, pagination)

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

func applyFilter(qb squirrel.SelectBuilder, filter Filter) squirrel.SelectBuilder {
	if len(filter.Statuses) > 0 {
		qb = qb.Where(squirrel.Eq{orderStatusColumn: slice.MapBatchNoErr(filter.Statuses, func(in model.OrderStatus) string {
			return in.String()
		})})
	}

	return qb
}

func applyPagination(qb squirrel.SelectBuilder, pagination Pagination) squirrel.SelectBuilder {
	return qb.Suffix(fmt.Sprintf("LIMIT %d OFFSET %d", pagination.Limit, (pagination.Page-1)*pagination.Limit))
}
