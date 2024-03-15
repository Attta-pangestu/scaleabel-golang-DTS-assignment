package main

import (
	"log"
	"net/http"
	"time"

	"banjir_server/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/banjir", handler.UpdateHandler)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		for {
			time.Sleep(15 * time.Second)
		}
	}()

	log.Fatal(srv.ListenAndServe())
}

