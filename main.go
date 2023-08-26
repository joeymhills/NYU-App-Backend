package main

import (
	"os"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.GET("/albums", func(c *gin.Context) {
			c.JSON(200, gin.H{
			"message": "pong",
		})	
	})
	var port = envPortOr("3000")

	router.Run("0.0.0.0" + port)

}

func envPortOr(port string) string {
	// If `PORT` variable in environment exists, return it
	if envPort := os.Getenv("PORT"); envPort != "" {
		return ":" + envPort
	}
	// Otherwise, return the value of `port` variable from function argument
	return ":" + port
}
