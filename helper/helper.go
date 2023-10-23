package helper

import (
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/wneessen/go-mail"
)

type succesResponse struct {
	Message string `json:"message"`
}

type updateAvatarRes struct {
	Message string `json:"message"`
	Avatar  string `json:"avatar"`
}

func SuccesResponse(message string) succesResponse {
	messageRes := succesResponse{
		Message: message,
	}

	return messageRes
}
func UpdateAvatarRes(message string, avatar string) updateAvatarRes {
	messageRes := updateAvatarRes{
		Message: message,
		Avatar:  avatar,
	}

	return messageRes
}

type generalResponse struct {
	Message string `json:"message"`
}

func GeneralResponse(message string) generalResponse {
	messageRes := generalResponse{
		Message: message,
	}

	return messageRes
}

type errorResponse struct {
	Message string `json:"message"`
	Error   any    `json:"error"`
}

func ErrorResponse(message string, err any) errorResponse {
	messageRes := errorResponse{
		Message: message,
		Error:   err,
	}

	return messageRes
}

type responseWithData struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ResponseWithData(message string, data interface{}) responseWithData {
	messageRes := responseWithData{
		Message: message,
		Data:    data,
	}

	return messageRes
}

func FormatValidationError(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}

func SendOTPByEmail(email, otp string) error {

	secret_user := os.Getenv("SMTP_USER")
	secret_pass := os.Getenv("SMTP_PASS")
	secret_port := os.Getenv("SMTP_PORT")

	convPort, err := strconv.Atoi(secret_port)
	if err != nil {
		return err
	}

	m := mail.NewMsg()
	if err := m.From(secret_user); err != nil {
		return err
	}
	if err := m.To(email); err != nil {
		return err
	}
	m.Subject("Verifikasi Email - RAISE UNITY")
	m.SetBodyString(mail.TypeTextPlain, "Kode OTP anda adalah : "+otp)

	c, err := mail.NewClient("smtp.gmail.com", mail.WithPort(convPort), mail.WithSMTPAuth(mail.SMTPAuthPlain), mail.WithUsername(secret_user), mail.WithPassword(secret_pass))
	if err != nil {
		return err
	}
	if err := c.DialAndSend(m); err != nil {
		return err
	}

	return nil
}

func GenerateRandomOTP(otpLent int) string {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	const n = "0123456789"

	otp := make([]byte, otpLent)
	for i := range otp {
		otp[i] = n[r.Intn(len(n))]
	}

	return string(otp)

}
