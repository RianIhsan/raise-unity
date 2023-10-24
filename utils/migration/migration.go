package migration

import (
	"fmt"
	"github.com/RianIhsan/raise-unity/campaign"
	"log"

	"github.com/RianIhsan/raise-unity/user"
	"github.com/RianIhsan/raise-unity/utils/database"
)

func GoMigrate() {
	if err := database.DB.AutoMigrate(&user.User{}, &user.OTP{}, &campaign.Campaign{}, &campaign.CampaignImage{}); err != nil {
		log.Fatal("Database migration failed")
	}

	fmt.Println("Successful database migration")

}
