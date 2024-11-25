package orderrepo

import (
	"github.com/Masterminds/squirrel"
	"github.com/aridae/gophermart-diploma/internal/database"
)

const (
	uniqueOrderNumberContraintName = "orders_order_number_key"
)

const (
	orderNumberColumn  = "order_number"
	orderStatusColumn  = "order_status"
	ownerLoginColumn   = "owner_login"
	accrualCentsColumn = "accrual_cents"
	createdAtColumn    = "created_at"
)

var psql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

var baseSelectQuery = psql.Select(
	orderNumberColumn,
	orderStatusColumn,
	ownerLoginColumn,
	accrualCentsColumn,
	createdAtColumn,
).From(database.OrdersTable)
