package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Status struct {
	Temperature string `json:"temp,omitempty"`
	Humidity    string `json:"humidity,omitempty"`
}

func main() {
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

		fmt.Println("The room temperature is: " + status.Temperature + "ºC")
		fmt.Println("The room humidity is: " + status.Humidity + "%")

		ctx.String(http.StatusOK, "OK")
	})
	router.Run(":8080")

}
