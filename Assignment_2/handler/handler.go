package handler

import (
	"net/http"
	"orders_management/database"
	"orders_management/model"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetOrders(c *gin.Context) {
	db := database.GetDB()

	var orders []model.Order
	if err := db.Preload("Items").Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mendapatkan daftar pesanan"})
		return
	}

	c.JSON(http.StatusOK, orders)
}

func CreateOrder(c *gin.Context) {
	var newOrder model.Order
	if err := c.BindJSON(&newOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format JSON tidak valid"})
		return
	}

	newOrder.OrderedAt = parseTime(newOrder.OrderedAt)

	db := database.GetDB()

	if err := db.Create(&newOrder).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat pesanan"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Pesanan berhasil dibuat", "data": newOrder})
}

func UpdateOrder(c *gin.Context) {
	orderID := c.Param("orderId")
	var updatedOrder model.Order
	if err := c.BindJSON(&updatedOrder); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format JSON tidak valid"})
		return
	}

	updatedOrder.OrderID = parseOrderID(orderID)

	db := database.GetDB()

	updatedOrder.OrderedAt = parseTime(updatedOrder.OrderedAt)

	if err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&model.Order{}).Where("order_id = ?", orderID).Updates(&updatedOrder).Error; err != nil {
			return err
		}
		for i := range updatedOrder.Items {
			if err := tx.Model(&model.Item{}).Where("item_id = ?", updatedOrder.Items[i].ItemID).Updates(&updatedOrder.Items[i]).Error; err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memperbarui pesanan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Pesanan berhasil diperbarui", "data": updatedOrder})
}

func DeleteOrder(c *gin.Context) {
	orderID := c.Param("orderId")

	db := database.GetDB()

	if err := db.Where("order_id = ?", orderID).Delete(&model.Order{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus pesanan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Pesanan berhasil dihapus"})
}

func parseTime(timeStr string) string {
	parsedTime, err := time.Parse("2006-01-02T15:04:05-07:00", timeStr)
	if err != nil {
		return timeStr
	}
	return parsedTime.Format("2006-01-02 15:04:05")
}

func parseOrderID(orderID string) uint {
	orderIDInt, _ := strconv.Atoi(orderID)
	return uint(orderIDInt)
}
