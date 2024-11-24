package userbalancerepo

import (
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/aridae/gophermart-diploma/internal/database"
)

var psql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

var baseSelectQuery = psql.Select(
	fmt.Sprintf("%s.login as %s", database.UsersTable, "user_login"),
	fmt.Sprintf("sum(%s.accrual_cents) as %s", database.OrdersTable, "current_balance_cents"),
	fmt.Sprintf("sum(%s.sum_cents) as %s", database.WithdrawalsTable, "withdrawn_cents"),
).From(database.UsersTable).
	Join(database.OrdersTable + " ON " + database.UsersTable + ".login = " + database.OrdersTable + ".owner_login").
	LeftJoin(database.WithdrawalsTable + " ON " + database.OrdersTable + ".order_number = " + database.WithdrawalsTable + ".order_number").
	GroupBy(fmt.Sprintf("%s.login", database.UsersTable))
