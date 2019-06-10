package http

import (
	"github.com/alfssobsd/minishop/dataproviders/mongodb"
	"github.com/alfssobsd/minishop/entrypoints/http/entities"
	_goodsUC "github.com/alfssobsd/minishop/usecases"
	_useCaseEntities "github.com/alfssobsd/minishop/usecases/entities"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"gopkg.in/mgo.v2"
	"net/http"
)

func GoodsRoutes(e *echo.Echo, mgoSession *mgo.Session) {
	//create repos and usecases
	goodsRepository := mongodb.NewGoodsRepository(mgoSession)
	goodsUseCase := _goodsUC.NewGoodsUseCase(goodsRepository)

	e.GET("/api/v1/goods", func(c echo.Context) error {
		return listGoodsController(c, goodsUseCase)
	})
	e.GET("/api/v1/goods/:id", func(c echo.Context) error {
		return showGoodsDetailInfoController(c, goodsUseCase)
	})

	e.POST("/api/v1/goods", func(c echo.Context) error {
		return createGoodsController(c, goodsUseCase)
	})
}

func listGoodsController(c echo.Context, goodsUseCase *_goodsUC.GoodsUseCase) error {
	log.Info("listGoodsController")
	goodsList := goodsUseCase.SearchGoodsUseCase()
	var responseGoodsList []entities.HttpGoodsResponseEntity
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

func showGoodsDetailInfoController(c echo.Context, goodsUseCase *_goodsUC.GoodsUseCase) error {
	id := c.Param("id")
	item := goodsUseCase.ShowGoodsDetailInfoUseCase(id)
	return c.JSON(http.StatusOK, entities.HttpGoodsResponseEntity{
		GoodsId:          item.GoodsId,
		GoodsCodeName:    item.GoodsCodeName,
		GoodsTitle:       item.GoodsTitle,
		GoodsDescription: item.GoodsDescrition,
		GoodsPrice:       item.GoodsPrice,
	})
}

func createGoodsController(c echo.Context, goodsUseCase *_goodsUC.GoodsUseCase) error {

	r := new(entities.HttpGoodsRequestEntity)
	_ = c.Bind(r)
	log.Info("createGoodsController ", r)

	goodsEntity := goodsUseCase.CreateGoodsUseCase(_useCaseEntities.GoodsUseCaseEntity{
		GoodsTitle:      r.GoodsTitle,
		GoodsCodeName:   r.GoodsCodeName,
		GoodsPrice:      r.GoodsPrice,
		GoodsDescrition: r.GoodsDescription,
	})

	return c.JSON(http.StatusOK, entities.HttpGoodsResponseEntity{
		GoodsId:          goodsEntity.GoodsId,
		GoodsCodeName:    goodsEntity.GoodsCodeName,
		GoodsTitle:       goodsEntity.GoodsTitle,
		GoodsDescription: goodsEntity.GoodsDescrition,
		GoodsPrice:       goodsEntity.GoodsPrice,
	})
}
