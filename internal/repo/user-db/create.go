package userdb

import (
	"context"
	"errors"
	"fmt"
	"github.com/aridae/gophermart-diploma/internal/model"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"time"
)

func (r *Repo) CreateUser(ctx context.Context, user model.UserCredentials, now time.Time) (int64, error) {
	queryable := r.txGetter.DefaultTrOrDB(ctx, r.db)

	qb := psql.Insert(usersTable).
		Columns(
			loginColumn,
			passwordHashColumn,
			createdAtColumn,
		).
		Values(
			user.Login,
			user.PasswordHash,
			now,
		).
		Suffix("RETURNING id")

	sql, args, err := qb.ToSql()
	if err != nil {
		return 0, fmt.Errorf("failed to build query: %w", err)
	}

	var id int64

	if err = pgxscan.Get(ctx, queryable, &id, sql, args...); err != nil {
		if pgerr := new(pgconn.PgError); errors.As(err, &pgerr) && isLoginUniqueConstraintViolated(pgerr) {
			return 0, LoginUniqueConstraintViolatedError
		}

		return 0, fmt.Errorf("failed to execute query: %w", err)
	}

	return id, nil
}

func isLoginUniqueConstraintViolated(pgerr *pgconn.PgError) bool {
	return pgerr.Code == pgerrcode.UniqueViolation && pgerr.ConstraintName == uniqueLoginContraintName
}
