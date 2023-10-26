package handler

import (
	"context"
	"net/http"
	"os"

	"github.com/RianIhsan/raise-unity/auth"
	"github.com/RianIhsan/raise-unity/helper"
	"github.com/RianIhsan/raise-unity/user"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	payload := user.RegisterUserInput{}

	err := c.ShouldBindJSON(&payload)
	if err != nil {
		errors := helper.FormatValidationError(err)

		response := helper.ErrorResponse("Register account failed", errors)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	existingUser, err := h.userService.GetUserByEmail(payload.Email)
	if err == nil && existingUser.ID > 0 {
		response := helper.ErrorResponse("Email alredy exist", err.Error())
		c.JSON(http.StatusConflict, response)
		return
	}

	_, err = h.userService.RegisterUser(payload)
	if err != nil {
		response := helper.ErrorResponse("Register account failed", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.SuccesResponse("Register account successfully, please check your email for OTP")
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) Login(c *gin.Context) {
	var input user.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)

		response := helper.ErrorResponse("Login failed", errors)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedinUser, err := h.userService.Login(input)
	if err != nil {
		response := helper.ErrorResponse("Login failed", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := h.authService.GenerateToken(loggedinUser.ID)
	if err != nil {
		response := helper.ErrorResponse("Login  failed", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := user.FormatUser(loggedinUser, token)

	response := helper.ResponseWithData("Succesfully Loggedin", formatter)
	c.JSON(http.StatusOK, response)

}

func (h *userHandler) VerifyEmail(c *gin.Context) {
	var payload user.VerifyEmailPayload

	if err := c.ShouldBindJSON(&payload); err != nil {
		response := helper.ErrorResponse("VerifyEmail failed", err)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	err := h.userService.VerifyEmail(payload.Email, payload.OTP)
	if err != nil {
		response := helper.ErrorResponse("VerifyEmail failed", err)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := helper.SuccesResponse("Email verified successfully")
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) ResendOTP(c *gin.Context) {
	var input user.ResendOTPInput

	if err := c.ShouldBindJSON(&input); err != nil {
		errors := helper.FormatValidationError(err)
		response := helper.ErrorResponse("Resend OTP failed", errors)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	otp, err := h.userService.ResendOTP(input.Email)
	if err != nil {
		response := helper.ErrorResponse("Resend OTP failed", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err = helper.SendOTPByEmail(input.Email, otp.OTP)
	if err != nil {
		response := helper.ErrorResponse("Error sending OTP", err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := helper.SuccesResponse("OTP has been resend")
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) UploadAvatar(c *gin.Context) {

	currentUser := c.MustGet("CurrentUser").(user.User)

	if err := c.ShouldBind(&currentUser.ID); err != nil {
		response := helper.ErrorResponse("Error Payload", err.Error())
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	fileHeader, _ := c.FormFile("avatar")
	file, _ := fileHeader.Open()
	ctx := context.Background()
	urlCloudinary := os.Getenv("CLOUDINARY_URL")
	cldService, _ := cloudinary.NewFromURL(urlCloudinary)
	response, _ := cldService.Upload.Upload(ctx, file, uploader.UploadParams{})

	_, err := h.userService.SaveAvatar(currentUser.ID, response.SecureURL)
	if err != nil {
		response := helper.ErrorResponse("Error update avatar", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	res := helper.UpdateAvatarRes("Success update avatar", response.SecureURL)

	c.JSON(http.StatusOK, res)

}
