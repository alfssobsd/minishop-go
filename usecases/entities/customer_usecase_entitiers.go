package entities

import uuid "github.com/satori/go.uuid"

type CustomerUseCaseEntity struct {
	CustomerId       uuid.UUID
	CustomerUsername string
	CustomerFullName string
}
