package main

import (
	"github.com/39penlight-api/mqtt"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	cl := mqtt.NewMQTTClient()
	defer cl.Disconnect()
	controller := NewColorController(cl)

	r := gin.Default()

	r.POST("color", func(context *gin.Context) {
		controller.ChangePenlightColor(context)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := r.Run(":" + port); err != nil {
		panic(err)
	}
}
