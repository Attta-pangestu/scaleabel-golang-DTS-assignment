package main

import (
	"log"
	"net/http"
	"time"

	"banjir_server/handler" // Sesuaikan impor package

	"github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    router.GET("/banjir", func(c *gin.Context) {
        handler.UpdateHandler(c) // Memanggil fungsi UpdateHandler dengan objek *gin.Context yang benar
    })

    srv := &http.Server{
        Addr:    ":8080",
        Handler: router,
    }

    go func() {
        for {
            time.Sleep(15 * time.Second)
            // handler.UpdateHandler(router) // Jangan gunakan router di sini, gunakan objek *gin.Context yang benar
        }
    }()

    log.Fatal(srv.ListenAndServe())
}
