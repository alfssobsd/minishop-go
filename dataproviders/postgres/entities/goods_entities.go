package entities

import uuid "github.com/satori/go.uuid"

type GoodsEntity struct {
	GoodsID         uuid.UUID `db:"uuid"`
	GoodsCodeName   string    `db:"code_name"`
	GoodsTitle      string    `db:"title"`
	GoodsDescrition string    `db:"description"`
	GoodsPrice      float64   `db:"price"`
}
