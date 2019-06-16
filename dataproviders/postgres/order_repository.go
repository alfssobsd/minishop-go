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
	PlusAmount(orderId uuid.UUID, goodsId uuid.UUID, amount int)
	MinusAmount(orderId uuid.UUID, goodsId uuid.UUID, amount int)
	AddGoods(orderId uuid.UUID, goodsId uuid.UUID)
	RemoveGoods(orderId uuid.UUID, goodsId uuid.UUID)
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

	var orderGoods []*entities.OrderGoodsEntity
	log.Info("Get list goods by orderID = ", order.OrderID)
	err = r.db.Select(&orderGoods, "SELECT * from order_goods where order_uuid=$1", order.OrderID)
	if err != nil {
		log.Error(err)
	}

	//need optimisation!!!
	log.Info("Get goods info by orderID = ", order.OrderID)
	for _, element := range orderGoods {
		log.Info("Search goodsID = ", element.GoodsId)
		goodsItem := entities.GoodsEntity{}
		//need handling error!!!
		err := r.db.Get(&goodsItem, "SELECT * FROM goods WHERE uuid=$1", element.GoodsId)
		if err == nil {
			log.Info("Found goods = ", goodsItem)
			order.OrderItems = append(order.OrderItems, entities.OrderItemEntity{
				GoodsItem:   goodsItem,
				GoodsAmount: element.GoodsAmount})
		}
	}
	return &order
}

func (r *orderRepository) PlusAmount(orderId uuid.UUID, goodsId uuid.UUID, amount int) {
	log.Info("PlusAmount  goods = ", goodsId, " to order = ", orderId, " amount = ", strconv.Itoa(amount))
	_ = r.db.MustExec("UPDATE order_goods SET goods_amount = goods_amount + $1  WHERE goods_uuid = $2 AND order_uuid = $3", amount, goodsId, orderId)
}

func (r *orderRepository) MinusAmount(orderId uuid.UUID, goodsId uuid.UUID, amount int) {
	log.Info("PlusAmount  goods = ", goodsId, " to order = ", orderId, " amount = ", strconv.Itoa(amount))
	_ = r.db.MustExec("UPDATE order_goods SET goods_amount = goods_amount - $1  WHERE goods_uuid = $2 AND order_uuid = $3", amount, goodsId, orderId)
}

func (r *orderRepository) AddGoods(orderId uuid.UUID, goodsId uuid.UUID) {
	//need refactoring, before add goods need check order status
	log.Info("Add goods = ", goodsId, " to order = ", orderId)
	_ = r.db.MustExec("INSERT INTO order_goods (goods_uuid, goods_amount, order_uuid) VALUES ($1, $2, $3) ON CONFLICT (goods_uuid, order_uuid) DO NOTHING", goodsId, 1, orderId)
}

func (r *orderRepository) RemoveGoods(orderId uuid.UUID, goodsId uuid.UUID) {
	//need refactoring, before add goods need check order status
	log.Info("Remove goods = ", goodsId, " from order = ", orderId)
	_ = r.db.MustExec("DELETE FROM order_goods WHERE goods_uuid = $1 AND order_uuid = $2", goodsId, orderId)
}
