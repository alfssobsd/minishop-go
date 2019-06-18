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

type ProductUseCase interface {
	SearchProductsUseCase() []entities.ProductUseCaseEntity
	ShowProductDetailInfoUseCase(productId uuid.UUID) (*entities.ProductUseCaseEntity, error)
	CreateProductUseCase(productEntity entities.ProductUseCaseEntity) entities.ProductUseCaseEntity
	CreateProductFromExcelUseCase(pathToExcel string) []entities.ProductUseCaseEntity
}

type productUseCase struct {
	productRepository _repo.ProductRepository
}

func NewProductUseCase(productRepository _repo.ProductRepository) *productUseCase {
	return &productUseCase{productRepository}
}

func (u *productUseCase) SearchProductsUseCase() []entities.ProductUseCaseEntity {
	log.Info("SearchProductsUseCase")
	productEntities := u.productRepository.FindAll()

	var resultEntities []entities.ProductUseCaseEntity
	for _, element := range productEntities {
		resultEntities = append(resultEntities,
			entities.ProductUseCaseEntity{
				ProductId:         element.ProductID,
				ProductTitle:      element.ProductTitle,
				ProductCodeName:   element.ProductCodeName,
				ProductDescrition: element.ProductDescrition,
				ProductPrice:      element.ProductPrice,
			})
	}

	return resultEntities
}

func (u *productUseCase) ShowProductDetailInfoUseCase(productId uuid.UUID) (*entities.ProductUseCaseEntity, error) {
	log.Info("ShowProductDetailInfoUseCase id = ", productId.String())

	productEntity := u.productRepository.FindById(productId)
	if productEntity == nil {
		return nil, errors.New("Not found product = " + productId.String())
	}
	return &entities.ProductUseCaseEntity{
		ProductId:         productEntity.ProductID,
		ProductPrice:      productEntity.ProductPrice,
		ProductDescrition: productEntity.ProductDescrition,
		ProductTitle:      productEntity.ProductTitle,
		ProductCodeName:   productEntity.ProductCodeName,
	}, nil
}

func (u *productUseCase) CreateProductUseCase(peroductEntity entities.ProductUseCaseEntity) entities.ProductUseCaseEntity {
	productId := uuid.NewV4()
	log.Info("CreateProductUseCase id = ", productId.String())
	u.productRepository.CreateOne(_repoEntities.ProductEntity{
		ProductID:         productId,
		ProductDescrition: peroductEntity.ProductDescrition,
		ProductCodeName:   peroductEntity.ProductCodeName,
		ProductPrice:      peroductEntity.ProductPrice,
		ProductTitle:      peroductEntity.ProductTitle,
	})

	resultEntity := u.productRepository.FindById(productId)
	return entities.ProductUseCaseEntity{
		ProductId:         resultEntity.ProductID,
		ProductPrice:      resultEntity.ProductPrice,
		ProductDescrition: resultEntity.ProductDescrition,
		ProductTitle:      resultEntity.ProductTitle,
		ProductCodeName:   resultEntity.ProductCodeName,
	}
}

func (u *productUseCase) CreateProductFromExcelUseCase(pathToExcel string) []entities.ProductUseCaseEntity {
	log.Info("CreateProductFromExcelUseCase")
	xlFile, err := xlsx.OpenFile(pathToExcel)
	if err != nil {
		log.Fatal(err)
	}

	for _, sheet := range xlFile.Sheets {
		for index, row := range sheet.Rows {

			product := u.productRepository.FindByCodeName(row.Cells[0].String())
			if product != nil {
				log.Info("Product ", row.Cells[0].String(), " already added")
				continue
			}

			price, err := row.Cells[3].Float()
			if err != nil {
				log.Error("Can't parse pice in row = ", index)
				continue
			}
			u.CreateProductUseCase(entities.ProductUseCaseEntity{
				ProductCodeName:   row.Cells[0].String(),
				ProductTitle:      row.Cells[1].String(),
				ProductDescrition: row.Cells[2].String(),
				ProductPrice:      price,
			})
		}
	}

	return u.SearchProductsUseCase()
}
