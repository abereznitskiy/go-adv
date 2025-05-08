package product

import (
	"go-adv/4-order-api/pkg/db"
	"go-adv/4-order-api/pkg/models"

	"gorm.io/gorm/clause"
)

type ProductRepositoryDeps struct {
	Database *db.Db
}

type ProductRepository struct {
	Database *db.Db
}

func NewProductRepository(database *db.Db) *ProductRepository {
	return &ProductRepository{Database: database}
}

func (repo *ProductRepository) GetById(id uint) (*models.Product, error) {
	var product models.Product
	result := repo.Database.DB.First(&product, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &product, nil
}

func (repo *ProductRepository) GetAll() (*[]models.Product, error) {
	var products []models.Product
	result := repo.Database.DB.Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}

	return &products, nil
}

func (repo *ProductRepository) Create(product *models.Product) (*models.Product, error) {
	result := repo.Database.DB.Create(product)
	if result.Error != nil {
		return nil, result.Error
	}

	return product, nil
}

func (repo *ProductRepository) Update(product *models.Product) (*models.Product, error) {
	result := repo.Database.DB.Clauses(clause.Returning{}).Updates(product)
	if result.Error != nil {
		return nil, result.Error
	}

	return product, nil
}

func (repo *ProductRepository) Delete(id uint) error {
	result := repo.Database.DB.Delete(&models.Product{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
