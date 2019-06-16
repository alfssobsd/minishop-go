package entities

import uuid "github.com/satori/go.uuid"

type CustomerEntity struct {
	CustomerId       uuid.UUID `db:"uuid"`
	CustomerUsername string    `db:"username"`
	CustomerFullName string    `db:"full_name"`
}
