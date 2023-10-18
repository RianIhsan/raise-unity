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

func (h *userHandler) Login(c *gin.Context) {
	var input user.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{
			"error": errors,
		}

		response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedinUser, err := h.userService.Login(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Login Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := user.FormatUser(loggedinUser, "eyawbidubawidlbalwdpa")

	response := helper.APIResponse("Succesfully Loggedin", http.StatusOK, "succes", formatter)
	c.JSON(http.StatusOK, response)

}
