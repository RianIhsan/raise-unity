package main

import (
	"log"

	"github.com/RianIhsan/raise-unity/auth"
	"github.com/RianIhsan/raise-unity/handler"
	"github.com/RianIhsan/raise-unity/user"
	"github.com/RianIhsan/raise-unity/utils/database"
	"github.com/RianIhsan/raise-unity/utils/migration"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to fetch .env file")
	}
	database.InitDB()
	migration.GoMigrate()

	userRepository := user.NewRepository(database.DB)
	userService := user.NewService(userRepository)
	authService := auth.NewService()

	userHandler := handler.NewUserHandler(userService, authService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/verify", userHandler.VerifyEmail)
	api.POST("/resend-otp", userHandler.ResendOTP)

	router.Run()
}
