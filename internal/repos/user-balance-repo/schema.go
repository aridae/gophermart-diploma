package userbalancerepo

import (
	"github.com/Masterminds/squirrel"
	"github.com/aridae/gophermart-diploma/internal/database"
)

var psql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

var baseSelectQuery = psql.Select(
	"o.owner_login as user_login",
	"sum(o.accrual_cents) as accrual_cents",
	"sum(w.sum_cents) as withdrawn_cents",
).
	From(database.OrdersTable + " o").
	LeftJoin(database.WithdrawalsTable + " w ON w.actor_login = o.owner_login").
	GroupBy("o.owner_login")
