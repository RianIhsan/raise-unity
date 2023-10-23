package middleware

import (
	"net/http"
	"strings"

	"github.com/RianIhsan/raise-unity/auth"
	"github.com/RianIhsan/raise-unity/helper"
	"github.com/RianIhsan/raise-unity/user"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.GeneralResponse("Unauthorized 1")
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}
		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.GeneralResponse("Unauthorized 2")
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := helper.GeneralResponse("Unauthorized 3")
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := int(claim["user_id"].(float64))

		user, err := userService.GetUserByID(userID)
		if err != nil {
			response := helper.GeneralResponse("Unauthorized 5")
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		c.Set("CurrentUser", user)
	}

}
