package entities

import uuid "github.com/satori/go.uuid"

type HttpShoppingCartResponseEntity struct {
	CustomerId uuid.UUID                             `json:"customer_id"`
	TotalGoods int                                   `json:"total_goods"`
	TotalPrice float64                               `json:"total_price"`
	Items      []HttpShoppingCartItemsResponseEntity `json:"items"`
}

type HttpShoppingCartItemsResponseEntity struct {
	Amount int                     `json:"amount"`
	Goods  HttpGoodsResponseEntity `json:"goods"`
}

type HttpShoppingCartAddGoodsRequestEntity struct {
	GoodsId uuid.UUID `json:"goods_id"`
}

type HttpShoppingCartRemoveGoodsRequestEntity struct {
	HttpShoppingCartAddGoodsRequestEntity
}
