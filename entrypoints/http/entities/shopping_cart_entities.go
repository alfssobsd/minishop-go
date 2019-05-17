package entities

type HttpShoppingCartEntity struct {
	ShoppingCartId string   `json:"shopping_cart_id"`
	TotalGoods     int      `json:"total_goods"`
	GoodsIds       []string `json:"goods_ids"`
}
