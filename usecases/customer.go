package usecases

import (
	"errors"
	_repo "github.com/alfssobsd/minishop/dataproviders/postgres"
	"github.com/alfssobsd/minishop/usecases/entities"
	"github.com/labstack/gommon/log"
	uuid "github.com/satori/go.uuid"
)

type CustomerUseCase interface {
	CreateNewCustomer(username string, fullName string) (*entities.CustomerUseCaseEntity, error)
}

type customerUseCase struct {
	custRepo _repo.CustomerRepository
}

func NewCustomerUseCase(custRepo _repo.CustomerRepository) *customerUseCase {
	return &customerUseCase{custRepo}
}

func (u *customerUseCase) CreateNewCustomer(username string, fullName string) (*entities.CustomerUseCaseEntity, error) {
	customerId := uuid.NewV4()
	log.Info("CreateNewCustomer id = ", customerId.String())

	customer, _ := u.custRepo.FindByUsername(username)
	if customer != nil {
		return nil, errors.New("customer already created")
	}

	err := u.custRepo.CreateCustomer(customerId, username, fullName)
	if err != nil {
		return nil, err
	}

	customer, err = u.custRepo.FindById(customerId)
	if err != nil {
		return nil, err
	}

	return &entities.CustomerUseCaseEntity{
		CustomerId:       customerId,
		CustomerFullName: customer.CustomerFullName,
		CustomerUsername: customer.CustomerUsername,
	}, nil
}
