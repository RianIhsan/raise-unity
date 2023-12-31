package campaign

import (
	"errors"
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]Campaign, error)
	FindByUserID(userID int) ([]Campaign, error)
	FindByID(ID int) (Campaign, error)
	Save(campaign Campaign) (Campaign, error)
	Update(campaign Campaign) (Campaign, error)
	CreateImage(campaignImage CampaignImage) (CampaignImage, error)
	MarkAllImagesAsNonPrimary(campaignID int) (bool, error)
	GetTotalCampaigns() (int64, error)
	GetPaginatedCampaigns(offset, limit int) ([]Campaign, error)
	GetTotalCampaignsByUserID(userID int) (int64, error)
	GetPaginatedCampaignsByUserID(userID, offset, limit int) ([]Campaign, error)
	SearchCampaignsByName(name string) ([]Campaign, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Campaign, error) {
	var campaigns []Campaign
	err := r.db.Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (r *repository) FindByUserID(userID int) ([]Campaign, error) {
	var campaigns []Campaign
	err := r.db.Where("user_id = ?", userID).Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (r *repository) FindByID(ID int) (Campaign, error) {
	var campaign Campaign
	err := r.db.Preload("User").Preload("CampaignImages").Where("id = ?", ID).Find(&campaign).Error
	if err != nil {
		return campaign, err
	}
	return campaign, nil
}

func (r *repository) Save(campaign Campaign) (Campaign, error) {
	err := r.db.Create(&campaign).Error
	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (r *repository) Update(campaign Campaign) (Campaign, error) {
	result := r.db.Where("id = ?", campaign.ID).Updates(&campaign)
	if result.Error != nil {
		return campaign, result.Error
	}

	if result.RowsAffected == 0 {
		return campaign, errors.New("no records were updated")
	}

	return campaign, nil
}

func (r *repository) CreateImage(campaignImage CampaignImage) (CampaignImage, error) {
	err := r.db.Create(&campaignImage).Error
	if err != nil {
		return campaignImage, err
	}

	return campaignImage, nil
}

func (r *repository) MarkAllImagesAsNonPrimary(campaignID int) (bool, error) {
	err := r.db.Model(&CampaignImage{}).Where("campaign_id = ?", campaignID).Update("is_primary", false).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *repository) GetTotalCampaigns() (int64, error) {
	var total int64
	if err := r.db.Model(&Campaign{}).Count(&total).Error; err != nil {
		return 0, err
	}
	return total, nil
}

func (r *repository) GetPaginatedCampaigns(offset, limit int) ([]Campaign, error) {
	var campaigns []Campaign
	err := r.db.Offset(offset).Limit(limit).Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (r *repository) GetTotalCampaignsByUserID(userID int) (int64, error) {
	var total int64
	if err := r.db.Model(&Campaign{}).Where("user_id = ?", userID).Count(&total).Error; err != nil {
		return 0, err
	}
	return total, nil
}

func (r *repository) GetPaginatedCampaignsByUserID(userID, offset, limit int) ([]Campaign, error) {
	var campaigns []Campaign
	err := r.db.Where("user_id = ?", userID).
		Offset(offset).
		Limit(limit).
		Preload("CampaignImages", "campaign_images.is_primary = 1").
		Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}

func (r *repository) SearchCampaignsByName(name string) ([]Campaign, error) {
	var campaigns []Campaign
	err := r.db.Where("name LIKE ?", "%"+name+"%").
		Preload("CampaignImages", "campaign_images.is_primary = 1").
		Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}
