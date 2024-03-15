package handler

import (
	"fmt"
	"net/http"

	"banjir_server/generator"
	"banjir_server/json_writer"

	"github.com/gin-gonic/gin"
)

func UpdateHandler(c *gin.Context) {
	var requestData struct {
		Water int `json:"water"`
		Wind  int `json:"wind"`
	}
	if err := c.BindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	water := requestData.Water
	wind := requestData.Wind

	status := generator.DetermineStatus(water, wind)

	response := json_writer.Response{
		Water:  water,
		Wind:   wind,
		Status: status,
	}

	c.JSON(http.StatusOK, response)

	err := json_writer.WriteResponseToJSON(response)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("Hasil permintaan client:")
	fmt.Printf("Water: %d, Wind: %d, Status: %s\n", water, wind, status)
}