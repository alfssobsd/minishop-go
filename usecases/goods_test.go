package usecases

import (
	_repoEntities "github.com/alfssobsd/minishop/dataproviders/mongodb/entities"
	"github.com/alfssobsd/minishop/usecases/entities"
	"github.com/stretchr/testify/assert"
	//"github.com/stretchr/testify/mock"
	"testing"
)

type MockRepo struct {
}

func (m MockRepo) FindById(id string) *_repoEntities.GoodsEntity {
	return &_repoEntities.GoodsEntity{
		GoodsID:         "c26e7e02-c1de-465d-88ff-b845abdc47f1",
		GoodsCodeName:   "0001",
		GoodsTitle:      "Плющевый медведь",
		GoodsDescrition: "Милый плющевый медведь",
		GoodsPrice:      255.5,
	}
}

func (m MockRepo) CreateOne(_repoEntities.GoodsEntity) {
}
func (m MockRepo) FindAll() []*_repoEntities.GoodsEntity {
	return []*_repoEntities.GoodsEntity{}
}

func (m MockRepo) FindByCodeName(codeName string) *_repoEntities.GoodsEntity {
	return &_repoEntities.GoodsEntity{}
}

func TestCorrectReturnFormatShowGoodsDetailInfoUseCase(t *testing.T) {

	goodsUC := NewGoodsUseCase(&MockRepo{})

	// call the code we are testing
	assert.Equal(t, goodsUC.ShowGoodsDetailInfoUseCase("c26e7e02-c1de-465d-88ff-b845abdc47f1"), entities.GoodsUseCaseEntity{
		GoodsId:         "c26e7e02-c1de-465d-88ff-b845abdc47f1",
		GoodsCodeName:   "0001",
		GoodsTitle:      "Плющевый медведь",
		GoodsDescrition: "Милый плющевый медведь",
		GoodsPrice:      255.5,
	})

}
