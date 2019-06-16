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
	GoodsItem   GoodsEntity
	GoodsAmount int
}

type OrderGoodsEntity struct {
	OrderID     uuid.UUID `db:"order_uuid"`
	GoodsId     uuid.UUID `db:"goods_uuid"`
	GoodsAmount int       `db:"goods_amount"`
}
