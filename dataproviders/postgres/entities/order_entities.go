package entities

import uuid "github.com/satori/go.uuid"

type OrderStatus int

const (
	OrderStatusActive OrderStatus = 0
	OrderStatusDone   OrderStatus = 1
)

type OrderEntity struct {
	OrderID         uuid.UUID   `db:"uuid"`
	OrderCustomerId uuid.UUID   `db:"customer_uuid"`
	OrderStatus     OrderStatus `db:"status"`
	OrderItems      []OrderItemEntity
}

type OrderItemEntity struct {
	ProductItem   ProductEntity
	ProductAmount int
}

type OrderProductEntity struct {
	OrderID       uuid.UUID `db:"order_uuid"`
	ProductId     uuid.UUID `db:"product_uuid"`
	ProductAmount int       `db:"product_amount"`
}
