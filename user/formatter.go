package user

type userFormatter struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Occupation string `json:"occupation"`
	Email      string `json:"email"`
	Token      string `json:"token"`
	Avatar     string `json:"avatar"`
}

func FormatUser(user User, token string) userFormatter {
	formatter := userFormatter{
		ID:         user.ID,
		Name:       user.Name,
		Occupation: user.Occupation,
		Email:      user.Email,
		Avatar:     user.Avatar,
		Token:      token,
	}

	return formatter
}

type VerifyEmailPayload struct {
	Email string `json:"email" binding:"required"`
	OTP   string `json:"otp" binding:"required"`
}
