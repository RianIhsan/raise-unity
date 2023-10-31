package campaign

import (
	"errors"
	"math"
)

type Service interface {
	FindCampaigns(userID int) ([]Campaign, error)
	FindCampaignByID(input GetCampaignDetailInput) (Campaign, error)
	CreateCampaign(input CreateCampaignInput) (Campaign, error)
	UpdateCampaign(inputID GetCampaignDetailInput, data CreateCampaignInput) (Campaign, error)
	SaveCampaignImage(input CreateCampaignImageInput, file string) (CampaignImage, error)
	GetPaginatedCampaigns(page, pageSize int) ([]Campaign, int, int, int, int, error)
	GetPaginatedCampaignsByUserID(userID, page, pageSize int) ([]Campaign, int, int, int, int, error)
	SearchCampaignsByName(name string) ([]Campaign, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindCampaigns(userID int) ([]Campaign, error) {
	if userID != 0 {
		campaigns, err := s.repository.FindByUserID(userID)
		if err != nil {
			return campaigns, err
		}
		return campaigns, nil
	}
	campaigns, err := s.repository.FindAll()
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}

func (s *service) FindCampaignByID(input GetCampaignDetailInput) (Campaign, error) {
	campaign, err := s.repository.FindByID(input.ID)
	if err != nil {
		return campaign, err
	}
	return campaign, nil
}

func (s *service) CreateCampaign(input CreateCampaignInput) (Campaign, error) {
	campaign := Campaign{}
	campaign.Name = input.Name
	campaign.ShortDescription = input.ShortDescription
	campaign.Description = input.Description
	campaign.GoalAmount = input.GoalAmount
	campaign.Perks = input.Perks
	campaign.UserID = input.User.ID
	campaign.Status = "unachieved"

	newCampaign, err := s.repository.Save(campaign)
	if err != nil {
		return newCampaign, err
	}

	return newCampaign, nil
}

func (s *service) UpdateCampaign(inputID GetCampaignDetailInput, data CreateCampaignInput) (Campaign, error) {
	campaign, err := s.repository.FindByID(inputID.ID)
	if err != nil {
		return campaign, err
	}

	if campaign.UserID != data.User.ID {
		return campaign, errors.New("Not an owner of the campaign")
	}
	campaign.Name = data.Name
	campaign.ShortDescription = data.ShortDescription
	campaign.Description = data.Description
	campaign.Perks = data.Perks
	campaign.GoalAmount = data.GoalAmount

	updatedCampaign, err := s.repository.Update(campaign)
	if err != nil {
		return updatedCampaign, err
	}
	return updatedCampaign, nil
}

func (s *service) SaveCampaignImage(input CreateCampaignImageInput, file string) (CampaignImage, error) {
	isPrimary := 0
	if input.IsPrimary == true {
		isPrimary = 1

		_, err := s.repository.MarkAllImagesAsNonPrimary(input.CampaignID)
		if err != nil {
			return CampaignImage{}, err
		}
	}
	campaignImage := CampaignImage{}
	campaignImage.CampaignID = input.CampaignID
	campaignImage.IsPrimary = isPrimary
	campaignImage.FileName = file

	newCampaignImage, err := s.repository.CreateImage(campaignImage)
	if err != nil {
		return newCampaignImage, err
	}

	return newCampaignImage, nil
}

func (s *service) GetPaginatedCampaigns(page, pageSize int) ([]Campaign, int, int, int, int, error) {
	totalCampaigns, err := s.repository.GetTotalCampaigns()
	if err != nil {
		return nil, 0, 0, 0, 0, err
	}

	totalPages := int(math.Ceil(float64(totalCampaigns) / float64(pageSize)))
	if page < 1 {
		page = 1
	}
	if page > totalPages {
		page = totalPages
	}
	offset := (page - 1) * pageSize
	campaigns, err := s.repository.GetPaginatedCampaigns(offset, pageSize)
	if err != nil {
		return nil, 0, 0, 0, 0, err
	}
	var nextPage, prevPage int
	if page < totalPages {
		nextPage = page + 1
	}
	if page > 1 {
		prevPage = page - 1
	}

	return campaigns, totalPages, page, nextPage, prevPage, nil
}

func (s *service) GetPaginatedCampaignsByUserID(userID, page, pageSize int) ([]Campaign, int, int, int, int, error) {
	totalCampaigns, err := s.repository.GetTotalCampaignsByUserID(userID)
	if err != nil {
		return nil, 0, 0, 0, 0, err
	}

	totalPages := int(math.Ceil(float64(totalCampaigns) / float64(pageSize)))
	if page < 1 {
		page = 1
	}
	if page > totalPages {
		page = totalPages
	}

	offset := (page - 1) * pageSize
	campaigns, err := s.repository.GetPaginatedCampaignsByUserID(userID, offset, pageSize)
	if err != nil {
		return nil, 0, 0, 0, 0, err
	}
	var nextPage, prevPage int
	if page < totalPages {
		nextPage = page + 1
	}

	if page > 1 {
		prevPage = page - 1
	}

	return campaigns, totalPages, page, nextPage, prevPage, nil
}

func (s *service) SearchCampaignsByName(name string) ([]Campaign, error) {
	campaigns, err := s.repository.SearchCampaignsByName(name)
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}
