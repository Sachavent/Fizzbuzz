package main

import (
	"net/http"

	"github.com/Fizzbuzz/models"
	"github.com/Fizzbuzz/services"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Create map between request object and number of time the request has been call
	var mapNumberRequest = make(map[models.FizzBuzzParams]int)

	// Return fizzbuzz result
	router.POST("/fizzbuzz", func(context *gin.Context) {
		var params models.FizzBuzzParams

		// Validate input paramters (matching to FizzBuzzParams model)
		if err := context.ShouldBindJSON(&params); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		result := services.GetFizzbuzzResult(params, mapNumberRequest)

		context.JSON(200, gin.H{
			"result": result,
		})
	})

	// Return most requested route and associated params
	router.GET("/data", func(context *gin.Context) {

		params, numberRequest, err := services.MostRequestedRequest(mapNumberRequest)

		if err != nil {
			context.JSON(err.StatusCode, gin.H{
				"message": err.Message,
			})
			return
		}

		context.JSON(200, gin.H{
			"params":        params,
			"numberRequest": numberRequest,
		})
	})

	// Listen and serve on 0.0.0.0:8080
	router.Run()
}
