package postgres

import (
	"github.com/alfssobsd/minishop/config"
	"github.com/alfssobsd/minishop/dataproviders/postgres/entities"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCustomerRepository_Create(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
	db := config.MakePostgresConnection()

	customerId := uuid.NewV4()
	username := "sergei"
	fullName := "Sergei Kravchuk"
	repo := NewCustomerRepository(db)
	_ = repo.CreateCustomer(customerId, username, fullName)

	expectedCustomer := entities.CustomerEntity{
		CustomerId:       customerId,
		CustomerUsername: username,
		CustomerFullName: fullName}

	customer, err := repo.FindById(customerId)
	assert.Equal(t, err, nil)
	assert.Equal(t, &expectedCustomer, customer)
}
