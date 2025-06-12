package usecases

import (
	"digital-wallet/domain/interfaces"
	"digital-wallet/domain/models"
	"fmt"
)

type WithdrawFunds struct {
	UserInterface interfaces.UserInterface
}

func (uc *WithdrawFunds) Execute(userID uint, amount float64) (*models.User, error) {
	user, err := uc.UserInterface.GetByID(userID)
	if err != nil {
		return nil, err
	}

	if user.Balance < amount {
		return nil, fmt.Errorf("insufficient balance")
	}

	user.Balance -= amount
	err = uc.UserInterface.UpdateBalance(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
