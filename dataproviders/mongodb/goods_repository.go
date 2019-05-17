package mongodb

import (
	"github.com/alfssobsd/minishop/dataproviders/mongodb/entities"
	uuid "github.com/satori/go.uuid"
	"gopkg.in/mgo.v2"
)

func GetGoodsById(id string) entities.GoodsEntity {
	return entities.GoodsEntity{
		"90e0cabf-101b-45db-8550-c342eabebd73",
		"GOODS01",
		"Плющевый медведь",
		"Милый плюшевый медведь",
		103.22,
	}
}

func FindAllGoods(db *mgo.Session) []*entities.GoodsEntity {

	item := &entities.GoodsEntity{
		GoodsID:         uuid.NewV4().String(),
		GoodsTitle:      "Плющевый медведь",
		GoodsDescrition: "Милый плюшевый медведь",
		GoodsPrice:      111.22,
		GoodsCodeName:   "CODENAME01",
	}

	_ = db.DB("minishop").C("goods").Insert(item)

	goodsEntities := []*entities.GoodsEntity{}

	dbconnect := db.Clone()
	defer dbconnect.Close()
	_ = dbconnect.DB("minishop").C("goods").Find(nil).All(&goodsEntities)
	//goodsEntities := []entities.GoodsEntity{
	//	entities.GoodsEntity{
	//		"90e0cabf-101b-45db-8550-c342eabebd73",
	//		"GOODS01",
	//		"Плющевый медведь",
	//		"Милый плюшевый медведь",
	//		103.22,
	//	},
	//	entities.GoodsEntity{
	//		"90e0cabf-101b-45db-8550-c342eabebd74",
	//		"GOODS02",
	//		"Плющевый жираф",
	//		"Милый плюшевый жираф",
	//		100.23,
	//	},
	//}
	//goodsEntities := []entities.GoodsEntity{}
	return goodsEntities
}

func CreateGoods(CodeName string) {
	uuid.NewV4()
	//	TODO: create goods item
}
