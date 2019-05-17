package entities

type GoodsEntity struct {
	GoodsID         string  `bson:"_id,omitempty"`
	GoodsCodeName   string  `bson:"goods_code_name"`
	GoodsTitle      string  `bson:"goods_title"`
	GoodsDescrition string  `bson:"goods_description"`
	GoodsPrice      float64 `bson:"goods_price"`
}
