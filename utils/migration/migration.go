package migration

import (
	"fmt"
	"log"

	"github.com/RianIhsan/raise-unity/user"
	"github.com/RianIhsan/raise-unity/utils/database"
)

func GoMigrate() {
	if err := database.DB.AutoMigrate(&user.User{}, &user.OTP{}); err != nil {
		log.Fatal("Database migration failed")
	}

	fmt.Println("Successful database migration")

}
