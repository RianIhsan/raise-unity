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

	existingUser, err := h.userService.GetUserByEmail(payload.Email)
	if err == nil && existingUser.ID > 0 {
		response := helper.APIResponse("Email already exists", http.StatusConflict, "error", nil)
		c.JSON(http.StatusConflict, response)
		return
	}

	newUser, err := h.userService.RegisterUser(payload)
	if err != nil {
		response := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := user.FormatUser(newUser, "eyawbidubawidlbalwdpa")

	response := helper.APIResponse("Berhasil Mendaftar, silahkan cek email untuk verifikasi OTP", http.StatusCreated, "success", formatter)
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

func (h *userHandler) VerifyEmail(c *gin.Context) {
	var payload user.VerifyEmailPayload

	if err := c.ShouldBindJSON(&payload); err != nil {
		response := helper.APIResponse("Invalid payload request", http.StatusUnprocessableEntity, "error", err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err := h.userService.VerifyEmail(payload.Email, payload.OTP)
	if err != nil {
		response := helper.APIResponse("Email verification failed", http.StatusUnprocessableEntity, "error", err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.APIResponse("Email verified successfully", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) ResendOTP(c *gin.Context) {
	var input user.ResendOTPInput

	if err := c.ShouldBindJSON(&input); err != nil {
		errors := helper.FormatValidationError(err)
		errMessage := gin.H{
			"errors": errors,
		}
		response := helper.APIResponse("Resend OTP failed", http.StatusUnprocessableEntity, "error", errMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	otp, err := h.userService.ResendOTP(input.Email)
	if err != nil {
		response := helper.APIResponse("Resend OTP failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err = helper.SendOTPByEmail(input.Email, otp.OTP)
	if err != nil {
		response := helper.APIResponse("Error sending OTP", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := helper.APIResponse("OTP has been resent", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}
