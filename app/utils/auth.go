package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/johncalvinroberts/furizu/app/users"
)

func Authenticate(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")

	if token == "" {
		RespondWithError(c, http.StatusUnauthorized, map[string]interface{}{
			"success": false,
		})
		return
	}

	decoded, err := FurizuJWT.ValidateFromToken(token)
	if err != nil {
		RespondWithError(c, http.StatusForbidden, map[string]interface{}{
			"success": false,
		})
		return
	}
	user, err := users.FindUserById(decoded["id"])
	if err != nil {
		RespondWithError(c, http.StatusForbidden, map[string]interface{}{
			"success": false,
			"message": "User not found",
		})
		return
	}
	c.Set("user", user)
	c.Next()
}
