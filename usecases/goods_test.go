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
	postgres.GoodsRepository
}

func (m MockRepo) FindById(goodsId uuid.UUID) *_repoEntities.GoodsEntity {
	return &_repoEntities.GoodsEntity{
		GoodsID:         uuid.FromStringOrNil("c26e7e02-c1de-465d-88ff-b845abdc47f1"),
		GoodsCodeName:   "0001",
		GoodsTitle:      "Плющевый медведь",
		GoodsDescrition: "Милый плющевый медведь",
		GoodsPrice:      255.5,
	}
}

func TestGoodsUseCase_ShowGoodsDetailInfoUseCase(t *testing.T) {

	goodsUC := NewGoodsUseCase(&MockRepo{})

	goods, _ := goodsUC.ShowGoodsDetailInfoUseCase(uuid.FromStringOrNil("c26e7e02-c1de-465d-88ff-b845abdc47f1"))
	// call the code we are testing
	assert.Equal(t, goods, &entities.GoodsUseCaseEntity{
		GoodsId:         uuid.FromStringOrNil("c26e7e02-c1de-465d-88ff-b845abdc47f1"),
		GoodsCodeName:   "0001",
		GoodsTitle:      "Плющевый медведь",
		GoodsDescrition: "Милый плющевый медведь",
		GoodsPrice:      255.5,
	})

}
