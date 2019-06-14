package usecases

import (
	"errors"
	_repo "github.com/alfssobsd/minishop/dataproviders/postgres"
	"github.com/alfssobsd/minishop/usecases/entities"
	uuid "github.com/satori/go.uuid"
)

type ShoppingCartUseCase interface {
	AddGoodsToCartUseCase(customer string, goodsId string) error
	RemoveGoodsFormCartUseCase(customer string, goodsId string) error
	ShowCartUseCase(customer string) (*entities.ShoppingCartUseCaseEntity, error)
}

type shoppingCartUseCase struct {
	goodsRepo _repo.GoodsRepository
	orderRepo _repo.OrderRepository
}

func NewShoppingCartUseCase(goodsRepository _repo.GoodsRepository, orderRepository _repo.OrderRepository) *shoppingCartUseCase {
	return &shoppingCartUseCase{goodsRepository, orderRepository}
}

func (u *shoppingCartUseCase) AddGoodsToCartUseCase(customer string, goodsId string) error {
	if u.goodsRepo.FindById(goodsId) == nil {
		return errors.New("Incorrect goodsId = " + goodsId)
	}

	order := u.orderRepo.GetFirstActiveOrder(customer)
	if order == nil {
		orderId := uuid.NewV4().String()
		u.orderRepo.CreateOrder(customer, orderId)
		order = u.orderRepo.GetFirstActiveOrder(customer)
	}

	isNotAddedGoods := true
	for _, element := range order.OrderItems {
		if element.GoodsItem.GoodsID == goodsId {
			u.orderRepo.PlusAmount(order.OrderID, goodsId, 1)
			isNotAddedGoods = false
			break
		}
	}

	if isNotAddedGoods {
		u.orderRepo.AddGoods(order.OrderID, goodsId)
	}
	order = u.orderRepo.GetFirstActiveOrder(customer)

	return nil
}

func (u *shoppingCartUseCase) RemoveGoodsFormCartUseCase(customer string, goodsId string) error {
	order := u.orderRepo.GetFirstActiveOrder(customer)
	if order == nil {
		return nil
	}

	for _, element := range order.OrderItems {
		if element.GoodsItem.GoodsID == goodsId {
			u.orderRepo.RemoveGoods(order.OrderID, goodsId)
			break
		}
	}
	return nil
}

func (u *shoppingCartUseCase) ShowCartUseCase(customer string) (*entities.ShoppingCartUseCaseEntity, error) {
	order := u.orderRepo.GetFirstActiveOrder(customer)
	if order == nil {
		return &entities.ShoppingCartUseCaseEntity{Customer: customer, TotalPrice: float64(0), GoodsItems: []entities.ShoppingCartGoodsItemUseCaseEntity{}}, nil
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

	return &entities.ShoppingCartUseCaseEntity{Customer: customer, TotalPrice: totalPrice, GoodsItems: goodsItems}, nil
}
