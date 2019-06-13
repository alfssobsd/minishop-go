package entities

type OrderStatus int

const (
	OrderStatusActive OrderStatus = 0
	OrderStatusDone   OrderStatus = 1
)

type OrderEntity struct {
	OrderID       string      `db:"uuid"`
	OrderCustomer string      `db:"customer"`
	OrderStatus   OrderStatus `db:"status"`
	OrderItems    []OrderItemEntity
}

type OrderItemEntity struct {
	GoodsItem   GoodsEntity
	GoodsAmount int
}

type OrderGoodsEntity struct {
	OrderID     string `db:"order_uuid"`
	GoodsId     string `db:"goods_uuid"`
	GoodsAmount int    `db:"goods_amount"`
}
