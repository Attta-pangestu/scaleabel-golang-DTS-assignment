package model

type Order struct {
	OrderID      uint   `gorm:"primaryKey" json:"orderId"`
	CustomerName string `json:"customerName"`
	OrderedAt    string `json:"orderedAt"`
	Items        []Item `gorm:"foreignKey:OrderID" json:"items"`
}
