package http

import (
	"github.com/alfssobsd/minishop/usecases/entities"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MyMockedObject struct {
	mock.Mock
}

func (m *MyMockedObject) ShowGoodsDetailInfoUseCase(id string) entities.GoodsUseCaseEntity {
	args := m.Called(id)
	return args.Get(0).(entities.GoodsUseCaseEntity)
}

var showGoodDetailJSON = `{"goods_id":"c26e7e02-c1de-465d-88ff-b845abdc47f1","goods_code_name":"0001","goods_title":"Плющевый медведь","goods_description":"Милый плющевый медведь","goods_price":255.5}`

func TestShowGoodsDetailInfoController(t *testing.T) {
	testObj := new(MyMockedObject)

	testObj.On("ShowGoodsDetailInfoUseCase", "c26e7e02-c1de-465d-88ff-b845abdc47f1").Return(entities.GoodsUseCaseEntity{
		GoodsId:         "c26e7e02-c1de-465d-88ff-b845abdc47f1",
		GoodsCodeName:   "0001",
		GoodsTitle:      "Плющевый медведь",
		GoodsDescrition: "Милый плющевый медведь",
		GoodsPrice:      255.5,
	})

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/v1/goods/:id")
	c.SetParamNames("id")
	c.SetParamValues("c26e7e02-c1de-465d-88ff-b845abdc47f1")

	if assert.NoError(t, showGoodsDetailInfoController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, showGoodDetailJSON, rec.Body.String()[:len(rec.Body.String())-1])
	}
}
