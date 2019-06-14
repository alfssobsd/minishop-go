package usecases

import (
	_repo "github.com/alfssobsd/minishop/dataproviders/postgres"
	"github.com/alfssobsd/minishop/usecases/entities"
	uuid "github.com/satori/go.uuid"
)

type ShoppingCartUseCase interface {
	AddGoodsToCartUseCase(customer string, goodsId string) *entities.ShoppingCartUseCaseEntity
	RemoveGoodsFormCartUseCase(customer string) *entities.ShoppingCartUseCaseEntity
	ShowCartUseCase(customer string) *entities.ShoppingCartUseCaseEntity
}

type shoppingCartUseCase struct {
	goodsRepo _repo.GoodsRepository
	orderRepo _repo.OrderRepository
}

func NewShoppingCartUseCase(goodsRepository _repo.GoodsRepository, orderRepository _repo.OrderRepository) *shoppingCartUseCase {
	return &shoppingCartUseCase{goodsRepository, orderRepository}
}

func (u *shoppingCartUseCase) AddGoodsToCartUseCase(customer string, goodsId string) *entities.ShoppingCartUseCaseEntity {
	if u.goodsRepo.FindById(goodsId) == nil {
		return nil
	}

	order := u.orderRepo.GetFirstActiveOrder(customer)
	if order == nil {
		orderId := uuid.NewV4().String()
		u.orderRepo.CreateOrder(customer, orderId)
		order = u.orderRepo.GetFirstActiveOrder(customer)
	}

	u.orderRepo.AddGoods(order.OrderID, goodsId)
	order = u.orderRepo.GetFirstActiveOrder(customer)

	totalPrice := float64(0)
	goodsItems := []entities.ShoppingCartGoodsItemUseCaseEntity{}
	for _, element := range order.OrderItems {
		goodsItems = append(goodsItems, entities.ShoppingCartGoodsItemUseCaseEntity{
			Goods: entities.GoodsUseCaseEntity{
				GoodsId:         element.GoodsItem.GoodsID,
				GoodsCodeName:   element.GoodsItem.GoodsCodeName,
				GoodsDescrition: element.GoodsItem.GoodsDescrition,
				GoodsTitle:      element.GoodsItem.GoodsTitle,
				GoodsPrice:      element.GoodsItem.GoodsPrice,
			},
			Amount: element.GoodsAmount,
		})
		for i := 0; i < element.GoodsAmount; i++ {
			totalPrice += element.GoodsItem.GoodsPrice
		}
	}

	return &entities.ShoppingCartUseCaseEntity{Customer: customer, TotalPrice: totalPrice, GoodsItems: goodsItems}
}

func (u *shoppingCartUseCase) RemoveGoodsFormCartUseCase(customer string) *entities.ShoppingCartUseCaseEntity {
	//TODO: need implement
	return nil
}

func (u *shoppingCartUseCase) ShowCartUseCase(customer string) *entities.ShoppingCartUseCaseEntity {
	order := u.orderRepo.GetFirstActiveOrder(customer)
	if order == nil {
		return &entities.ShoppingCartUseCaseEntity{Customer: customer, TotalPrice: float64(0), GoodsItems: []entities.ShoppingCartGoodsItemUseCaseEntity{}}
	}

	totalPrice := float64(0)
	goodsItems := []entities.ShoppingCartGoodsItemUseCaseEntity{}
	for _, element := range order.OrderItems {
		goodsItems = append(goodsItems, entities.ShoppingCartGoodsItemUseCaseEntity{
			Goods: entities.GoodsUseCaseEntity{
				GoodsId:         element.GoodsItem.GoodsID,
				GoodsCodeName:   element.GoodsItem.GoodsCodeName,
				GoodsDescrition: element.GoodsItem.GoodsDescrition,
				GoodsTitle:      element.GoodsItem.GoodsTitle,
				GoodsPrice:      element.GoodsItem.GoodsPrice,
			},
			Amount: element.GoodsAmount,
		})
		for i := 0; i < element.GoodsAmount; i++ {
			totalPrice += element.GoodsItem.GoodsPrice
		}
	}

	return &entities.ShoppingCartUseCaseEntity{Customer: customer, TotalPrice: totalPrice, GoodsItems: goodsItems}
}
