package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Status struct {
	Temperature string `json:"temp,omitempty"`
	Humidity    string `json:"humidity,omitempty"`
}

func main() {
	actualTemp := Status{}
	router := gin.Default()
	router.POST("/temp", func(ctx *gin.Context) {
		var temp Status

		if err := ctx.ShouldBindJSON(&temp); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}



		fmt.Println("The room temperature is: " + temp.Temperature + "ºC")

		ctx.String(http.StatusOK, temp.Temperature)
	})

	router.POST("/humidity", func(ctx *gin.Context) {
		var humidity Status

		if err := ctx.ShouldBindJSON(&humidity); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		fmt.Println("The room humidity is: " + humidity.Humidity + "%")

		ctx.String(http.StatusOK, humidity.Humidity)
	})

	router.POST("/status", func(ctx *gin.Context) {
		var status Status

		if err := ctx.ShouldBindJSON(&status); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		
		temperature, err := strconv.ParseFloat(status.Temperature, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		humidity, err := strconv.ParseFloat(status.Humidity, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if temperature >= 0 && humidity >= 0 {
			actualTemp.Temperature = status.Temperature
			actualTemp.Humidity = status.Humidity
		}

		fmt.Println("The room temperature is: " + status.Temperature + "ºC")
		fmt.Println("The room humidity is: " + status.Humidity + "%")

		ctx.String(http.StatusOK, "OK")
	})

	router.GET("/status", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"Temperature": actualTemp.Temperature, "Humidity": actualTemp.Humidity})
	})

	router.GET("/test", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello world")
	})
	router.Run(":8080")

}
