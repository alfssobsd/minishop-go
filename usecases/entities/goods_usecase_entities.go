package entities

import uuid "github.com/satori/go.uuid"

type GoodsUseCaseEntity struct {
	GoodsId         uuid.UUID
	GoodsCodeName   string
	GoodsTitle      string
	GoodsDescrition string
	GoodsPrice      float64
}
