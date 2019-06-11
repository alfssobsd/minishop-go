package mongodb

import (
	"github.com/alfssobsd/minishop/dataproviders/mongodb/entities"
	"github.com/labstack/gommon/log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type GoodsRepository struct {
	mgoSession *mgo.Session
}

func NewGoodsRepository(mgoSession *mgo.Session) *GoodsRepository {
	return &GoodsRepository{mgoSession: mgoSession}
}

func (repository *GoodsRepository) CreateOne(goodsEntity entities.GoodsEntity) {
	log.Info("CreateGoods ", goodsEntity)
	mgoSession := repository.mgoSession.Clone()
	defer mgoSession.Close()
	_ = mgoSession.DB("minishop").C("goods").Insert(goodsEntity)
}

func (repository *GoodsRepository) FindById(id string) *entities.GoodsEntity {
	log.Info("FindById ", id)
	mgoSession := repository.mgoSession.Clone()
	defer mgoSession.Close()

	goodsItem := &entities.GoodsEntity{}
	_ = mgoSession.DB("minishop").C("goods").FindId(id).One(goodsItem)

	return goodsItem
}

func (repository *GoodsRepository) FindByCodeName(codeName string) *entities.GoodsEntity {
	log.Info("FindByCodeName ", codeName)
	mgoSession := repository.mgoSession.Clone()
	defer mgoSession.Close()

	goodsItem := &entities.GoodsEntity{}
	_ = mgoSession.DB("minishop").C("goods").Find(bson.M{"goods_code_name": codeName}).One(goodsItem)

	return goodsItem
}

func (repository *GoodsRepository) FindAll() []*entities.GoodsEntity {
	log.Info("FindAll ")
	mgoSession := repository.mgoSession.Clone()
	defer mgoSession.Close()

	var goodsEntities []*entities.GoodsEntity
	_ = mgoSession.DB("minishop").C("goods").Find(nil).All(&goodsEntities)

	return goodsEntities
}
