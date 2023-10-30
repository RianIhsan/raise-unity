package campaign

import (
	"github.com/RianIhsan/raise-unity/user"
	"time"
)

type Campaign struct {
	ID               int             `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID           int             `json:"user_id" gorm:"index"`
	Name             string          `json:"name" gorm:"type:varchar(255)"`
	ShortDescription string          `json:"short_description" gorm:"type:varchar(255)"`
	Description      string          `json:"description" gorm:"type:text"`
	Perks            string          `json:"perks" gorm:"type:text"`
	BackerCount      int             `json:"backer_count"`
	GoalAmount       int             `json:"goal_amount"`
	CurrentAmount    int             `json:"current_amount"`
	CreatedAt        time.Time       `json:"created_at"`
	UpdatedAt        time.Time       `json:"updated_at"`
	CampaignImages   []CampaignImage `json:"campaign_images" gorm:"foreignKey:CampaignID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	User             user.User       `json:"user"`
}

type CampaignImage struct {
	ID         int       `json:"id" gorm:"primaryKey;autoIncrement"`
	CampaignID int       `json:"campaign_id" gorm:"index"`
	FileName   string    `json:"file_name" form:"file_name" gorm:"type:varchar(255)"`
	IsPrimary  int       `json:"is_primary" gorm:"type:tinyint(4)"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
