package http

import (
	"github.com/alfssobsd/minishop/usecases"
	"github.com/alfssobsd/minishop/usecases/entities"
	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var showGoodDetailJSON = `{"goods_id":"c26e7e02-c1de-465d-88ff-b845abdc47f1","goods_code_name":"0001","goods_title":"Плющевый медведь","goods_description":"Милый плющевый медведь","goods_price":255.5}`

type MockGoodsUC struct {
	usecases.GoodsUseCase
}

func (mcu MockGoodsUC) ShowGoodsDetailInfoUseCase(goodsId uuid.UUID) (*entities.GoodsUseCaseEntity, error) {
	return &entities.GoodsUseCaseEntity{
		GoodsId:         uuid.FromStringOrNil("c26e7e02-c1de-465d-88ff-b845abdc47f1"),
		GoodsCodeName:   "0001",
		GoodsTitle:      "Плющевый медведь",
		GoodsDescrition: "Милый плющевый медведь",
		GoodsPrice:      255.5,
	}, nil
}

func TestShowGoodsDetailInfoController(t *testing.T) {

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/v1/goods/:id")
	c.SetParamNames("id")
	c.SetParamValues("c26e7e02-c1de-465d-88ff-b845abdc47f1")

	if assert.NoError(t, showGoodsDetailInfoController(c, MockGoodsUC{})) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, showGoodDetailJSON, rec.Body.String()[:len(rec.Body.String())-1])
	}
}
