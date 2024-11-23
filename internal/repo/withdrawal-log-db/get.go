package withdrawallogdb

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/aridae/gophermart-diploma/internal/model"
	"github.com/aridae/gophermart-diploma/pkg/slice"
	"github.com/georgysavva/scany/v2/pgxscan"
)

func (r *Repository) GetByActor(ctx context.Context, actor model.User) ([]model.WithdrawalLog, error) {
	queryable := r.txGetter.DefaultTrOrDB(ctx, r.db)

	qb := baseSelectQuery.Where(squirrel.Eq{actorLoginColumn: actor.Login})

	sql, args, err := qb.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %w", err)
	}

	var dtos []withdrawalLogDTO
	if err = pgxscan.Select(ctx, queryable, &dtos, sql, args...); err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}

	withdrawalsLogs := slice.MapBatchNoErr(dtos, mapDTOToDomainWithdrawalLog)

	return withdrawalsLogs, nil
}
