package handler

import (
	"github.com/RianIhsan/raise-unity/helper"
	"github.com/RianIhsan/raise-unity/transaction"
	"github.com/RianIhsan/raise-unity/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

type transactionHandler struct {
	service transaction.Service
}

func NewTransactionHandler(service transaction.Service) *transactionHandler {
	return &transactionHandler{service}
}

func (h *transactionHandler) GetCampaignTransactions(c *gin.Context) {
	var input transaction.GetCampaignTransactionInput
	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.ErrorResponse("Failed to get campaign's transactions", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	currentUser := c.MustGet("CurrentUser").(user.User)
	input.User = currentUser
	transactions, err := h.service.GetTransactionByCampaignID(input)
	if err != nil {
		response := helper.ErrorResponse("Failed to get campaign's transactions", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ResponseWithData("Campaign details", transaction.FormatCampaignTransactions(transactions))
	c.JSON(http.StatusOK, response)
}
