package main

import (
	"banjir_client/generator"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
    url := "http://localhost:8080/banjir"

    for {
        water := generator.GenerateWater()
        wind := generator.GenerateWind()

        reqBody := gin.H{"water": water, "wind": wind}

        reqBodyJSON, err := json.Marshal(reqBody)
        if err != nil {
            fmt.Println("Error encoding reqBody into JSON:", err)
            continue
        }

        reqBuffer := bytes.NewBuffer(reqBodyJSON)

        resp, err := http.Post(url, "application/json", reqBuffer)
        if err != nil {
            fmt.Println("Error while making request to the server:", err)
            continue
        }
        defer resp.Body.Close()

        body, err := io.ReadAll(resp.Body)
        if err != nil {
            fmt.Println("Error reading response body:", err)
        } else {
            fmt.Println(string(body))
        }

        time.Sleep(15 * time.Second)
    }
}