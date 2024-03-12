package model

type Item struct {
	ItemID      uint   `gorm:"primaryKey" json:"itemId"`
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	OrderID     uint   `json:"orderId"`
}
