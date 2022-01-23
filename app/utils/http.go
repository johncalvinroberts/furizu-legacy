package utils

import (
	"context"
	"fmt"
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

func RespondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}

func GinContextFromContext(ctx context.Context) (*gin.Context, error) {
	ginContext := ctx.Value("GinContextKey")
	if ginContext == nil {
		err := fmt.Errorf("could not retrieve gin.Context")
		return nil, err
	}

	gc, ok := ginContext.(*gin.Context)
	if !ok {
		err := fmt.Errorf("gin.Context has wrong type")
		return nil, err
	}
	return gc, nil
}
