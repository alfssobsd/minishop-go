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

var showProductDetailJSON = `{"product_id":"c26e7e02-c1de-465d-88ff-b845abdc47f1","product_code_name":"0001","product_title":"Плющевый медведь","product_description":"Милый плющевый медведь","product_price":255.5}`

type MockGoodsUC struct {
	usecases.ProductUseCase
}

func (mcu MockGoodsUC) ShowProductDetailInfoUseCase(goodsId uuid.UUID) (*entities.ProductUseCaseEntity, error) {
	return &entities.ProductUseCaseEntity{
		ProductId:         uuid.FromStringOrNil("c26e7e02-c1de-465d-88ff-b845abdc47f1"),
		ProductCodeName:   "0001",
		ProductTitle:      "Плющевый медведь",
		ProductDescrition: "Милый плющевый медведь",
		ProductPrice:      255.5,
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

	if assert.NoError(t, showProductDetailInfoController(c, MockGoodsUC{})) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, showProductDetailJSON, rec.Body.String()[:len(rec.Body.String())-1])
	}
}
