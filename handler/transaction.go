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
	response := helper.ResponseWithData("Campaign transactions", transaction.FormatCampaignTransactions(transactions))
	c.JSON(http.StatusOK, response)
}
func (h *transactionHandler) GetUserTransactions(c *gin.Context) {
	currentUser := c.MustGet("CurrentUser").(user.User)
	userID := currentUser.ID

	transactions, err := h.service.GetTransactionByUserID(userID)
	if err != nil {
		response := helper.ErrorResponse("Failed to get user's transactions", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ResponseWithData("User's transaction", transaction.FormatUserTransactions(transactions))
	c.JSON(http.StatusOK, response)
}

func (h *transactionHandler) CreateTransaction(c *gin.Context) {
	var input transaction.CreateTransactionInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		response := helper.ErrorResponse("Failed to create transaction", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("CurrentUser").(user.User)
	input.User = currentUser

	newTransaction, err := h.service.CreateTransaction(input)
	if err != nil {
		response := helper.ErrorResponse("Failed to create transaction", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ResponseWithData("Success to create transaction", transaction.FormatTransaction(newTransaction))
	c.JSON(http.StatusOK, response)

}
