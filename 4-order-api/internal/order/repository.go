package order

import (
	"go-adv/4-order-api/pkg/db"
	"go-adv/4-order-api/pkg/models"
)

type OrderRepository struct {
	Database *db.Db
}

func NewOrderRepository(database *db.Db) *OrderRepository {
	return &OrderRepository{
		Database: database,
	}
}

func (repo *OrderRepository) Create(order *models.Order) (*models.Order, error) {
	result := repo.Database.DB.Create(order)
	if result.Error != nil {
		return nil, result.Error
	}

	return order, nil
}

func (repo *OrderRepository) GetById(id string) (*models.Order, error) {
	var order models.Order
	result := repo.Database.DB.Preload("Products").First(&order, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &order, nil
}

func (repo *OrderRepository) GetAllByUserId(id string) (*[]models.Order, error) {
	var orders []models.Order
	result := repo.Database.DB.Preload("Products").Where("user_id = ?", id).Find(&orders)

	if result.Error != nil {
		return nil, result.Error
	}
	return &orders, nil
}
