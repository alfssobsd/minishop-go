package mongodb

import (
	mongo_config "github.com/alfssobsd/minishop/config"
	"github.com/alfssobsd/minishop/dataproviders/mongodb/entities"
	"github.com/labstack/gommon/log"
)

func CreateGoods(goodsEntity entities.GoodsEntity) {
	log.Info("CreateGoods ", goodsEntity)
	session := mongo_config.GetMongoSession()
	defer mongo_config.CloseMongoSession(session)
	_ = session.DB("minishop").C("goods").Insert(goodsEntity)
}

func FindGoodsById(id string) *entities.GoodsEntity {
	goodsItem := &entities.GoodsEntity{}

	session := mongo_config.GetMongoSession()
	defer mongo_config.CloseMongoSession(session)
	_ = session.DB("minishop").C("goods").FindId(id).One(goodsItem)

	return goodsItem
}

func FindAllGoods() []*entities.GoodsEntity {
	goodsEntities := []*entities.GoodsEntity{}

	session := mongo_config.GetMongoSession()
	defer mongo_config.CloseMongoSession(session)
	_ = session.DB("minishop").C("goods").Find(nil).All(&goodsEntities)

	return goodsEntities
}
