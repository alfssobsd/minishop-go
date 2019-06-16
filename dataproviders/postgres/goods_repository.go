package postgres

import (
	"github.com/alfssobsd/minishop/dataproviders/postgres/entities"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
)

type GoodsRepository interface {
	FindAll() []*entities.GoodsEntity
	CreateOne(entities.GoodsEntity)
	FindById(goodsId uuid.UUID) *entities.GoodsEntity
	FindByCodeName(codeName string) *entities.GoodsEntity
}

type goodsRepository struct {
	db *sqlx.DB
}

func NewGoodsRepository(db *sqlx.DB) *goodsRepository {
	return &goodsRepository{db: db}
}

func (r *goodsRepository) CreateOne(goodsEntity entities.GoodsEntity) {
	log.Info("CreateGoods ", goodsEntity)
	result, err := r.db.NamedExec("INSERT INTO goods (uuid, title, code_name, description, price) VALUES (:uuid, :title, :code_name, :description, :price)", &goodsEntity)
	if err != nil {
		log.Error(err)
	}
	log.Info(result)
}

func (r *goodsRepository) FindById(goodsId uuid.UUID) *entities.GoodsEntity {
	log.Info("FindById ", goodsId)

	goodsItem := entities.GoodsEntity{}
	err := r.db.Get(&goodsItem, "SELECT * FROM goods WHERE uuid=$1", goodsId)
	if err != nil {
		log.Error(err)
		return nil
	}
	return &goodsItem
}

func (r *goodsRepository) FindByCodeName(codeName string) *entities.GoodsEntity {
	log.Info("FindByCodeName ", codeName)
	goodsItem := entities.GoodsEntity{}
	err := r.db.Get(&goodsItem, "SELECT * FROM goods WHERE code_name=$1", codeName)
	if err != nil {
		log.Error(err)
		return nil
	}
	return &goodsItem
}

func (r *goodsRepository) FindAll() []*entities.GoodsEntity {
	log.Info("FindAll ")
	var goodsEntities []*entities.GoodsEntity
	err := r.db.Select(&goodsEntities, "SELECT * FROM goods ORDER BY code_name ASC")
	if err != nil {
		log.Error(err)
	}

	return goodsEntities
}
