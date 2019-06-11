package mongodb

import (
	"github.com/alfssobsd/minishop/dataproviders/mongodb/entities"
	"github.com/labstack/gommon/log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type GoodsRepository interface {
	FindAll() []*entities.GoodsEntity
	CreateOne(entities.GoodsEntity)
	FindById(id string) *entities.GoodsEntity
	FindByCodeName(codeName string) *entities.GoodsEntity
}

type goodsRepository struct {
	mgoSession *mgo.Session
}

func NewGoodsRepository(mgoSession *mgo.Session) *goodsRepository {
	return &goodsRepository{mgoSession: mgoSession}
}

func (repository *goodsRepository) CreateOne(goodsEntity entities.GoodsEntity) {
	log.Info("CreateGoods ", goodsEntity)
	mgoSession := repository.mgoSession.Clone()
	defer mgoSession.Close()
	_ = mgoSession.DB("minishop").C("goods").Insert(goodsEntity)
}

func (repository *goodsRepository) FindById(id string) *entities.GoodsEntity {
	log.Info("FindById ", id)
	mgoSession := repository.mgoSession.Clone()
	defer mgoSession.Close()

	goodsItem := &entities.GoodsEntity{}
	_ = mgoSession.DB("minishop").C("goods").FindId(id).One(goodsItem)

	return goodsItem
}

func (repository *goodsRepository) FindByCodeName(codeName string) *entities.GoodsEntity {
	log.Info("FindByCodeName ", codeName)
	mgoSession := repository.mgoSession.Clone()
	defer mgoSession.Close()

	goodsItem := &entities.GoodsEntity{}
	_ = mgoSession.DB("minishop").C("goods").Find(bson.M{"goods_code_name": codeName}).One(goodsItem)

	return goodsItem
}

func (repository *goodsRepository) FindAll() []*entities.GoodsEntity {
	log.Info("FindAll ")
	mgoSession := repository.mgoSession.Clone()
	defer mgoSession.Close()

	var goodsEntities []*entities.GoodsEntity
	_ = mgoSession.DB("minishop").C("goods").Find(nil).All(&goodsEntities)

	return goodsEntities
}
