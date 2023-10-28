package campaign

import "strings"

type CampaignFormatter struct {
	ID               int    `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID           int    `json:"user_id" gorm:"index"`
	Name             string `json:"name" gorm:"type:varchar(255)"`
	ShortDescription string `json:"short_description" gorm:"type:varchar(255)"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
	CampaignImages   string `json:"campaign_images,omitempty" gorm:"foreignKey:CampaignID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func FormatCampaign(campaign Campaign) CampaignFormatter {
	campaignFormatter := CampaignFormatter{}
	campaignFormatter.ID = campaign.ID
	campaignFormatter.UserID = campaign.UserID
	campaignFormatter.Name = campaign.Name
	campaignFormatter.ShortDescription = campaign.ShortDescription
	campaignFormatter.GoalAmount = campaign.GoalAmount
	campaignFormatter.CurrentAmount = campaign.CurrentAmount

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

type CampaignDetailsFormatter struct {
	ID               int                       `json:"id"`
	Name             string                    `json:"name"`
	ShortDescription string                    `json:"short_description"`
	Description      string                    `json:"description"`
	CampaignImages   string                    `json:"campaign_images" gorm:"foreignKey:CampaignID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	GoalAmount       int                       `json:"goal_amount"`
	CurrentAmount    int                       `json:"current_amount"`
	BackerCount      int                       `json:"backer_count"`
	UserID           int                       `json:"user_id"`
	Perks            []string                  `json:"perks"`
	User             CampaignUserFormatter     `json:"user"`
	Images           []CampaigsImagesFormatter `json:"images"`
}

type CampaignUserFormatter struct {
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

type CampaigsImagesFormatter struct {
	CampaignImages string `json:"campaign_image"`
	IsPrimary      bool   `json:"is_primary"`
}

func FormatCampaignDetails(campaign Campaign) CampaignDetailsFormatter {
	campaignDetailsFormatter := CampaignDetailsFormatter{}
	campaignDetailsFormatter.ID = campaign.ID
	campaignDetailsFormatter.Name = campaign.Name
	campaignDetailsFormatter.ShortDescription = campaign.ShortDescription
	campaignDetailsFormatter.Description = campaign.Description
	campaignDetailsFormatter.GoalAmount = campaign.GoalAmount
	campaignDetailsFormatter.CurrentAmount = campaign.CurrentAmount
	campaignDetailsFormatter.BackerCount = campaign.BackerCount
	campaignDetailsFormatter.UserID = campaign.UserID
	campaignDetailsFormatter.CampaignImages = ""

	if len(campaign.CampaignImages) > 0 {
		var campaignImage string

		for _, image := range campaign.CampaignImages {
			if image.IsPrimary > 0 {
				campaignImage = image.FileName
				break
			}
		}

		campaignDetailsFormatter.CampaignImages = campaignImage
	}
	var perks []string
	for _, perk := range strings.Split(campaign.Perks, ",") {
		perks = append(perks, strings.TrimSpace(perk))
	}
	campaignDetailsFormatter.Perks = perks

	user := campaign.User
	campaignUserFormatter := CampaignUserFormatter{}
	campaignUserFormatter.Name = user.Name
	campaignUserFormatter.Avatar = user.Avatar

	campaignDetailsFormatter.User = campaignUserFormatter

	images := []CampaigsImagesFormatter{}
	for _, img := range campaign.CampaignImages {
		campaignImagesFormatter := CampaigsImagesFormatter{}
		campaignImagesFormatter.CampaignImages = img.FileName
		isPrimary := true
		if img.IsPrimary == 0 {
			isPrimary = false
		}
		campaignImagesFormatter.IsPrimary = isPrimary
		images = append(images, campaignImagesFormatter)
	}
	campaignDetailsFormatter.Images = images
	return campaignDetailsFormatter
}
