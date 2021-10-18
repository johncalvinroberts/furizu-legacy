package utils

import (
	"os"

	"github.com/gin-gonic/gin"
)

var secure bool
var httpOnly bool

func init() {
	secure = os.Getenv("COOKIE_SECURE") == "true"
	httpOnly = os.Getenv("COOKIE_HTTP_ONLY") == "true"
}

func SetCookie(c *gin.Context, token string) {
	// NOTE: ttlMs set in .jwt.go
	c.SetCookie("tk", token, ttlMs, "/", "", secure, httpOnly)
}
