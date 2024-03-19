package handler

import (
	"net/http"
	"orders_management/database"
	"orders_management/model"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetOrders(c *gin.Context) {
	db, err := database.Connect()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal terhubung ke database"})
		return
	}
	defer database.Close(db)


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

    db, err := database.Connect()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal terhubung ke database"})
        return
    }
    defer database.Close(db)

    // Konversi format tanggal
    orderedAt, err := time.Parse("2006-01-02T15:04:05-07:00", newOrder.OrderedAt)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengonversi format tanggal"})
        return
    }
    newOrder.OrderedAt = orderedAt.Format("2006-01-02 15:04:05")

    // Simpan pesanan baru ke dalam database
    if err := db.Create(&newOrder).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat pesanan"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "Pesanan berhasil dibuat", "data": newOrder})
}



func UpdateOrder(c *gin.Context) {
    orderID := c.Param("orderId") // Mengambil orderId dari path
    
    var updatedOrder model.Order
    if err := c.BindJSON(&updatedOrder); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Format JSON tidak valid"})
        return
    }

    orderIDInt, _ := strconv.Atoi(orderID)
   
    updatedOrder.OrderID = uint(orderIDInt)

    db, err := database.Connect()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal terhubung ke database"})
        return
    }
    defer database.Close(db)

    // Konversi format tanggal
    orderedAt, err := time.Parse("2006-01-02T15:04:05-07:00", updatedOrder.OrderedAt)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengonversi format tanggal"})
        return
    }
    updatedOrder.OrderedAt = orderedAt.Format("2006-01-02 15:04:05")

    // Perbarui pesanan yang ada di database berdasarkan orderId
    if err := db.Where("order_id = ?", orderID).Updates(&updatedOrder).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memperbarui pesanan"})
        return
    }

    // Update item sesuai dengan ItemId
    for i := range updatedOrder.Items {
        if err := db.Model(&model.Item{}).Where("item_id = ?", updatedOrder.Items[i].ItemID).Updates(&updatedOrder.Items[i]).Error; err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memperbarui item"})
            return
        }
    }

    c.JSON(http.StatusOK, gin.H{"message": "Pesanan berhasil diperbarui", "data": updatedOrder})
}



func DeleteOrder(c *gin.Context) {
	orderID := c.Param("orderId")

	db, err := database.Connect()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal terhubung ke database"})
		return
	}
	defer database.Close(db)

	// Hapus pesanan dari database
	if err := db.Where("order_id = ?", orderID).Delete(&model.Order{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus pesanan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Pesanan berhasil dihapus"})
}

