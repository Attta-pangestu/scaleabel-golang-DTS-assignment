package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

// Comment represents the model for an Comment
type Comment struct {
	GormModel
	Message string `gorm:"not null" json:"message" form:"message" valid:"required~Comment message is required"`
	PhotoID uint   `gorm:"not null" json:"photo_id" form:"photo_id" valid:"required~Photo is required"`
	UserID  uint   `gorm:"not null" json:"user_id"`
	Photo *Photo 
	User  *User  
	
}

func (c *Comment) TableName() string {
	return "tb_comments"
}

func (c *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(c)
	if err != nil {
		return
	}

	return
}

func (c *Comment) BeforeUpdate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(c)
	if err != nil {
		return
	}

	return
}
