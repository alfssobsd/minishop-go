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
create table IF NOT EXISTS goods
(
    uuid uuid not null,
    title varchar not null,
    code_name varchar not null,
    description varchar,
    price decimal(10,2) not null,
	PRIMARY KEY (uuid)
);

create unique index IF NOT EXISTS goods_uuid_uindex
    on goods (uuid);`

	db.MustExec(schema)
}
