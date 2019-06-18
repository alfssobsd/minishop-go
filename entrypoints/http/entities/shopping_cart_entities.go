package entities

import uuid "github.com/satori/go.uuid"

type HttpShoppingCartResponseEntity struct {
	CustomerId    uuid.UUID                             `json:"customer_id"`
	TotalProducts int                                   `json:"total_products"`
	TotalPrice    float64                               `json:"total_price"`
	Items         []HttpShoppingCartItemsResponseEntity `json:"items"`
}

type HttpShoppingCartItemsResponseEntity struct {
	Amount  int                       `json:"amount"`
	Product HttpProductResponseEntity `json:"product"`
}

type HttpShoppingCartAddProductRequestEntity struct {
	ProductId uuid.UUID `json:"product_id"`
}

type HttpShoppingCartRemoveProductRequestEntity struct {
	HttpShoppingCartAddProductRequestEntity
}
