package usecases

import (
	"github.com/alfssobsd/minishop/dataproviders/postgres"
	_repoEntities "github.com/alfssobsd/minishop/dataproviders/postgres/entities"
	"github.com/alfssobsd/minishop/usecases/entities"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

type MockRepo struct {
	postgres.ProductRepository
}

func (m MockRepo) FindById(productId uuid.UUID) *_repoEntities.ProductEntity {
	return &_repoEntities.ProductEntity{
		ProductID:         uuid.FromStringOrNil("c26e7e02-c1de-465d-88ff-b845abdc47f1"),
		ProductCodeName:   "0001",
		ProductTitle:      "Плющевый медведь",
		ProductDescrition: "Милый плющевый медведь",
		ProductPrice:      255.5,
	}
}

func TestGoodsUseCase_ShowProductDetailInfoUseCase(t *testing.T) {
	productUseCase := NewProductUseCase(&MockRepo{})
	product, _ := productUseCase.ShowProductDetailInfoUseCase(uuid.FromStringOrNil("c26e7e02-c1de-465d-88ff-b845abdc47f1"))
	// call the code we are testing
	assert.Equal(t, product, &entities.ProductUseCaseEntity{
		ProductId:         uuid.FromStringOrNil("c26e7e02-c1de-465d-88ff-b845abdc47f1"),
		ProductCodeName:   "0001",
		ProductTitle:      "Плющевый медведь",
		ProductDescrition: "Милый плющевый медведь",
		ProductPrice:      255.5,
	})

}
