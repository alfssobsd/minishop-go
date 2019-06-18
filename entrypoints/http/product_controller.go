package http

import (
	"github.com/alfssobsd/minishop/dataproviders/postgres"
	"github.com/alfssobsd/minishop/entrypoints/http/entities"
	"github.com/alfssobsd/minishop/usecases"
	_useCaseEntities "github.com/alfssobsd/minishop/usecases/entities"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	uuid "github.com/satori/go.uuid"
	"io"
	"io/ioutil"
	"net/http"
)

func ProductRoutes(e *echo.Echo, db *sqlx.DB) {
	//create repos and usecases
	productRepository := postgres.NewProductRepository(db)
	productUseCase := usecases.NewProductUseCase(productRepository)

	e.GET("/api/v1/products", func(c echo.Context) error {
		return searchProductsController(c, productUseCase)
	})
	e.GET("/api/v1/products/:id", func(c echo.Context) error {
		return showProductDetailInfoController(c, productUseCase)
	})

	e.POST("/api/v1/products", func(c echo.Context) error {
		return createProductController(c, productUseCase)
	})

	e.POST("/api/v1/products/excel", func(c echo.Context) error {
		return createProductsFromExcelController(c, productUseCase)
	})
}

func searchProductsController(c echo.Context, productUseCase usecases.ProductUseCase) error {
	log.Info("searchProductsController")

	productList := productUseCase.SearchProductsUseCase()
	var responseProductList []entities.HttpProductResponseEntity
	responseProductList = []entities.HttpProductResponseEntity{}

	for _, element := range productList {
		responseProductList = append(responseProductList, entities.HttpProductResponseEntity{
			ProductId:          element.ProductId,
			ProductCodeName:    element.ProductCodeName,
			ProductTitle:       element.ProductTitle,
			ProductDescription: element.ProductDescrition,
			ProductPrice:       element.ProductPrice,
		})
	}
	return c.JSON(http.StatusOK, entities.HttpProductListResponseEntity{
		Total:  len(responseProductList),
		Offset: 0,
		Items:  responseProductList,
	})
}

func showProductDetailInfoController(c echo.Context, productUseCase usecases.ProductUseCase) error {
	id := c.Param("id")
	item, err := productUseCase.ShowProductDetailInfoUseCase(uuid.FromStringOrNil(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, entities.HttpActionResponseEntity{
			Code:    http.StatusNotFound,
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, entities.HttpProductResponseEntity{
		ProductId:          item.ProductId,
		ProductCodeName:    item.ProductCodeName,
		ProductTitle:       item.ProductTitle,
		ProductDescription: item.ProductDescrition,
		ProductPrice:       item.ProductPrice,
	})
}

func createProductController(c echo.Context, productUseCase usecases.ProductUseCase) error {

	r := new(entities.HttpProductRequestEntity)
	_ = c.Bind(r)
	log.Info("createProductController ", r)

	productEntity := productUseCase.CreateProductUseCase(_useCaseEntities.ProductUseCaseEntity{
		ProductTitle:      r.ProductTitle,
		ProductCodeName:   r.ProductCodeName,
		ProductPrice:      r.ProductPrice,
		ProductDescrition: r.ProductDescription,
	})

	return c.JSON(http.StatusOK, entities.HttpProductResponseEntity{
		ProductId:          productEntity.ProductId,
		ProductCodeName:    productEntity.ProductCodeName,
		ProductTitle:       productEntity.ProductTitle,
		ProductDescription: productEntity.ProductDescrition,
		ProductPrice:       productEntity.ProductPrice,
	})
}

func createProductsFromExcelController(c echo.Context, productUseCase usecases.ProductUseCase) error {
	log.Info("createProductsFromExcelController ")

	//prepare temp file for parsing
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	defer func() {
		_ = src.Close()
	}()

	tmpfile, err := ioutil.TempFile("", "products.*.xlsx")
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if _, err = io.Copy(tmpfile, src); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	productsList := productUseCase.CreateProductFromExcelUseCase(tmpfile.Name())
	var responseList []entities.HttpProductResponseEntity
	for _, element := range productsList {
		responseList = append(responseList, entities.HttpProductResponseEntity{
			ProductId:          element.ProductId,
			ProductCodeName:    element.ProductCodeName,
			ProductTitle:       element.ProductTitle,
			ProductDescription: element.ProductDescrition,
			ProductPrice:       element.ProductPrice,
		})
	}
	return c.JSON(http.StatusOK, entities.HttpProductListResponseEntity{
		Total:  len(responseList),
		Offset: 0,
		Items:  responseList,
	})
}
