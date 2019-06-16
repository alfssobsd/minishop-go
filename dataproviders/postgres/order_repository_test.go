package postgres

import (
	"fmt"
	"github.com/alfssobsd/minishop/config"
	uuid "github.com/satori/go.uuid"
	"testing"
)

func TestOrderRepository_CreateOrder(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	db := config.MakePostgresConnection()
	orderRepository := NewOrderRepository(db)
	customerRepository := NewCustomerRepository(db)
	customer, _ := customerRepository.FindByUsername("sergei")
	orderRepository.CreateOrder(customer.CustomerId, uuid.FromStringOrNil("2d98c5f9-2a4c-4286-921a-1c2a7c92a452"))
}

func TestOrderRepository_GetFirstActiveOrder(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	db := config.MakePostgresConnection()
	orderRepository := NewOrderRepository(db)
	customerRepository := NewCustomerRepository(db)
	customer, _ := customerRepository.FindByUsername("sergei")
	order := orderRepository.GetFirstActiveOrder(customer.CustomerId)
	fmt.Println(order)
}

func TestOrderRepository_AddGoods(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	db := config.MakePostgresConnection()
	orderRepository := NewOrderRepository(db)
	customerRepository := NewCustomerRepository(db)
	customer, _ := customerRepository.FindByUsername("sergei")
	order := orderRepository.GetFirstActiveOrder(customer.CustomerId)
	orderRepository.AddGoods(order.OrderID, uuid.FromStringOrNil("2d98c5f9-2a4c-4286-921a-1c2a7c92a451"))
	order = orderRepository.GetFirstActiveOrder(customer.CustomerId)
	fmt.Println(order)
}

func TestOrderRepository_RemoveGoods(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	db := config.MakePostgresConnection()
	orderRepository := NewOrderRepository(db)
	customerRepository := NewCustomerRepository(db)
	customer, _ := customerRepository.FindByUsername("sergei")
	order := orderRepository.GetFirstActiveOrder(customer.CustomerId)
	orderRepository.RemoveGoods(order.OrderID, uuid.FromStringOrNil("2d98c5f9-2a4c-4286-921a-1c2a7c92a451"))
	order = orderRepository.GetFirstActiveOrder(customer.CustomerId)
	fmt.Println(order)
}
