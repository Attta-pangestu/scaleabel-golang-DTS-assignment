package handler

import (
	"banjir_server/json_writer"
	"fmt"
	"net/http"

	"banjir_server/generator" // Import package generator

	"github.com/gin-gonic/gin"
)



func UpdateHandler(c *gin.Context) {
    water := generator.GenerateWater()
    wind := generator.GenerateWind()
    status := generator.DetermineStatus(water, wind)

    // Membuat objek response
    response := json_writer.Response{
        Water:  water,
        Wind:   wind,
        Status: status,
    }

    // Mengirim response ke client
    c.JSON(http.StatusOK, response)

    // Menulis data response ke dalam file JSON
    err := json_writer.WriteResponseToJSON(response)
    if err != nil {
        // Jika terjadi error saat menulis ke dalam file JSON, tangani sesuai kebutuhan
        c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // Menampilkan hasil dari permintaan client di terminal
    fmt.Println("Hasil permintaan client:")
    fmt.Printf("Water: %d, Wind: %d, Status: %s\n", water, wind, status)
}
