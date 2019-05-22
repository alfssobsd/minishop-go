package usecases

import (
	entities2 "github.com/alfssobsd/minishop/dataproviders/mongodb/entities"
	"github.com/alfssobsd/minishop/usecases/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MyMockedObject struct {
	mock.Mock
}

func (m *MyMockedObject) FindGoodsById(id string) entities2.GoodsEntity {
	args := m.Called(id)
	return args.Get(0).(entities2.GoodsEntity)
}

func TestCorrectReturnFormatShowGoodsDetailInfoUseCase(t *testing.T) {
	testObj := new(MyMockedObject)

	testObj.On("FindGoodsById", "c26e7e02-c1de-465d-88ff-b845abdc47f1").Return(entities2.GoodsEntity{
		GoodsID:         "c26e7e02-c1de-465d-88ff-b845abdc47f1",
		GoodsCodeName:   "0001",
		GoodsTitle:      "Плющевый медведь",
		GoodsDescrition: "Милый плющевый медведь",
		GoodsPrice:      255.5,
	})

	// call the code we are testing
	assert.Equal(t, ShowGoodsDetailInfoUseCase("c26e7e02-c1de-465d-88ff-b845abdc47f1"), entities.GoodsUseCaseEntity{
		GoodsId:         "c26e7e02-c1de-465d-88ff-b845abdc47f1",
		GoodsCodeName:   "0001",
		GoodsTitle:      "Плющевый медведь",
		GoodsDescrition: "Милый плющевый медведь",
		GoodsPrice:      255.5,
	})

}
