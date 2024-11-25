package userrepo

import (
	"github.com/Masterminds/squirrel"
	"github.com/aridae/gophermart-diploma/internal/database"
)

const (
	uniqueLoginContraintName = "users_login_key"
)

const (
	loginColumn        = "login"
	passwordHashColumn = "pwd_hash"
	createdAtColumn    = "created_at"
)

var psql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

var baseSelectQuery = psql.Select(
	loginColumn,
	passwordHashColumn,
).From(database.UsersTable)
