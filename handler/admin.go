package handler

import (
	"github.com/RianIhsan/raise-unity/admin"
	"github.com/RianIhsan/raise-unity/auth"
	"github.com/RianIhsan/raise-unity/helper"
	"github.com/RianIhsan/raise-unity/transaction"
	"github.com/RianIhsan/raise-unity/user"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type adminHandler struct {
	service     admin.Service
	authService auth.Service
}

func NewAdminHandler(service admin.Service, authService auth.Service) *adminHandler {
	return &adminHandler{service, authService}
}

func (h *adminHandler) GetAllUsers(c *gin.Context) {
	currentUser := c.MustGet("CurrentUser").(user.User)
	if currentUser.Role != "admin" {
		response := helper.GeneralResponse("Access denied")
		c.JSON(http.StatusUnauthorized, response)
		return
	}
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	var users []user.User
	var totalPages, currentPage, nextPage, prevPage int
	var err error

	searchNameUser := c.Query("name")
	if searchNameUser != "" {
		users, err = h.service.SearchUserByName(searchNameUser)
	} else {
		users, totalPages, currentPage, nextPage, prevPage, err = h.service.GetUsersPagination(page, pageSize)
	}
	if err != nil {
		response := helper.ErrorResponse("Error to get users", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	var nonAdminUsers []user.User
	for _, u := range users {
		if u.Role != "admin" {
			nonAdminUsers = append(nonAdminUsers, u)
		}
	}
	if totalPages > 1 {
		if currentPage < totalPages {
			nextPage = currentPage + 1
		} else {
			nextPage = -1
		}
		if currentPage > 1 {
			prevPage = currentPage - 1
		} else {
			prevPage = -1
		}
	}

	response := helper.ResponseWithPaginationAndNextPrev("List of users", admin.FormatterUsers(nonAdminUsers), currentPage, totalPages, nextPage, prevPage)
	c.JSON(http.StatusOK, response)
}

func (h *adminHandler) GetAllUsersTransactions(c *gin.Context) {
	currentUser := c.MustGet("CurrentUser").(user.User)
	if currentUser.Role != "admin" {
		response := helper.GeneralResponse("Access denied")
		c.JSON(http.StatusUnauthorized, response)
		return
	}
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	var transactions []transaction.Transaction
	var totalPages, currentPage, nextPage, prevPage int
	var err error

	searchTransactionByUsername := c.Query("user_name")
	if searchTransactionByUsername != "" {
		transactions, err = h.service.SearchTransactionByUsername(searchTransactionByUsername)
	} else {
		transactions, totalPages, currentPage, nextPage, prevPage, err = h.service.GetTransactionsPagination(page, pageSize)
	}
	if err != nil {
		response := helper.ErrorResponse("Error to get transactions", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if totalPages > 1 {
		if currentPage < totalPages {
			nextPage = currentPage + 1
		} else {
			nextPage = -1
		}
		if currentPage > 1 {
			prevPage = currentPage - 1
		} else {
			prevPage = -1
		}
	}

	response := helper.ResponseWithPaginationAndNextPrev("List of users", admin.FormatterTransactions(transactions), currentPage, totalPages, nextPage, prevPage)
	c.JSON(http.StatusOK, response)
}

func (h *adminHandler) DeleteUser(c *gin.Context) {
	currentUser := c.MustGet("CurrentUser").(user.User)
	if currentUser.Role != "admin" {
		response := helper.GeneralResponse("Access denied")
		c.JSON(http.StatusUnauthorized, response)
		return
	}
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response := helper.ErrorResponse("Failed get user", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	user, err := h.service.FindUserById(userId)
	if err != nil {
		response := helper.ErrorResponse("Failed get user", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.service.DeleteUserById(user.ID)
	if err != nil {
		response := helper.ErrorResponse("Failed delete user", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.SuccesResponse("Success Delete user")
	c.JSON(http.StatusOK, response)
}

func (h *adminHandler) DeleteCampaign(c *gin.Context) {
	currentUser := c.MustGet("CurrentUser").(user.User)
	if currentUser.Role != "admin" {
		response := helper.GeneralResponse("Access denied")
		c.JSON(http.StatusUnauthorized, response)
		return
	}
	campaign, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response := helper.ErrorResponse("Failed get campaign", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	campaignId, err := h.service.FindCampaignById(campaign)
	if err != nil {
		response := helper.ErrorResponse("Failed get campaign", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.service.DeleteCampaignById(campaignId.ID)
	if err != nil {
		response := helper.ErrorResponse("Failed delete campaign", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.GeneralResponse("Success Delete Campaign")
	c.JSON(http.StatusOK, response)
}
