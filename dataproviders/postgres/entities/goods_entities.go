package entities

type GoodsEntity struct {
	GoodsID         string  `db:"uuid"`
	GoodsCodeName   string  `db:"code_name"`
	GoodsTitle      string  `db:"title"`
	GoodsDescrition string  `db:"description"`
	GoodsPrice      float64 `db:"price"`
}
