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
	c.SetCookie("tk", token, jwtTtlMs, "/", "", secure, httpOnly)
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

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: check gin mode, only turn on if gin mode is debug
		c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
