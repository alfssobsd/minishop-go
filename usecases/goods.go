package usecases

import (
	_mongoRepsitories "github.com/alfssobsd/minishop/dataproviders/mongodb"
	_repoEntities "github.com/alfssobsd/minishop/dataproviders/mongodb/entities"
	"github.com/alfssobsd/minishop/usecases/entities"
	"github.com/labstack/gommon/log"
	uuid "github.com/satori/go.uuid"
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
func ShowGoodsDetailInfoUseCase(db *mgo.Session, id string) entities.GoodsUseCaseEntity {
	log.Info("ShowGoodsDetailInfoUseCase id = ", id)

	goodsItem := _mongoRepsitories.FindGoodsById(db, id)

	return entities.GoodsUseCaseEntity{
		GoodsId:         goodsItem.GoodsID,
		GoodsPrice:      goodsItem.GoodsPrice,
		GoodsDescrition: goodsItem.GoodsDescrition,
		GoodsTitle:      goodsItem.GoodsTitle,
		GoodsCodeName:   goodsItem.GoodsCodeName,
	}
}

func CreateGoodsUseCase(db *mgo.Session, goodsEntity entities.GoodsUseCaseEntity) entities.GoodsUseCaseEntity {
	id := uuid.NewV4().String()
	log.Info("CreateGoodsUseCase id = ", id)
	_mongoRepsitories.CreateGoods(db, _repoEntities.GoodsEntity{
		GoodsID:         id,
		GoodsDescrition: goodsEntity.GoodsDescrition,
		GoodsCodeName:   goodsEntity.GoodsCodeName,
		GoodsPrice:      goodsEntity.GoodsPrice,
		GoodsTitle:      goodsEntity.GoodsTitle,
	})

	goodsItem := _mongoRepsitories.FindGoodsById(db, id)
	return entities.GoodsUseCaseEntity{
		GoodsId:         goodsItem.GoodsID,
		GoodsPrice:      goodsItem.GoodsPrice,
		GoodsDescrition: goodsItem.GoodsDescrition,
		GoodsTitle:      goodsItem.GoodsTitle,
		GoodsCodeName:   goodsItem.GoodsCodeName,
	}
}
