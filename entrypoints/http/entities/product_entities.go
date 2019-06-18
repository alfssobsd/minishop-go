package entities

import uuid "github.com/satori/go.uuid"

type HttpProductListResponseEntity struct {
	Total  int                         `json:"total"`
	Offset int                         `json:"offset"`
	Items  []HttpProductResponseEntity `json:"items"`
}

type HttpProductResponseEntity struct {
	ProductId          uuid.UUID `json:"product_id"`
	ProductCodeName    string    `json:"product_code_name"`
	ProductTitle       string    `json:"product_title"`
	ProductDescription string    `json:"product_description"`
	ProductPrice       float64   `json:"product_price"`
}

type HttpProductRequestEntity struct {
	ProductCodeName    string  `json:"product_code_name"`
	ProductTitle       string  `json:"product_title"`
	ProductDescription string  `json:"product_description"`
	ProductPrice       float64 `json:"product_price"`
}
