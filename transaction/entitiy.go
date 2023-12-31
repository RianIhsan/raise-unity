package transaction

import (
	"github.com/RianIhsan/raise-unity/campaign"
	"github.com/RianIhsan/raise-unity/user"
	"time"
)

type Transaction struct {
	ID         int               `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID     int               `json:"user_id" gorm:"index"`
	CampaignID int               `json:"campaign_id" gorm:"index"`
	Amount     int               `json:"amount"`
	Status     string            `json:"status"`
	Code       string            `json:"code"`
	PaymentURL string            `json:"payment_url" gorm:"type:varchar(255)"`
	User       user.User         `json:"user" gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	Campaign   campaign.Campaign `json:"campaign" gorm:"constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
	CreatedAt  time.Time         `json:"created_at"`
	UpdatedAt  time.Time         `json:"updated_at"`
}
