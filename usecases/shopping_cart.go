package usecases

import (
	"errors"
	_repo "github.com/alfssobsd/minishop/dataproviders/postgres"
	"github.com/alfssobsd/minishop/usecases/entities"
	uuid "github.com/satori/go.uuid"
)

type ShoppingCartUseCase interface {
	AddProductToCartUseCase(username string, productId uuid.UUID) error
	RemoveProductFormCartUseCase(username string, productId uuid.UUID) error
	ShowCartUseCase(username string) (*entities.ShoppingCartUseCaseEntity, error)
}

type shoppingCartUseCase struct {
	productRepo _repo.ProductRepository
	orderRepo   _repo.OrderRepository
	custRepo    _repo.CustomerRepository
}

func NewShoppingCartUseCase(productRepository _repo.ProductRepository, orderRepository _repo.OrderRepository, custRepository _repo.CustomerRepository) *shoppingCartUseCase {
	return &shoppingCartUseCase{productRepository, orderRepository, custRepository}
}

func (u *shoppingCartUseCase) AddProductToCartUseCase(username string, productId uuid.UUID) error {
	if u.productRepo.FindById(productId) == nil {
		return errors.New("Incorrect productId = " + productId.String())
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

	isNotAddedProduct := true
	for _, element := range order.OrderItems {
		if element.ProductItem.ProductID == productId {
			u.orderRepo.PlusAmount(order.OrderID, productId, 1)
			isNotAddedProduct = false
			break
		}
	}

	if isNotAddedProduct {
		u.orderRepo.AddProduct(order.OrderID, productId)
	}
	order = u.orderRepo.GetFirstActiveOrder(customer.CustomerId)

	return nil
}

func (u *shoppingCartUseCase) RemoveProductFormCartUseCase(username string, productId uuid.UUID) error {

	customer, err := u.custRepo.FindByUsername(username)
	if err != nil {
		return err
	}

	order := u.orderRepo.GetFirstActiveOrder(customer.CustomerId)
	if order == nil {
		return nil
	}

	for _, element := range order.OrderItems {
		if element.ProductItem.ProductID == productId {
			u.orderRepo.RemoveProduct(order.OrderID, productId)
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
		return &entities.ShoppingCartUseCaseEntity{CustomerId: customer.CustomerId, TotalPrice: float64(0), ProductItems: []entities.ShoppingCartProductItemUseCaseEntity{}}, nil
	}

	totalPrice := float64(0)
	productItems := []entities.ShoppingCartProductItemUseCaseEntity{}
	for _, element := range order.OrderItems {
		productItems = append(productItems, entities.ShoppingCartProductItemUseCaseEntity{
			Product: entities.ProductUseCaseEntity{
				ProductId:         element.ProductItem.ProductID,
				ProductCodeName:   element.ProductItem.ProductCodeName,
				ProductDescrition: element.ProductItem.ProductDescrition,
				ProductTitle:      element.ProductItem.ProductTitle,
				ProductPrice:      element.ProductItem.ProductPrice,
			},
			Amount: element.ProductAmount,
		})
		for i := 0; i < element.ProductAmount; i++ {
			totalPrice += element.ProductItem.ProductPrice
		}
	}

	return &entities.ShoppingCartUseCaseEntity{CustomerId: customer.CustomerId, TotalPrice: totalPrice, ProductItems: productItems}, nil
}
