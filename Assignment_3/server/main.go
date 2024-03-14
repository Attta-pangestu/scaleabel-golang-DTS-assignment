package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/update", updateHandler)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		for {
			time.Sleep(15 * time.Second)
			updateHandler(nil)
		}
	}()

	log.Fatal(srv.ListenAndServe())
}