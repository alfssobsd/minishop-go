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
	productRepo := postgres.NewProductRepository(db)
	orderRepo := postgres.NewOrderRepository(db)
	custRepo := postgres.NewCustomerRepository(db)
	cartUseCase := usecases.NewShoppingCartUseCase(productRepo, orderRepo, custRepo)

	e.GET("/api/v1/cart/:customer", func(c echo.Context) error {
		return showCustomerCartController(c, cartUseCase)
	})
	e.POST("/api/v1/cart/:customer/add", func(c echo.Context) error {
		return addProductToCustomerCartController(c, cartUseCase)
	})

	e.POST("/api/v1/cart/:customer/remove", func(c echo.Context) error {
		return removeProductFromCustomerCartController(c, cartUseCase)
	})
}

//Remove Product from shopping cart
func removeProductFromCustomerCartController(c echo.Context, cartUseCase usecases.ShoppingCartUseCase) error {
	customer := c.Param("customer")
	httpRequest := &entities.HttpShoppingCartAddProductRequestEntity{}
	if err := c.Bind(httpRequest); err != nil {
		return c.JSON(http.StatusBadRequest, entities.HttpActionResponseEntity{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	if err := cartUseCase.RemoveProductFormCartUseCase(customer, httpRequest.ProductId); err != nil {
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

//add Product to shopping cart
func addProductToCustomerCartController(c echo.Context, cartUseCase usecases.ShoppingCartUseCase) error {
	customer := c.Param("customer")
	httpRequest := &entities.HttpShoppingCartAddProductRequestEntity{}
	if err := c.Bind(httpRequest); err != nil {
		return c.JSON(http.StatusBadRequest, entities.HttpActionResponseEntity{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	if err := cartUseCase.AddProductToCartUseCase(customer, httpRequest.ProductId); err != nil {
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
	totalProducts := 0
	for _, element := range cart.ProductItems {
		items = append(items, entities.HttpShoppingCartItemsResponseEntity{
			Product: entities.HttpProductResponseEntity{
				ProductId:          element.Product.ProductId,
				ProductCodeName:    element.Product.ProductCodeName,
				ProductTitle:       element.Product.ProductTitle,
				ProductDescription: element.Product.ProductDescrition,
				ProductPrice:       element.Product.ProductPrice,
			},
			Amount: element.Amount,
		})
		totalProducts += element.Amount
	}

	return c.JSON(http.StatusOK, entities.HttpShoppingCartResponseEntity{
		CustomerId:    cart.CustomerId,
		TotalProducts: totalProducts,
		TotalPrice:    cart.TotalPrice,
		Items:         items,
	})
}
