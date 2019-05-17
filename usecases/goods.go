package usecases

import (
	_mongoRepsitories "github.com/alfssobsd/minishop/dataproviders/mongodb"
	"github.com/alfssobsd/minishop/usecases/entities"
	"github.com/labstack/gommon/log"
	"gopkg.in/mgo.v2"
)

func SearchGoodsUseCase(db *mgo.Session) []entities.GoodsUseCaseEntity {
	log.Info("SearchGoodsUseCase")
	goodsEntities := _mongoRepsitories.FindAllGoods(db)

	resultEntities := []entities.GoodsUseCaseEntity{}
	for _, element := range goodsEntities {
		resultEntities = append(resultEntities,
			entities.GoodsUseCaseEntity{
				GoodsId:         element.GoodsID,
				GoodsTitle:      element.GoodsTitle,
				GoodsCodeName:   element.GoodsCodeName,
				GoodsDescrition: element.GoodsDescrition,
				GoodsPrice:      element.GoodsPrice,
			})
	}

	return resultEntities
}
func ShowGoodsDetailInfoUseCase(id string) string {
	log.Info("ShowGoodsDetailInfoUseCase id = ", id)
	return "Detail Info"
}
