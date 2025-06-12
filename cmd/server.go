package main

import (
	"digital-wallet/api"
	"digital-wallet/internal"
	"digital-wallet/internal/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	internal.InitDB()

	r := gin.Default()

	r.POST("/login", api.Login)
	r.POST("/user", api.AddUser) // Add User route (open for creation)

	r.Use(middleware.AuthMiddleware)
	r.POST("/withdraw", api.Withdraw)
	r.POST("/deposit", api.Deposit)
	r.GET("/balance", api.BalanceInquiry)

	r.Run(":8080")
}
