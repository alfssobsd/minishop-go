package http

import (
	"github.com/alfssobsd/minishop/entrypoints/http/entities"
	_goodsUsecases "github.com/alfssobsd/minishop/usecases"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"gopkg.in/mgo.v2"
	"net/http"
)

func GoodsRoutes(e *echo.Echo, db *mgo.Session) {
	e.GET("/goods", func(c echo.Context) error {
		return listGoodsController(c, db)
	})
	e.GET("/goods/:id", func(c echo.Context) error {
		return showGoodsDetailInfoController(c)
	})
}

func listGoodsController(c echo.Context, db *mgo.Session) error {
	log.Info("listGoodsController")
	goodsList := _goodsUsecases.SearchGoodsUseCase(db)
	responseGoodsList := []entities.HttpGoodsResponseEntity{}
	for _, element := range goodsList {
		responseGoodsList = append(responseGoodsList, entities.HttpGoodsResponseEntity{
			GoodsId:          element.GoodsId,
			GoodsCodeName:    element.GoodsCodeName,
			GoodsTitle:       element.GoodsTitle,
			GoodsDescription: element.GoodsDescrition,
			GoodsPrice:       element.GoodsPrice,
		})
	}
	return c.JSON(http.StatusOK, entities.HttpGoodsListResponseEntity{
		Total:  len(responseGoodsList),
		Offset: 0,
		Items:  responseGoodsList,
	})
}

func showGoodsDetailInfoController(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusOK, _goodsUsecases.ShowGoodsDetailInfoUseCase(id))
}
