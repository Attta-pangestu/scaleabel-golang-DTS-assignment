package models

import (
	"time"

	"gorm.io/gorm"
)

type GormModel struct {
    ID        uint           `gorm:"primary_key;autoIncrement" json:"id"`
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt gorm.DeletedAt `gorm:"index"`
}
