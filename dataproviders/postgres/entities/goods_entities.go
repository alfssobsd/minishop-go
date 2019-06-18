package entities

import uuid "github.com/satori/go.uuid"

type ProductEntity struct {
	ProductID         uuid.UUID `db:"uuid"`
	ProductCodeName   string    `db:"code_name"`
	ProductTitle      string    `db:"title"`
	ProductDescrition string    `db:"description"`
	ProductPrice      float64   `db:"price"`
}
