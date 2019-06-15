package usecases

import (
	"github.com/alfssobsd/minishop/dataproviders/postgres"
	entities2 "github.com/alfssobsd/minishop/dataproviders/postgres/entities"
	"github.com/alfssobsd/minishop/usecases/entities"
	"github.com/stretchr/testify/assert"
	"testing"
)

type MockOrderRepo struct {
	postgres.OrderRepository
}

type MockGoodsRepo struct {
	postgres.GoodsRepository
}

func (r *MockOrderRepo) GetFirstActiveOrder(customer string) *entities2.OrderEntity {
	return nil
}

func TestShoppingCartUseCase_ShowCartUseCase_NoOneOrder(t *testing.T) {
	cartUC := NewShoppingCartUseCase(&MockGoodsRepo{}, &MockOrderRepo{})

	cart, _ := cartUC.ShowCartUseCase("sergei")

	assert.Equal(t, cart, &entities.ShoppingCartUseCaseEntity{Customer: "sergei", TotalPrice: float64(0), GoodsItems: []entities.ShoppingCartGoodsItemUseCaseEntity{}})
}

type MockOrderRepoEmptyOrder struct {
	postgres.OrderRepository
}

func (r *MockOrderRepoEmptyOrder) GetFirstActiveOrder(customer string) *entities2.OrderEntity {
	return &entities2.OrderEntity{OrderID: "cadceb09-308a-4c60-9def-94c435e77be3",
		OrderStatus: entities2.OrderStatusActive, OrderCustomer: "sergei", OrderItems: []entities2.OrderItemEntity{}}
}

func TestShoppingCartUseCase_ShowCartUseCase_EmptyOrder(t *testing.T) {
	cartUC := NewShoppingCartUseCase(&MockGoodsRepo{}, &MockOrderRepoEmptyOrder{})

	cart, _ := cartUC.ShowCartUseCase("sergei")
	assert.Equal(t, cart, &entities.ShoppingCartUseCaseEntity{Customer: "sergei", TotalPrice: float64(0),
		GoodsItems: []entities.ShoppingCartGoodsItemUseCaseEntity{}})
}

type MockOrderRepoOneItemTwoAmount struct {
	postgres.OrderRepository
}

func (r *MockOrderRepoOneItemTwoAmount) GetFirstActiveOrder(customer string) *entities2.OrderEntity {
	items := []entities2.OrderItemEntity{}
	items = append(items, entities2.OrderItemEntity{
		GoodsItem: entities2.GoodsEntity{
			GoodsID:         "2d98c5f9-2a4c-4286-921a-1c2a7c92a451",
			GoodsCodeName:   "TOY04",
			GoodsTitle:      "Little Toy1",
			GoodsDescrition: "Description L Toy1",
			GoodsPrice:      10.3,
		},
		GoodsAmount: 2,
	})
	return &entities2.OrderEntity{OrderID: "cadceb09-308a-4c60-9def-94c435e77be3",
		OrderStatus: entities2.OrderStatusActive, OrderCustomer: "sergei", OrderItems: items}
}

func TestShoppingCartUseCase_ShowCartUseCase_OneItemTwoAmount(t *testing.T) {
	cartUC := NewShoppingCartUseCase(&MockGoodsRepo{}, &MockOrderRepoOneItemTwoAmount{})

	cart, _ := cartUC.ShowCartUseCase("sergei")
	items := []entities.ShoppingCartGoodsItemUseCaseEntity{}
	items = append(items, entities.ShoppingCartGoodsItemUseCaseEntity{
		Goods: entities.GoodsUseCaseEntity{
			GoodsId:         "2d98c5f9-2a4c-4286-921a-1c2a7c92a451",
			GoodsCodeName:   "TOY04",
			GoodsTitle:      "Little Toy1",
			GoodsDescrition: "Description L Toy1",
			GoodsPrice:      10.3,
		},
		Amount: 2,
	})
	assert.Equal(t, cart, &entities.ShoppingCartUseCaseEntity{Customer: "sergei", TotalPrice: 20.6,
		GoodsItems: items})
}
