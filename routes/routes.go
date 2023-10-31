package routes

import (
	"github.com/RianIhsan/raise-unity/admin"
	"github.com/RianIhsan/raise-unity/auth"
	"github.com/RianIhsan/raise-unity/campaign"
	"github.com/RianIhsan/raise-unity/handler"
	"github.com/RianIhsan/raise-unity/middleware"
	"github.com/RianIhsan/raise-unity/payment"
	"github.com/RianIhsan/raise-unity/transaction"
	"github.com/RianIhsan/raise-unity/user"
	"github.com/RianIhsan/raise-unity/utils/database"
	"github.com/gin-gonic/gin"
)

func SetupRoute(router *gin.Engine) {

	adminRepo := admin.NewRepository(database.DB)
	userRepository := user.NewRepository(database.DB)
	campRepository := campaign.NewRepository(database.DB)
	transactionRepository := transaction.NewRepository(database.DB)

	authService := auth.NewService()
	adminService := admin.NewService(adminRepo)
	userService := user.NewService(userRepository)
	campService := campaign.NewService(campRepository)
	paymentService := payment.NewService()
	transactionService := transaction.NewService(transactionRepository, campRepository, paymentService)

	adminHandler := handler.NewAdminHandler(adminService, authService)
	userHandler := handler.NewUserHandler(userService, authService)
	campHandler := handler.NewCampaignHandler(campService)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	api := router.Group("/api/v1")

	api.GET("/healthcheck", handler.HealthCheck)
	api.POST("/register", userHandler.RegisterUser)
	api.POST("/login", userHandler.Login)
	api.POST("/verify", userHandler.VerifyEmail)
	api.POST("/resend-otp", userHandler.ResendOTP)
	api.PATCH("/avatar", middleware.AuthMiddleware(authService, userService), userHandler.UploadAvatar)

	api.GET("/campaigns", campHandler.GetCampaigns)
	api.GET("/campaign/:id", campHandler.GetCampaign)
	api.POST("/campaign", middleware.AuthMiddleware(authService, userService), campHandler.CreateCampaign)
	api.PUT("/campaign/:id", middleware.AuthMiddleware(authService, userService), campHandler.UpdateCampaign)
	api.POST("/campaign-images", middleware.AuthMiddleware(authService, userService), campHandler.UploadImage)

	api.GET("/campaigns/:id/transactions", middleware.AuthMiddleware(authService, userService), transactionHandler.GetCampaignTransactions)
	api.GET("/transactions", middleware.AuthMiddleware(authService, userService), transactionHandler.GetUserTransactions)
	api.POST("/transactions", middleware.AuthMiddleware(authService, userService), transactionHandler.CreateTransaction)
	api.POST("/transactions/notification", transactionHandler.GetNotification)

	api.GET("/users", middleware.AuthMiddleware(authService, userService), adminHandler.GetAllUsers)
	api.GET("/users/transactions", middleware.AuthMiddleware(authService, userService), adminHandler.GetAllUsersTransactions)
	api.DELETE("/users/:id", middleware.AuthMiddleware(authService, userService), adminHandler.DeleteUser)
	api.DELETE("/campaign/:id", middleware.AuthMiddleware(authService, userService), adminHandler.DeleteCampaign)
}
