package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/johncalvinroberts/furizu/src/utils"
)

func Authenticate(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")

	if token == "" {
		utils.RespondWithError(c, http.StatusUnauthorized, map[string]interface{}{
			"success": false,
		})
		return
	}

	_, err := utils.FurizuJWT.ValidateFromToken(token)
	if err != nil {
		utils.RespondWithError(c, http.StatusForbidden, map[string]interface{}{
			"success": false,
		})
	}
	c.Next()
}
