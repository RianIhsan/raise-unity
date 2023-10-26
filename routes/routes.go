package routes

import (
	"github.com/RianIhsan/raise-unity/auth"
	"github.com/RianIhsan/raise-unity/campaign"
	"github.com/RianIhsan/raise-unity/handler"
	"github.com/RianIhsan/raise-unity/middleware"
	"github.com/RianIhsan/raise-unity/user"
	"github.com/RianIhsan/raise-unity/utils/database"
	"github.com/gin-gonic/gin"
)

func SetupRoute(router *gin.Engine) {
	authService := auth.NewService()
	userRepository := user.NewRepository(database.DB)
	campRepository := campaign.NewRepository(database.DB)

	userService := user.NewService(userRepository)
	campService := campaign.NewService(campRepository)

	userHandler := handler.NewUserHandler(userService, authService)
	campHandler := handler.NewCampaignHandler(campService)
	api := router.Group("/api/v1")

	api.POST("/register", userHandler.RegisterUser)
	api.POST("/login", userHandler.Login)
	api.POST("/verify", userHandler.VerifyEmail)
	api.POST("/resend-otp", userHandler.ResendOTP)
	api.PATCH("/avatar", middleware.AuthMiddleware(authService, userService), userHandler.UploadAvatar)

	api.GET("/campaigns", campHandler.GetCampaigns)
	api.GET("/campaign/:id", campHandler.GetCampaign)
	api.POST("/campaign", middleware.AuthMiddleware(authService, userService), campHandler.CreateCampaign)
	api.PUT("/campaign/:id", middleware.AuthMiddleware(authService, userService), campHandler.UpdateCampaign)
}
