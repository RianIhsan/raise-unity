package campaign

type CampaignFormatter struct {
	ID               int    `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID           int    `json:"user_id" gorm:"index"`
	Name             string `json:"name" gorm:"type:varchar(255)"`
	ShortDescription string `json:"short_description" gorm:"type:varchar(255)"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
	CampaignImages   string `json:"campaign_images" gorm:"foreignKey:CampaignID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func FormatCampaign(campaign Campaign) CampaignFormatter {
	campaignFormatter := CampaignFormatter{}
	campaignFormatter.ID = campaign.ID
	campaignFormatter.UserID = campaign.UserID
	campaignFormatter.Name = campaign.Name
	campaignFormatter.ShortDescription = campaign.ShortDescription
	campaignFormatter.GoalAmount = campaign.GoalAmount
	campaignFormatter.CurrentAmount = campaign.CurrentAmount
	campaignFormatter.CampaignImages = ""

	if len(campaign.CampaignImages) > 0 {
		campaignFormatter.CampaignImages = campaign.CampaignImages[0].FileName
	}

	return campaignFormatter
}

func FormatCampaigns(campaigns []Campaign) []CampaignFormatter {
	var campaignsFormatter []CampaignFormatter

	for _, campaign := range campaigns {
		campaignFormatter := FormatCampaign(campaign)
		campaignsFormatter = append(campaignsFormatter, campaignFormatter)
	}

	return campaignsFormatter
}
