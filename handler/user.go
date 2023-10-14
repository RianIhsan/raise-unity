package handler

import (
	"net/http"

	"github.com/RianIhsan/raise-unity/helper"
	"github.com/RianIhsan/raise-unity/user"
	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	payload := user.RegisterUserInput{}

	err := c.ShouldBindJSON(&payload)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errMessage := gin.H{
			"errors": errors,
		}
		response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", errMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.RegisterUser(payload)
	if err != nil {
		response := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := user.FormatUser(newUser, "eyawbidubawidlbalwdpa")

	response := helper.APIResponse("Account has been registered", http.StatusCreated, "success", formatter)
	c.JSON(http.StatusOK, response)
}
