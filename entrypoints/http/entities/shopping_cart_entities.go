package entities

type HttpShoppingCartEntity struct {
	Customer   string                    `json:"customer"`
	TotalGoods int                       `json:"total_goods"`
	TotalPrice float64                   `json:"total_price"`
	GoodsItems []HttpGoodsResponseEntity `json:"goods_items"`
}
