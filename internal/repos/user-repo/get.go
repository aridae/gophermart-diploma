package userrepo

import (
	"context"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/aridae/gophermart-diploma/internal/model"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
)

func (r *Repository) GetUserCredentials(ctx context.Context, login string) (*model.UserCredentials, error) {
	queryable := r.txGetter.DefaultTrOrDB(ctx, r.db)

	qb := baseSelectQuery.Where(squirrel.Eq{loginColumn: login})

	sql, args, err := qb.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %w", err)
	}

	var dto userDTO
	if err = pgxscan.Get(ctx, queryable, &dto, sql, args...); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	user := mapDTOToDomainUserCredentials(dto)

	return &user, nil
}
