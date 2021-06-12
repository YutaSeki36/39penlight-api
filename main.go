package main

import (
	"github.com/39penlight-api/mqtt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"os"
	"time"
)

func main() {
	cl := mqtt.NewMQTTClient()
	defer cl.Disconnect()
	controller := NewColorController(cl)

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"https://zen-fermat-8df333.netlify.app",
			"http://localhost:3000",
		},
		AllowMethods: []string{
			"POST",
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	}))

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
