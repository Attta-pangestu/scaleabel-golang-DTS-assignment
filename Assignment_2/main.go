package main

import (
	"orders_management/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/orders", handler.CreateOrder)
	router.GET("/orders", handler.GetOrders)
	router.PUT("/orders/:orderId", handler.UpdateOrder)
	router.DELETE("/orders/:orderId", handler.DeleteOrder)

	router.Run(":8080")
}
