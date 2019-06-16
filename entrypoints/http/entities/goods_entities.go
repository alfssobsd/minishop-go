package entities

import uuid "github.com/satori/go.uuid"

type HttpGoodsListResponseEntity struct {
	Total  int                       `json:"total"`
	Offset int                       `json:"offset"`
	Items  []HttpGoodsResponseEntity `json:"items"`
}

type HttpGoodsResponseEntity struct {
	GoodsId          uuid.UUID `json:"goods_id"`
	GoodsCodeName    string    `json:"goods_code_name"`
	GoodsTitle       string    `json:"goods_title"`
	GoodsDescription string    `json:"goods_description"`
	GoodsPrice       float64   `json:"goods_price"`
}

type HttpGoodsRequestEntity struct {
	GoodsCodeName    string  `json:"goods_code_name"`
	GoodsTitle       string  `json:"goods_title"`
	GoodsDescription string  `json:"goods_description"`
	GoodsPrice       float64 `json:"goods_price"`
}
