package archives

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindMany(c *gin.Context) {
	fmt.Printf("%T", c.Request)
	data := map[string]interface{}{
		"success": true,
	}
	c.IndentedJSON(http.StatusAccepted, data)
	return
}

func Create(c *gin.Context) {
	fmt.Printf("%T", c.Request)
	data := map[string]interface{}{
		"success": true,
	}
	c.IndentedJSON(http.StatusAccepted, data)
	return
}

func FindOne(c *gin.Context) {
	fmt.Printf("%T", c.Request)
	data := map[string]interface{}{
		"success": true,
	}
	c.IndentedJSON(http.StatusAccepted, data)
	return
}

func DestroyOne(c *gin.Context) {
	fmt.Printf("%T", c.Request)
	data := map[string]interface{}{
		"success": true,
	}
	c.IndentedJSON(http.StatusAccepted, data)
	return
}
