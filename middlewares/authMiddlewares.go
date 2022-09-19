package middlewares

import (
	"go-pos-api/helpers"
	"go-pos-api/services"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(userServices services.UserServices, authService services.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			response := helpers.APIResponse(http.StatusUnauthorized, "Error", "Unauthorized", "Dont Have Authorization")
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}
		result, userId, err := authService.ValidateToken(tokenString)
		if err != nil && !result && userId == 0 {
			response := helpers.APIResponse(http.StatusUnauthorized, "Error", "Unauthorized", "Dont Have Authorization")
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		} else {
			user, err := userServices.GetUsersByID(userId)
			if err != nil {
				response := helpers.APIResponse(http.StatusUnauthorized, "Error", "Unauthorized", "Dont Have Authorization")
				c.AbortWithStatusJSON(http.StatusUnauthorized, response)
				return
			}
			c.Set("currentUser", user)
		}
	}
}
