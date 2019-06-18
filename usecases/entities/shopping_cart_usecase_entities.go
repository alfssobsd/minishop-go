package entities

import uuid "github.com/satori/go.uuid"

type ShoppingCartUseCaseEntity struct {
	CustomerId   uuid.UUID
	TotalPrice   float64
	ProductItems []ShoppingCartProductItemUseCaseEntity
}

type ShoppingCartProductItemUseCaseEntity struct {
	Amount  int
	Product ProductUseCaseEntity
}
