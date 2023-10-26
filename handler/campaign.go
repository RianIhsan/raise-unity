package handler

import (
	"github.com/RianIhsan/raise-unity/campaign"
	"github.com/RianIhsan/raise-unity/helper"
	"github.com/RianIhsan/raise-unity/user"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type campaignHandler struct {
	service campaign.Service
}

func NewCampaignHandler(service campaign.Service) *campaignHandler {
	return &campaignHandler{service}
}

func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))

	campaigns, err := h.service.FindCampaigns(userID)
	if err != nil {
		response := helper.ErrorResponse("Error to get campaigns", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ResponseWithData("List of campaigns", campaign.FormatCampaigns(campaigns))
	c.JSON(http.StatusOK, response)

}

func (h *campaignHandler) GetCampaign(c *gin.Context) {
	var input campaign.GetCampaignDetailInput
	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.ErrorResponse("Failed to get campaign 1", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	campaignDetails, err := h.service.FindCampaignByID(input)
	if err != nil {
		response := helper.ErrorResponse("Failed to get campaign 2", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ResponseWithData("Campaign Details", campaign.FormatCampaignDetails(campaignDetails))
	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) CreateCampaign(c *gin.Context) {
	var input campaign.CreateCampaignInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		response := helper.ErrorResponse("Payload invalid", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("CurrentUser").(user.User)
	input.User = currentUser

	newCampaign, err := h.service.CreateCampaign(input)
	if err != nil {
		response := helper.ErrorResponse("Failed create campaign", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ErrorResponse("Success creating campaign", campaign.FormatCampaign(newCampaign))
	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) UpdateCampaign(c *gin.Context) {
	var input campaign.GetCampaignDetailInput
	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.ErrorResponse("Failed to update campaign", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	var data campaign.CreateCampaignInput
	err = c.ShouldBindJSON(&data)
	if err != nil {
		errors := helper.FormatValidationError(err)
		response := helper.ErrorResponse("Payload invalid", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	currentUser := c.MustGet("CurrentUser").(user.User)
	data.User = currentUser

	updatedCampaign, err := h.service.UpdateCampaign(input, data)
	if err != nil {
		response := helper.ErrorResponse("Failed to update campaign", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ResponseWithData("Success updating campaign", campaign.FormatCampaign(updatedCampaign))
	c.JSON(http.StatusOK, response)
}
