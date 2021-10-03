package files

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Cmd(c *gin.Context) {
	fmt.Printf("%T", c.Request)
	data := map[string]interface{}{
		"success": true,
	}
	c.IndentedJSON(http.StatusAccepted, data)
	return
}

func Qry(c *gin.Context) {
	fmt.Printf("%T", c.Request)
	data := map[string]interface{}{
		"success": true,
	}
	c.IndentedJSON(http.StatusAccepted, data)
	return
}
