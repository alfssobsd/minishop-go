package usecases

import (
	"errors"
	_repo "github.com/alfssobsd/minishop/dataproviders/postgres"
	"github.com/alfssobsd/minishop/usecases/entities"
	uuid "github.com/satori/go.uuid"
)

type ShoppingCartUseCase interface {
	AddGoodsToCartUseCase(username string, goodsId uuid.UUID) error
	RemoveGoodsFormCartUseCase(username string, goodsId uuid.UUID) error
	ShowCartUseCase(username string) (*entities.ShoppingCartUseCaseEntity, error)
}

type shoppingCartUseCase struct {
	goodsRepo _repo.GoodsRepository
	orderRepo _repo.OrderRepository
	custRepo  _repo.CustomerRepository
}

func NewShoppingCartUseCase(goodsRepository _repo.GoodsRepository, orderRepository _repo.OrderRepository, custRepository _repo.CustomerRepository) *shoppingCartUseCase {
	return &shoppingCartUseCase{goodsRepository, orderRepository, custRepository}
}

func (u *shoppingCartUseCase) AddGoodsToCartUseCase(username string, goodsId uuid.UUID) error {
	if u.goodsRepo.FindById(goodsId) == nil {
		return errors.New("Incorrect goodsId = " + goodsId.String())
	}

	customer, err := u.custRepo.FindByUsername(username)
	if err != nil {
		return err
	}

	order := u.orderRepo.GetFirstActiveOrder(customer.CustomerId)
	if order == nil {
		orderId := uuid.NewV4()
		u.orderRepo.CreateOrder(customer.CustomerId, orderId)
		order = u.orderRepo.GetFirstActiveOrder(customer.CustomerId)
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
	order = u.orderRepo.GetFirstActiveOrder(customer.CustomerId)

	return nil
}

func (u *shoppingCartUseCase) RemoveGoodsFormCartUseCase(username string, goodsId uuid.UUID) error {

	customer, err := u.custRepo.FindByUsername(username)
	if err != nil {
		return err
	}

	order := u.orderRepo.GetFirstActiveOrder(customer.CustomerId)
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

func (u *shoppingCartUseCase) ShowCartUseCase(username string) (*entities.ShoppingCartUseCaseEntity, error) {
	customer, err := u.custRepo.FindByUsername(username)
	if err != nil {
		return nil, err
	}

	order := u.orderRepo.GetFirstActiveOrder(customer.CustomerId)

	if order == nil {
		return &entities.ShoppingCartUseCaseEntity{CustomerId: customer.CustomerId, TotalPrice: float64(0), GoodsItems: []entities.ShoppingCartGoodsItemUseCaseEntity{}}, nil
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

	return &entities.ShoppingCartUseCaseEntity{CustomerId: customer.CustomerId, TotalPrice: totalPrice, GoodsItems: goodsItems}, nil
}
