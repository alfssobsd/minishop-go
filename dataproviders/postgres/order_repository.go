package postgres

import (
	"github.com/alfssobsd/minishop/dataproviders/postgres/entities"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
	"strconv"
)

type OrderRepository interface {
	CreateOrder(customer string, uuid string)
	GetFirstActiveOrder(customer string) *entities.OrderEntity
	PlusAmount(orderId string, goodsId string, amount int)
	MinusAmount(orderId string, goodsId string, amount int)
	AddGoods(orderId string, goodsId string)
	RemoveGoods(orderId string, goodsId string)
}

type orderRepository struct {
	db *sqlx.DB
}

func NewOrderRepository(db *sqlx.DB) *orderRepository {
	return &orderRepository{db: db}
}

func (r *orderRepository) CreateOrder(customer string, uuid string) {
	log.Info("CreateOrder customer = ", customer)
	_ = r.db.MustExec("INSERT INTO orders (uuid, customer, status) VALUES ($1, $2, $3)", uuid, customer, entities.OrderStatus(entities.OrderStatusActive))
}

func (r *orderRepository) GetFirstActiveOrder(customer string) *entities.OrderEntity {
	log.Info("Get Active Order customer = ", customer)

	order := entities.OrderEntity{}
	order.OrderItems = []entities.OrderItemEntity{}

	err := r.db.Get(&order, "SELECT * FROM orders WHERE status=$1 and customer=$2", entities.OrderStatusActive, customer)
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

func (r *orderRepository) PlusAmount(orderId string, goodsId string, amount int) {
	log.Info("PlusAmount  goods = ", goodsId, " to order = ", orderId, " amount = ", strconv.Itoa(amount))
	_ = r.db.MustExec("UPDATE order_goods SET goods_amount = goods_amount + $1  WHERE goods_uuid = $2 AND order_uuid = $3", amount, goodsId, orderId)
}

func (r *orderRepository) MinusAmount(orderId string, goodsId string, amount int) {
	log.Info("PlusAmount  goods = ", goodsId, " to order = ", orderId, " amount = ", strconv.Itoa(amount))
	_ = r.db.MustExec("UPDATE order_goods SET goods_amount = goods_amount - $1  WHERE goods_uuid = $2 AND order_uuid = $3", amount, goodsId, orderId)
}

func (r *orderRepository) AddGoods(orderId string, goodsId string) {
	//need refactoring, before add goods need check order status
	log.Info("Add goods = ", goodsId, " to order = ", orderId)
	_ = r.db.MustExec("INSERT INTO order_goods (goods_uuid, goods_amount, order_uuid) VALUES ($1, $2, $3) ON CONFLICT (goods_uuid, order_uuid) DO NOTHING", goodsId, 1, orderId)
}

func (r *orderRepository) RemoveGoods(orderId string, goodsId string) {
	//need refactoring, before add goods need check order status
	log.Info("Remove goods = ", goodsId, " from order = ", orderId)
	_ = r.db.MustExec("DELETE FROM order_goods WHERE goods_uuid = $1 AND order_uuid = $2", goodsId, orderId)
}
