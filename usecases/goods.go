package usecases

import (
	"errors"
	_repo "github.com/alfssobsd/minishop/dataproviders/postgres"
	_repoEntities "github.com/alfssobsd/minishop/dataproviders/postgres/entities"
	"github.com/alfssobsd/minishop/usecases/entities"
	"github.com/labstack/gommon/log"
	uuid "github.com/satori/go.uuid"
	"github.com/tealeg/xlsx"
)

type GoodsUseCase interface {
	SearchGoodsUseCase() []entities.GoodsUseCaseEntity
	ShowGoodsDetailInfoUseCase(goodsId uuid.UUID) (*entities.GoodsUseCaseEntity, error)
	CreateGoodsUseCase(goodsEntity entities.GoodsUseCaseEntity) entities.GoodsUseCaseEntity
	CreateFromExcelUseCase(pathToExcel string) []entities.GoodsUseCaseEntity
}

type goodsUseCase struct {
	goodsRepository _repo.GoodsRepository
}

func NewGoodsUseCase(goodsRepository _repo.GoodsRepository) *goodsUseCase {
	return &goodsUseCase{goodsRepository}
}

func (u *goodsUseCase) SearchGoodsUseCase() []entities.GoodsUseCaseEntity {
	log.Info("SearchGoodsUseCase")
	goodsEntities := u.goodsRepository.FindAll()

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

func (u *goodsUseCase) ShowGoodsDetailInfoUseCase(goodsId uuid.UUID) (*entities.GoodsUseCaseEntity, error) {
	log.Info("ShowGoodsDetailInfoUseCase id = ", goodsId.String())

	goodsEntity := u.goodsRepository.FindById(goodsId)
	if goodsEntity == nil {
		return nil, errors.New("Not found goods = " + goodsId.String())
	}
	return &entities.GoodsUseCaseEntity{
		GoodsId:         goodsEntity.GoodsID,
		GoodsPrice:      goodsEntity.GoodsPrice,
		GoodsDescrition: goodsEntity.GoodsDescrition,
		GoodsTitle:      goodsEntity.GoodsTitle,
		GoodsCodeName:   goodsEntity.GoodsCodeName,
	}, nil
}

func (u *goodsUseCase) CreateGoodsUseCase(goodsEntity entities.GoodsUseCaseEntity) entities.GoodsUseCaseEntity {
	goodsId := uuid.NewV4()
	log.Info("CreateGoodsUseCase id = ", goodsId.String())
	u.goodsRepository.CreateOne(_repoEntities.GoodsEntity{
		GoodsID:         goodsId,
		GoodsDescrition: goodsEntity.GoodsDescrition,
		GoodsCodeName:   goodsEntity.GoodsCodeName,
		GoodsPrice:      goodsEntity.GoodsPrice,
		GoodsTitle:      goodsEntity.GoodsTitle,
	})

	goodsResultEntity := u.goodsRepository.FindById(goodsId)
	return entities.GoodsUseCaseEntity{
		GoodsId:         goodsResultEntity.GoodsID,
		GoodsPrice:      goodsResultEntity.GoodsPrice,
		GoodsDescrition: goodsResultEntity.GoodsDescrition,
		GoodsTitle:      goodsResultEntity.GoodsTitle,
		GoodsCodeName:   goodsResultEntity.GoodsCodeName,
	}
}

func (u *goodsUseCase) CreateFromExcelUseCase(pathToExcel string) []entities.GoodsUseCaseEntity {
	log.Info("CreateFromExcelUseCase")
	xlFile, err := xlsx.OpenFile(pathToExcel)
	if err != nil {
		log.Fatal(err)
	}

	for _, sheet := range xlFile.Sheets {
		for index, row := range sheet.Rows {

			goods := u.goodsRepository.FindByCodeName(row.Cells[0].String())
			if goods != nil {
				log.Info("Goods ", row.Cells[0].String(), " already added")
				continue
			}

			price, err := row.Cells[3].Float()
			if err != nil {
				log.Error("Can't parse pice in row = ", index)
				continue
			}
			u.CreateGoodsUseCase(entities.GoodsUseCaseEntity{
				GoodsCodeName:   row.Cells[0].String(),
				GoodsTitle:      row.Cells[1].String(),
				GoodsDescrition: row.Cells[2].String(),
				GoodsPrice:      price,
			})
		}
	}

	return u.SearchGoodsUseCase()
}
