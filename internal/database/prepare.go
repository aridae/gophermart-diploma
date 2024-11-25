package database

import (
	"context"
	"fmt"
	"sync"

	"github.com/jackc/pgx/v5/pgconn"
)

var (
	_once sync.Once
)

type sqlQueryable interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (commandTag pgconn.CommandTag, err error)
}

func PrepareSchema(ctx context.Context, db sqlQueryable) error {
	var err error
	_once.Do(func() {
		_, err = db.Exec(ctx, schemaDDL)
	})
	if err != nil {
		return fmt.Errorf("executing schema ddl: %w", err)
	}

	return nil
}
