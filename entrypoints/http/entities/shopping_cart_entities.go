package entities

type HttpShoppingCartResponseEntity struct {
	Customer   string                                `json:"customer"`
	TotalGoods int                                   `json:"total_goods"`
	TotalPrice float64                               `json:"total_price"`
	Items      []HttpShoppingCartItemsResponseEntity `json:"items"`
}

type HttpShoppingCartItemsResponseEntity struct {
	Amount int                     `json:"amount"`
	Goods  HttpGoodsResponseEntity `json:"goods"`
}

type HttpShoppingCartAddGoodsRequestEntity struct {
	GoodsId string `json:"goods_id"`
}

type HttpShoppingCartRemoveGoodsRequestEntity struct {
	HttpShoppingCartAddGoodsRequestEntity
}
