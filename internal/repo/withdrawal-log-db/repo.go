package withdrawallogdb

import (
	"context"
	trmpgx "github.com/avito-tech/go-transaction-manager/drivers/pgxv5/v2"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type sqlQueryable interface {
	Begin(ctx context.Context) (pgx.Tx, error)
	CopyFrom(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error)
	SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults
	Exec(ctx context.Context, sql string, arguments ...interface{}) (commandTag pgconn.CommandTag, err error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
}

type sqlTransaxtionManager interface {
	DefaultTrOrDB(ctx context.Context, db trmpgx.Tr) trmpgx.Tr
}

type Repository struct {
	db       sqlQueryable
	txGetter sqlTransaxtionManager
}

func NewRepository(
	db sqlQueryable,
	txGetter sqlTransaxtionManager,
) *Repository {
	imp := &Repository{
		db:       db,
		txGetter: txGetter,
	}

	return imp
}
