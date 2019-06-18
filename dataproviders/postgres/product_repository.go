package postgres

import (
	"github.com/alfssobsd/minishop/dataproviders/postgres/entities"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
)

type ProductRepository interface {
	FindAll() []*entities.ProductEntity
	CreateOne(entities.ProductEntity)
	FindById(productId uuid.UUID) *entities.ProductEntity
	FindByCodeName(codeName string) *entities.ProductEntity
}

type productRepository struct {
	db *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) *productRepository {
	return &productRepository{db: db}
}

func (r *productRepository) CreateOne(productEntity entities.ProductEntity) {
	log.Info("CreateOne ", productEntity)
	result, err := r.db.NamedExec("INSERT INTO products (uuid, title, code_name, description, price) VALUES (:uuid, :title, :code_name, :description, :price)", &productEntity)
	if err != nil {
		log.Error(err)
	}
	log.Info(result)
}

func (r *productRepository) FindById(productId uuid.UUID) *entities.ProductEntity {
	log.Info("FindById ", productId)

	product := entities.ProductEntity{}
	err := r.db.Get(&product, "SELECT * FROM products WHERE uuid=$1", productId)
	if err != nil {
		log.Error(err)
		return nil
	}
	return &product
}

func (r *productRepository) FindByCodeName(codeName string) *entities.ProductEntity {
	log.Info("FindByCodeName ", codeName)
	product := entities.ProductEntity{}
	err := r.db.Get(&product, "SELECT * FROM products WHERE code_name=$1", codeName)
	if err != nil {
		log.Error(err)
		return nil
	}
	return &product
}

func (r *productRepository) FindAll() []*entities.ProductEntity {
	log.Info("FindAll ")
	var productEntities []*entities.ProductEntity
	err := r.db.Select(&productEntities, "SELECT * FROM products ORDER BY code_name ASC")
	if err != nil {
		log.Error(err)
	}

	return productEntities
}
