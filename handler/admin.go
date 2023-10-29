package handler

import (
	"github.com/RianIhsan/raise-unity/admin"
	"github.com/RianIhsan/raise-unity/auth"
	"github.com/RianIhsan/raise-unity/helper"
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
		response := helper.ErrorResponse("Error to get users", err)
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