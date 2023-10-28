package main

import (
	// "log"

	"github.com/RianIhsan/raise-unity/routes"
	"github.com/RianIhsan/raise-unity/utils/database"
	"github.com/RianIhsan/raise-unity/utils/migration"
	"github.com/gin-gonic/gin"
	// "github.com/joho/godotenv"
)

func main() {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Failed to fetch .env file")
	// }
	database.InitDB()
	migration.GoMigrate()

	app := gin.Default()
	routes.SetupRoute(app)
	app.Run()
}
