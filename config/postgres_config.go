package config

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
)

func MakePostgresConnection() *sqlx.DB {
	connSettings := "postgres://postgres:password@localhost/minishop?sslmode=disable"
	db, err := sqlx.Open("postgres", connSettings)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func RunMigration(db *sqlx.DB) {
	log.Info("Run DB Migration")
	schema := `
create table IF NOT EXISTS products
(
    uuid uuid not null,
    title varchar not null,
    code_name varchar not null,
    description varchar,
    price decimal(10,2) not null,
	PRIMARY KEY (uuid)
);

create unique index IF NOT EXISTS products_uuid_uindex
    on products (uuid);

create table IF NOT EXISTS order_products
(
	product_uuid uuid not null,
	product_amount int default 1 not null,
	order_uuid uuid not null
);
create unique index IF NOT EXISTS order_products_product_uuid_order_uuid_uindex
	on order_products (product_uuid, order_uuid);

create table IF NOT EXISTS orders
(
    uuid uuid not null UNIQUE,
    customer_uuid varchar not null,
    status integer not null,
	PRIMARY KEY (uuid)
);
create table IF NOT EXISTS customers
(
	uuid uuid not null,
	username varchar not null,
	full_name varchar not null,
	PRIMARY KEY (uuid)
);

create unique index IF NOT EXISTS customers_username_uindex
	on customers (username);

create unique index IF NOT EXISTS customers_uuid_uindex
	on customers (uuid);
`

	db.MustExec(schema)
}
