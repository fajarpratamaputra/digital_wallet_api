package api

import (
	"digital-wallet/domain/models"
	"digital-wallet/infrastructure/repository"
	"digital-wallet/internal"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func AddUser(c *gin.Context) {
	var req struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
		return
	}

	user := models.User{
		Name:     req.Name,
		Password: string(hashedPassword),
		Balance:  0, // Default balance
	}

	userRepo := repository.UserRepository{DB: internal.DB}
	if err := userRepo.Create(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully", "user": user})
}
