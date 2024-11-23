package database

const schemaDDL = `
create table if not exists withdrawals (
    id serial,
    order_number text,
    owner_login text,
    sum_cents bigint,
    created_at time
);

create table if not exists orders (
    id serial,
    order_number text unique,
    order_status text,
    owner_login text,
    accrual_cents bigint default 0,
    created_at time,
    updated_at time,
);

create table if not exists users (
    id serial,
    login text unique,
    pwd_hash bytea,
    created_at time
);
`

const (
	WithdrawalsTable = "withdrawals"
	OrdersTable      = "orders"
	UsersTable       = "users"
)
