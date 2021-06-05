package main

import (
	"github.com/39penlight-api/mqtt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func main() {
	cl := mqtt.NewMQTTClient()
	//cl.Publish("start")
	defer cl.Disconnect()
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := r.Run(":" + port); err != nil {
		panic(err)
	}
}
