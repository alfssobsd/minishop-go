package entities

import uuid "github.com/satori/go.uuid"

type ProductUseCaseEntity struct {
	ProductId         uuid.UUID
	ProductCodeName   string
	ProductTitle      string
	ProductDescrition string
	ProductPrice      float64
}
