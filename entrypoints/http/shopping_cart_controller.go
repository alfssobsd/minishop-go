package http

import (
	"github.com/alfssobsd/minishop/dataproviders/postgres"
	"github.com/alfssobsd/minishop/entrypoints/http/entities"
	"github.com/alfssobsd/minishop/usecases"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"net/http"
)

func CartRoutes(e *echo.Echo, db *sqlx.DB) {
	//create repos and usecases
	goodsRepo := postgres.NewGoodsRepository(db)
	orderRepo := postgres.NewOrderRepository(db)
	cartUseCase := usecases.NewShoppingCartUseCase(goodsRepo, orderRepo)

	e.GET("/api/v1/cart/:customer", func(c echo.Context) error {
		return showCustomerCart(c, cartUseCase)
	})
}

func showCustomerCart(c echo.Context, cartUseCase usecases.ShoppingCartUseCase) error {
	customer := c.Param("customer")
	cart := cartUseCase.ShowCartUseCase(customer)

	httpGoodsItems := []entities.HttpGoodsResponseEntity{}
	for _, element := range cart.GoodsItems {
		httpGoodsItems = append(httpGoodsItems, entities.HttpGoodsResponseEntity{
			GoodsId:          element.GoodsId,
			GoodsCodeName:    element.GoodsCodeName,
			GoodsTitle:       element.GoodsTitle,
			GoodsDescription: element.GoodsDescrition,
			GoodsPrice:       element.GoodsPrice,
		})
	}

	return c.JSON(http.StatusOK, entities.HttpShoppingCartEntity{
		Customer: customer,
		//TODO: incorrect value, need to consider amount
		TotalGoods: len(cart.GoodsItems),
		TotalPrice: cart.TotalPrice,
		//TODO: incorrect value, need show amount
		GoodsItems: httpGoodsItems,
	})
}
