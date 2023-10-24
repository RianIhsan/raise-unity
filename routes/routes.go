package routes

import (
	"github.com/RianIhsan/raise-unity/auth"
	"github.com/RianIhsan/raise-unity/handler"
	"github.com/RianIhsan/raise-unity/middleware"
	"github.com/RianIhsan/raise-unity/user"
	"github.com/RianIhsan/raise-unity/utils/database"
	"github.com/gin-gonic/gin"
)

func SetupRoute(router *gin.Engine) {
	authService := auth.NewService()
	userRepository := user.NewRepository(database.DB)

	userService := user.NewService(userRepository)

	userHandler := handler.NewUserHandler(userService, authService)
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/verify", userHandler.VerifyEmail)
	api.POST("/resend-otp", userHandler.ResendOTP)
	api.PATCH("/avatar", middleware.AuthMiddleware(authService, userService), userHandler.UploadAvatar)
}
