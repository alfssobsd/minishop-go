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
