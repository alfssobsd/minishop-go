package http

import (
	"github.com/alfssobsd/minishop/entrypoints/http/entities"
	_goodsUsecases "github.com/alfssobsd/minishop/usecases"
	entities2 "github.com/alfssobsd/minishop/usecases/entities"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"net/http"
)

func GoodsRoutes(e *echo.Echo) {
	e.GET("/api/v1/goods", func(c echo.Context) error {
		return listGoodsController(c)
	})
	e.GET("/api/v1/goods/:id", func(c echo.Context) error {
		return showGoodsDetailInfoController(c)
	})

	e.POST("/api/v1/goods", func(c echo.Context) error {
		return createGoodsController(c)
	})
}

func listGoodsController(c echo.Context) error {
	log.Info("listGoodsController")
	goodsList := _goodsUsecases.SearchGoodsUseCase()
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
	item := _goodsUsecases.ShowGoodsDetailInfoUseCase(id)
	return c.JSON(http.StatusOK, entities.HttpGoodsResponseEntity{
		GoodsId:          item.GoodsId,
		GoodsCodeName:    item.GoodsCodeName,
		GoodsTitle:       item.GoodsTitle,
		GoodsDescription: item.GoodsDescrition,
		GoodsPrice:       item.GoodsPrice,
	})
}

func createGoodsController(c echo.Context) error {

	r := new(entities.HttpGoodsRequestEntity)
	_ = c.Bind(r)
	log.Info("createGoodsController ", r)

	item := _goodsUsecases.CreateGoodsUseCase(entities2.GoodsUseCaseEntity{
		GoodsTitle:      r.GoodsTitle,
		GoodsCodeName:   r.GoodsCodeName,
		GoodsPrice:      r.GoodsPrice,
		GoodsDescrition: r.GoodsDescription,
	})

	return c.JSON(http.StatusOK, entities.HttpGoodsResponseEntity{
		GoodsId:          item.GoodsId,
		GoodsCodeName:    item.GoodsCodeName,
		GoodsTitle:       item.GoodsTitle,
		GoodsDescription: item.GoodsDescrition,
		GoodsPrice:       item.GoodsPrice,
	})
}
