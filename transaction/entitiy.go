package transaction

import (
	"github.com/RianIhsan/raise-unity/user"
	"time"
)

type Transaction struct {
	ID         int       `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID     int       `json:"user_id" gorm:"index"`
	CampaignID int       `json:"campaign_id" gorm:"index"`
	Amount     int       `json:"amount"`
	Status     string    `json:"status"`
	Code       string    `json:"code"`
	User       user.User `json:"user"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
