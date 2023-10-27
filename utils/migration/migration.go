package migration

import (
	"fmt"
	entity2 "github.com/RianIhsan/raise-unity/campaign"
	"github.com/RianIhsan/raise-unity/transaction"
	"github.com/RianIhsan/raise-unity/user"
	"log"

	"github.com/RianIhsan/raise-unity/utils/database"
)

func GoMigrate() {
	if err := database.DB.AutoMigrate(&user.User{}, &user.OTP{}, &entity2.Campaign{}, &entity2.CampaignImage{}, &transaction.Transaction{}); err != nil {
		log.Fatal("Database migration failed")
	}

	fmt.Println("Successful database migration")

}
