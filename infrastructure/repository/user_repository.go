package repository

import (
	"digital-wallet/domain/models"

	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (repo *UserRepository) GetByID(userID uint) (*models.User, error) {
	var user models.User
	if err := repo.DB.First(&user, userID).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepository) UpdateBalance(user *models.User) error {
	return repo.DB.Save(user).Error
}

func (repo *UserRepository) Create(user *models.User) error {
	return repo.DB.Create(user).Error
}
