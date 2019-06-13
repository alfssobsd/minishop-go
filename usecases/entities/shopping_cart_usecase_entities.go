package entities

type ShoppingCartUseCaseEntity struct {
	Customer   string
	TotalPrice float64
	GoodsItems []GoodsUseCaseEntity
}
