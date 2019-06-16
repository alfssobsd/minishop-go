package entities

import uuid "github.com/satori/go.uuid"

type ShoppingCartUseCaseEntity struct {
	CustomerId uuid.UUID
	TotalPrice float64
	GoodsItems []ShoppingCartGoodsItemUseCaseEntity
}

type ShoppingCartGoodsItemUseCaseEntity struct {
	Amount int
	Goods  GoodsUseCaseEntity
}
