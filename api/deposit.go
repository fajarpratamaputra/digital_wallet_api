package api

import (
	"digital-wallet/infrastructure/repository"
	"digital-wallet/internal"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Deposit(c *gin.Context) {
	var req struct {
		Amount float64 `json:"amount"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	userID, _ := c.Get("user_id")

	userRepo := repository.UserRepository{DB: internal.DB}
	user, err := userRepo.GetByID(userID.(uint))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	user.Balance += req.Amount

	if err := userRepo.UpdateBalance(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to deposit funds"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Deposit successful", "balance": user.Balance})
}
