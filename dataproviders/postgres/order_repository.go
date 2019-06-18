package postgres

import (
	"github.com/alfssobsd/minishop/dataproviders/postgres/entities"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
	uuid "github.com/satori/go.uuid"
	"strconv"
)

type OrderRepository interface {
	CreateOrder(customerId uuid.UUID, orderId uuid.UUID)
	GetFirstActiveOrder(customerId uuid.UUID) *entities.OrderEntity
	PlusAmount(orderId uuid.UUID, productId uuid.UUID, amount int)
	MinusAmount(orderId uuid.UUID, productId uuid.UUID, amount int)
	AddProduct(orderId uuid.UUID, productId uuid.UUID)
	RemoveProduct(orderId uuid.UUID, productId uuid.UUID)
}

type orderRepository struct {
	db *sqlx.DB
}

func NewOrderRepository(db *sqlx.DB) *orderRepository {
	return &orderRepository{db: db}
}

func (r *orderRepository) CreateOrder(customerId uuid.UUID, orderId uuid.UUID) {
	log.Info("CreateOrder customerId = ", customerId)
	_ = r.db.MustExec("INSERT INTO orders (uuid, customer_uuid, status) VALUES ($1, $2, $3)", orderId, customerId, entities.OrderStatus(entities.OrderStatusActive))
}

func (r *orderRepository) GetFirstActiveOrder(customerId uuid.UUID) *entities.OrderEntity {
	log.Info("Get Active Order customerId = ", customerId)

	order := entities.OrderEntity{}
	order.OrderItems = []entities.OrderItemEntity{}

	err := r.db.Get(&order, "SELECT * FROM orders WHERE status=$1 and customer_uuid=$2", entities.OrderStatusActive, customerId)
	if err != nil {
		log.Error(err)
		return nil
	}

	var orderProducts []*entities.OrderProductEntity
	log.Info("Get list products by orderID = ", order.OrderID)
	err = r.db.Select(&orderProducts, "SELECT * from order_products where order_uuid=$1", order.OrderID)
	if err != nil {
		log.Error(err)
	}

	//need optimisation!!!
	log.Info("Get products info by orderID = ", order.OrderID)
	for _, element := range orderProducts {
		log.Info("Search productID = ", element.ProductId)
		productItem := entities.ProductEntity{}
		//need handling error!!!
		err := r.db.Get(&productItem, "SELECT * FROM products WHERE uuid=$1", element.ProductId)
		if err == nil {
			log.Info("Found product = ", productItem)
			order.OrderItems = append(order.OrderItems, entities.OrderItemEntity{
				ProductItem:   productItem,
				ProductAmount: element.ProductAmount})
		}
	}
	return &order
}

func (r *orderRepository) PlusAmount(orderId uuid.UUID, productId uuid.UUID, amount int) {
	log.Info("PlusAmount  product = ", productId, " to order = ", orderId, " amount = ", strconv.Itoa(amount))
	_ = r.db.MustExec("UPDATE order_products SET product_amount = product_amount + $1  WHERE product_uuid = $2 AND order_uuid = $3", amount, productId, orderId)
}

func (r *orderRepository) MinusAmount(orderId uuid.UUID, productId uuid.UUID, amount int) {
	log.Info("Minus  product = ", productId, " to order = ", orderId, " amount = ", strconv.Itoa(amount))
	_ = r.db.MustExec("UPDATE order_products SET product_amount = product_amount - $1  WHERE product_uuid = $2 AND order_uuid = $3", amount, productId, orderId)
}

func (r *orderRepository) AddProduct(orderId uuid.UUID, productId uuid.UUID) {
	//need refactoring, before add product need check order status
	log.Info("Add product = ", productId, " to order = ", orderId)
	_ = r.db.MustExec("INSERT INTO order_products (product_uuid, product_amount, order_uuid) VALUES ($1, $2, $3) ON CONFLICT (product_uuid, order_uuid) DO NOTHING", productId, 1, orderId)
}

func (r *orderRepository) RemoveProduct(orderId uuid.UUID, productId uuid.UUID) {
	//need refactoring, before add product need check order status
	log.Info("Remove product = ", productId, " from order = ", orderId)
	_ = r.db.MustExec("DELETE FROM order_products WHERE product_uuid = $1 AND order_uuid = $2", productId, orderId)
}
