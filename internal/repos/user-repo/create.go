package userrepo

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

func (r *Repository) CreateUser(ctx context.Context, user model.UserCredentials, now time.Time) error {
	queryable := r.txGetter.DefaultTrOrDB(ctx, r.db)

	qb := psql.Insert(database.UsersTable).
		Columns(
			loginColumn,
			passwordHashColumn,
			createdAtColumn,
		).
		Values(
			user.Login,
			user.PasswordHash,
			now,
		)

	sql, args, err := qb.ToSql()
	if err != nil {
		return fmt.Errorf("failed to build query: %w", err)
	}

	if _, err = queryable.Exec(ctx, sql, args...); err != nil {
		if pgerr := new(pgconn.PgError); errors.As(err, &pgerr) && isLoginUniqueConstraintViolated(pgerr) {
			return LoginUniqueConstraintViolatedError
		}

		return fmt.Errorf("failed to execute query: %w", err)
	}

	return nil
}

func isLoginUniqueConstraintViolated(pgerr *pgconn.PgError) bool {
	return pgerr.Code == pgerrcode.UniqueViolation && pgerr.ConstraintName == uniqueLoginContraintName
}
