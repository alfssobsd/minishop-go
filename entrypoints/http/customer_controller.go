package http

import (
	"github.com/alfssobsd/minishop/dataproviders/postgres"
	"github.com/alfssobsd/minishop/entrypoints/http/entities"
	"github.com/alfssobsd/minishop/usecases"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"net/http"
)

func CustomerRoutes(e *echo.Echo, db *sqlx.DB) {
	//create repos and usecases
	custRepo := postgres.NewCustomerRepository(db)
	custUseCase := usecases.NewCustomerUseCase(custRepo)

	e.POST("/api/v1/customer/registration", func(c echo.Context) error {
		return registrationCustomerController(c, custUseCase)
	})
}

func registrationCustomerController(c echo.Context, custUseCase usecases.CustomerUseCase) error {
	r := new(entities.HttpCustomerRegistrationRequestEntity)
	err := c.Bind(r)
	log.Info("registrationCustomerController ", r)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entities.HttpActionResponseEntity{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	custEntity, err := custUseCase.CreateNewCustomer(r.CustomerUsername, r.CustomerFullName)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entities.HttpActionResponseEntity{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, entities.HttpCustomerRegistrationResponseEntity{
		CustomerId:       custEntity.CustomerId,
		CustomerUsername: custEntity.CustomerUsername,
		CustomerFullName: custEntity.CustomerFullName,
	})
}
