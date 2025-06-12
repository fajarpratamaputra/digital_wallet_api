package api

import (
	"digital-wallet/infrastructure/repository"
	"digital-wallet/internal"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BalanceInquiry(c *gin.Context) {
	userID, _ := c.Get("user_id")

	userRepo := repository.UserRepository{DB: internal.DB}
	user, err := userRepo.GetByID(userID.(uint))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"balance": user.Balance})
}
