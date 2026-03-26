package main

import (
	"fmt"
	"net/http"
    "os" 

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "This won't build!",
		})
	})
	
	r.StartServer(":8081") 
}