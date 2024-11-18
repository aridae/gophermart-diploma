package userdb

import (
	"github.com/Masterminds/squirrel"
)

const schemaDDL = `create table if not exists users (
    id serial,
    login text unique,
    pwd_hash bytea,
    created_at time
);`

const (
	usersTable               = "users"
	uniqueLoginContraintName = "users_login_key"
)

const (
	idColumn           = "id"
	loginColumn        = "login"
	passwordHashColumn = "pwd_hash"
	createdAtColumn    = "created_at"
)

var psql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

var baseSelectQuery = psql.Select(
	idColumn,
	loginColumn,
	passwordHashColumn,
).From(usersTable)
