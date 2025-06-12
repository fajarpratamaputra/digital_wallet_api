package interfaces

import "digital-wallet/domain/models"

type UserInterface interface {
	GetByID(id uint) (*models.User, error)
	UpdateBalance(user *models.User) error
}
