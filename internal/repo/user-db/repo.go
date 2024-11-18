package userdb

import (
	"context"
	"fmt"
	trmpgx "github.com/avito-tech/go-transaction-manager/drivers/pgxv5/v2"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"sync"
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

type Repo struct {
	db       sqlQueryable
	txGetter sqlTransaxtionManager
}

func NewRepo(
	ctx context.Context,
	db sqlQueryable,
	txGetter sqlTransaxtionManager,
) (*Repo, error) {
	imp := &Repo{
		db:       db,
		txGetter: txGetter,
	}

	err := imp.prepareSchema(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare schema: %w", err)
	}

	return imp, nil
}

var (
	_once sync.Once
)

func (r *Repo) prepareSchema(ctx context.Context) error {
	queryable := r.txGetter.DefaultTrOrDB(ctx, r.db)

	var err error
	_once.Do(func() {
		_, err = queryable.Exec(ctx, schemaDDL)
	})
	if err != nil {
		return fmt.Errorf("executing schema ddl: %w", err)
	}

	return nil
}
