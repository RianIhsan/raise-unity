package handler

import (
	"context"
	"github.com/RianIhsan/raise-unity/campaign"
	"github.com/RianIhsan/raise-unity/helper"
	"github.com/RianIhsan/raise-unity/user"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
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

func (h *campaignHandler) UploadImage(c *gin.Context) {
	currentUser := c.MustGet("CurrentUser").(user.User)
	if err := c.ShouldBind(&currentUser.ID); err != nil {
		errors := helper.FormatValidationError(err)
		response := helper.ErrorResponse("Unauthorized", errors)
		c.JSON(http.StatusUnauthorized, response)
		return
	}
	var input campaign.CreateCampaignImageInput
	err := c.ShouldBind(&input)
	if err != nil {
		response := helper.ErrorResponse("Failed to upload campaign image", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	fileHeader, err := c.FormFile("image")
	file, err := fileHeader.Open()
	ctx := context.Background()
	urlCloudinary := os.Getenv("CLOUDINARY_URL")
	cldService, _ := cloudinary.NewFromURL(urlCloudinary)
	response, _ := cldService.Upload.Upload(ctx, file, uploader.UploadParams{})
	_, err = h.service.SaveCampaignImage(input, response.SecureURL)
	if err != nil {
		response := helper.ErrorResponse("Failed to upload campaign image", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	res := helper.SuccesResponse("Successfully added campaign image to this campaign ")
	c.JSON(http.StatusOK, res)
}
