package postgres

import (
	"github.com/alfssobsd/minishop/dataproviders/postgres/entities"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
	uuid "github.com/satori/go.uuid"
)

type CustomerRepository interface {
	FindByUsername(username string) (*entities.CustomerEntity, error)
	FindById(customerId uuid.UUID) (*entities.CustomerEntity, error)
	CreateCustomer(customerId uuid.UUID, username string, fullName string) error
}

type customerRepository struct {
	db *sqlx.DB
}

func NewCustomerRepository(db *sqlx.DB) *customerRepository {
	return &customerRepository{db: db}
}

func (r *customerRepository) FindByUsername(username string) (*entities.CustomerEntity, error) {
	log.Info("FindByUsername ", username)
	customerEntity := entities.CustomerEntity{}
	err := r.db.Get(&customerEntity, "SELECT * FROM customers WHERE username=$1", username)
	if err != nil {
		return nil, err
	}
	return &customerEntity, nil
}

func (r *customerRepository) FindById(customerId uuid.UUID) (*entities.CustomerEntity, error) {
	log.Info("FindByUuid ", customerId)
	customerEntity := entities.CustomerEntity{}
	err := r.db.Get(&customerEntity, "SELECT * FROM customers WHERE uuid=$1", customerId)
	if err != nil {
		return nil, err
	}
	return &customerEntity, nil
}

func (r *customerRepository) CreateCustomer(customerId uuid.UUID, username string, fullName string) error {
	log.Infof("Create uuid = %s, username = %s, fullName = %s", customerId, username, fullName)
	_, err := r.db.Exec("INSERT INTO customers (uuid, username, full_name) VALUES ($1, $2, $3)", customerId, username, fullName)
	return err
}
