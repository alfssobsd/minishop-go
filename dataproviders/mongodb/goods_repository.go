package mongodb

import (
	"github.com/alfssobsd/minishop/dataproviders/mongodb/entities"
	"github.com/labstack/gommon/log"
	"gopkg.in/mgo.v2"
)

func CreateGoods(db *mgo.Session, goodsEntity entities.GoodsEntity) {
	log.Info("CreateGoods ", goodsEntity)
	dbconnect := db.Clone()
	defer dbconnect.Close()
	_ = dbconnect.DB("minishop").C("goods").Insert(goodsEntity)
}

func FindGoodsById(db *mgo.Session, id string) *entities.GoodsEntity {
	goodsItem := &entities.GoodsEntity{}

	dbconnect := db.Clone()
	defer dbconnect.Close()
	_ = dbconnect.DB("minishop").C("goods").FindId(id).One(goodsItem)

	return goodsItem
}

func FindAllGoods(db *mgo.Session) []*entities.GoodsEntity {
	goodsEntities := []*entities.GoodsEntity{}

	dbconnect := db.Clone()
	defer dbconnect.Close()
	_ = dbconnect.DB("minishop").C("goods").Find(nil).All(&goodsEntities)

	return goodsEntities
}
