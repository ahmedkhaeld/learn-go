package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello world",
	})
}
func SetupServer() *gin.Engine {
	r := gin.Default()
	r.GET("/", IndexHandler)
	return r
}

func main() {
	SetupServer().Run()
}
