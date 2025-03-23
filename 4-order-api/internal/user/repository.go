package user

import (
	"errors"
	"go-adv/4-order-api/pkg/db"
	"go-adv/4-order-api/pkg/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepositoryDeps struct {
	Database *db.Db
}

type UserRepository struct {
	Database *db.Db
}

func NewUserRepository(database *db.Db) *UserRepository {
	return &UserRepository{Database: database}
}

func (repo *UserRepository) GetByPhoneNumber(phoneNumber string) (*models.User, error) {
	var user models.User
	result := repo.Database.DB.First(&user, "phone_number = ?", phoneNumber)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return &user, nil
}

func (repo *UserRepository) GetBySessionId(sessionId string) (*models.User, error) {
	var user models.User
	result := repo.Database.DB.First(&user, "session_id = ?", sessionId)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return &user, nil
}

func (repo *UserRepository) Create(user *models.User) (*models.User, error) {
	result := repo.Database.DB.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (repo *UserRepository) Update(user *models.User) (*models.User, error) {
	result := repo.Database.DB.Clauses(clause.Returning{}).Updates(user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}
