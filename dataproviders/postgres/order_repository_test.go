package postgres

import (
	"fmt"
	"github.com/alfssobsd/minishop/config"
	"testing"
)

func TestOrderRepository_CreateOrder(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	db := config.MakePostgresConnection()
	orderRepository := NewOrderRepository(db)
	orderRepository.CreateOrder("sergei", "2d98c5f9-2a4c-4286-921a-1c2a7c92a452")
}

func TestOrderRepository_GetFirstActiveOrder(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	db := config.MakePostgresConnection()
	orderRepository := NewOrderRepository(db)
	order := orderRepository.GetFirstActiveOrder("sergei")
	fmt.Println(order)
}

func TestOrderRepository_AddGoods(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	db := config.MakePostgresConnection()
	orderRepository := NewOrderRepository(db)
	order := orderRepository.GetFirstActiveOrder("sergei")
	orderRepository.AddGoods(order.OrderID, "2d98c5f9-2a4c-4286-921a-1c2a7c92a451")
	order = orderRepository.GetFirstActiveOrder("sergei")
	fmt.Println(order)
}

func TestOrderRepository_RemoveGoods(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	db := config.MakePostgresConnection()
	orderRepository := NewOrderRepository(db)
	order := orderRepository.GetFirstActiveOrder("sergei")
	orderRepository.RemoveGoods(order.OrderID, "2d98c5f9-2a4c-4286-921a-1c2a7c92a451")
	order = orderRepository.GetFirstActiveOrder("sergei")
	fmt.Println(order)
}
