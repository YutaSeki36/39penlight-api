package main

import (
	"errors"
	"fmt"
	"github.com/39penlight-api/mqtt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Request struct {
	ColorCode string `json:"color_code"`
}

type Response struct {
	Result bool `json:"result"`
}

type ColorController struct {
	MQTTClient mqtt.ClientInterface
}

func NewColorController(client mqtt.ClientInterface) *ColorController {
	return &ColorController{MQTTClient: client}
}

func (controller *ColorController) ChangePenlightColor(c *gin.Context) {
	var request Request
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	colorCode, err := NewColorCode(request.ColorCode)
	if err != nil {
		c.JSON(500, fmt.Errorf("カラーコード生成に失敗しました. code: %s", request.ColorCode))
	}

	if !controller.MQTTClient.Publish(colorCode.ToString()) {
		c.JSON(500, errors.New("publishに失敗"))
	}

	resp := Response{}
	resp.Result = true

	c.JSON(200, resp)
}
