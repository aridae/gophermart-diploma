package withdrawallogdb

import (
	"github.com/Masterminds/squirrel"
	"github.com/aridae/gophermart-diploma/internal/database"
)

const (
	sumCentsColumn    = "sum_cents"
	orderNumberColumn = "order_number"
	actorLoginColumn  = "actor_login"
	requestedAtColumn = "requested_at"
)

var psql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

var baseSelectQuery = psql.Select(
	orderNumberColumn,
	actorLoginColumn,
	sumCentsColumn,
	requestedAtColumn,
).From(database.WithdrawalsTable)
