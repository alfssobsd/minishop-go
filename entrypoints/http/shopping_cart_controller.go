package http

import (
	"github.com/alfssobsd/minishop/dataproviders/postgres"
	"github.com/alfssobsd/minishop/entrypoints/http/entities"
	"github.com/alfssobsd/minishop/usecases"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"net/http"
)

//Create Routes for Cart
func CartRoutes(e *echo.Echo, db *sqlx.DB) {
	//create repos and usecases
	goodsRepo := postgres.NewGoodsRepository(db)
	orderRepo := postgres.NewOrderRepository(db)
	cartUseCase := usecases.NewShoppingCartUseCase(goodsRepo, orderRepo)

	e.GET("/api/v1/cart/:customer", func(c echo.Context) error {
		return showCustomerCartController(c, cartUseCase)
	})
	e.POST("/api/v1/cart/:customer/add", func(c echo.Context) error {
		return addGoodsToCustomerCartController(c, cartUseCase)
	})

	e.POST("/api/v1/cart/:customer/remove", func(c echo.Context) error {
		return removeGoodsFromCustomerCartController(c, cartUseCase)
	})
}

//Remove Goods from shopping cart
func removeGoodsFromCustomerCartController(c echo.Context, cartUseCase usecases.ShoppingCartUseCase) error {
	customer := c.Param("customer")
	httpRequest := &entities.HttpShoppingCartAddGoodsRequestEntity{}
	if err := c.Bind(httpRequest); err != nil {
		return c.JSON(http.StatusBadRequest, entities.HttpActionResponseEntity{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	if err := cartUseCase.RemoveGoodsFormCartUseCase(customer, httpRequest.GoodsId); err != nil {
		return c.JSON(http.StatusBadRequest, entities.HttpActionResponseEntity{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, entities.HttpActionResponseEntity{
		Code:    http.StatusOK,
		Message: "DONE",
	})
}

//add Goods to shopping cart
func addGoodsToCustomerCartController(c echo.Context, cartUseCase usecases.ShoppingCartUseCase) error {
	customer := c.Param("customer")
	httpRequest := &entities.HttpShoppingCartAddGoodsRequestEntity{}
	if err := c.Bind(httpRequest); err != nil {
		return c.JSON(http.StatusBadRequest, entities.HttpActionResponseEntity{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	if err := cartUseCase.AddGoodsToCartUseCase(customer, httpRequest.GoodsId); err != nil {
		return c.JSON(http.StatusBadRequest, entities.HttpActionResponseEntity{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, entities.HttpActionResponseEntity{
		Code:    http.StatusOK,
		Message: "DONE",
	})
}

//show customer cart
func showCustomerCartController(c echo.Context, cartUseCase usecases.ShoppingCartUseCase) error {
	customer := c.Param("customer")
	cart, _ := cartUseCase.ShowCartUseCase(customer)

	items := []entities.HttpShoppingCartItemsResponseEntity{}
	totalGoods := 0
	for _, element := range cart.GoodsItems {
		items = append(items, entities.HttpShoppingCartItemsResponseEntity{
			Goods: entities.HttpGoodsResponseEntity{
				GoodsId:          element.Goods.GoodsId,
				GoodsCodeName:    element.Goods.GoodsCodeName,
				GoodsTitle:       element.Goods.GoodsTitle,
				GoodsDescription: element.Goods.GoodsDescrition,
				GoodsPrice:       element.Goods.GoodsPrice,
			},
			Amount: element.Amount,
		})
		totalGoods += element.Amount
	}

	return c.JSON(http.StatusOK, entities.HttpShoppingCartResponseEntity{
		Customer:   customer,
		TotalGoods: totalGoods,
		TotalPrice: cart.TotalPrice,
		Items:      items,
	})
}
