package entities

type ShoppingCartUseCaseEntity struct {
	Customer   string
	TotalPrice float64
	GoodsItems []ShoppingCartGoodsItemUseCaseEntity
}

type ShoppingCartGoodsItemUseCaseEntity struct {
	Amount int
	Goods  GoodsUseCaseEntity
}
