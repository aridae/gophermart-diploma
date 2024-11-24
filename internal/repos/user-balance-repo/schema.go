package userbalancerepo

import (
	"github.com/Masterminds/squirrel"
	"github.com/aridae/gophermart-diploma/internal/database"
)

var psql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

var currentBalanceSubquery, _, _ = psql.Select(
	"owner_login",
	"sum(accrual_cents) as current_balance_cents",
).From(database.OrdersTable).GroupBy("owner_login").ToSql()

var withdrawnSubquery, _, _ = psql.Select(
	"actor_login",
	"sum(sum_cents) as withdrawn_cents",
).From(database.WithdrawalsTable).GroupBy("actor_login").ToSql()

var baseSelectQuery = psql.Select(
	"users.login as user_login",
	"ws.withdrawn_cents as withdrawn_cents",
	"cbs.current_balance_cents as current_balance_cents",
).From(database.UsersTable).
	Join("(" + currentBalanceSubquery + ") cbs ON " + database.UsersTable + ".login = cbs.owner_login").
	LeftJoin("(" + withdrawnSubquery + ") ws ON " + database.UsersTable + ".login = ws.actor_login")
