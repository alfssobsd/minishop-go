package usecases

import (
	_mongoRepsitories "github.com/alfssobsd/minishop/dataproviders/mongodb"
	_repoEntities "github.com/alfssobsd/minishop/dataproviders/mongodb/entities"
	"github.com/alfssobsd/minishop/usecases/entities"
	"github.com/labstack/gommon/log"
	uuid "github.com/satori/go.uuid"
)

type GoodsUseCase struct {
	goodsRepository *_mongoRepsitories.GoodsRepository
}

func NewGoodsUseCase(goodsRepository *_mongoRepsitories.GoodsRepository) *GoodsUseCase {
	return &GoodsUseCase{goodsRepository}
}

func (goodsUseCase *GoodsUseCase) SearchGoodsUseCase() []entities.GoodsUseCaseEntity {
	log.Info("SearchGoodsUseCase")
	goodsEntities := goodsUseCase.goodsRepository.FindAll()

	var resultEntities []entities.GoodsUseCaseEntity
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

func (goodsUseCase *GoodsUseCase) ShowGoodsDetailInfoUseCase(id string) entities.GoodsUseCaseEntity {
	log.Info("ShowGoodsDetailInfoUseCase id = ", id)

	goodsEntity := goodsUseCase.goodsRepository.FindById(id)
	return entities.GoodsUseCaseEntity{
		GoodsId:         goodsEntity.GoodsID,
		GoodsPrice:      goodsEntity.GoodsPrice,
		GoodsDescrition: goodsEntity.GoodsDescrition,
		GoodsTitle:      goodsEntity.GoodsTitle,
		GoodsCodeName:   goodsEntity.GoodsCodeName,
	}
}

func (goodsUseCase *GoodsUseCase) CreateGoodsUseCase(goodsEntity entities.GoodsUseCaseEntity) entities.GoodsUseCaseEntity {
	id := uuid.NewV4().String()
	log.Info("CreateGoodsUseCase id = ", id)
	goodsUseCase.goodsRepository.CreateOne(_repoEntities.GoodsEntity{
		GoodsID:         id,
		GoodsDescrition: goodsEntity.GoodsDescrition,
		GoodsCodeName:   goodsEntity.GoodsCodeName,
		GoodsPrice:      goodsEntity.GoodsPrice,
		GoodsTitle:      goodsEntity.GoodsTitle,
	})

	goodsResultEntity := goodsUseCase.goodsRepository.FindById(id)
	return entities.GoodsUseCaseEntity{
		GoodsId:         goodsResultEntity.GoodsID,
		GoodsPrice:      goodsResultEntity.GoodsPrice,
		GoodsDescrition: goodsResultEntity.GoodsDescrition,
		GoodsTitle:      goodsResultEntity.GoodsTitle,
		GoodsCodeName:   goodsResultEntity.GoodsCodeName,
	}
}
