package handler

import (
	"github.com/RianIhsan/raise-unity/campaign"
	"github.com/RianIhsan/raise-unity/helper"
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
	//user id di ubah ke string
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
