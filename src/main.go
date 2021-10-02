package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getHello(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Jhello")
}

func main() {
	router := gin.Default()
	router.GET("/hello", getHello)
	router.Run("localhost:3000")
}
