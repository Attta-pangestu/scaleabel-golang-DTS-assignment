package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

// Photo represents the model for an Photo
type Photo struct {
	GormModel
	Title    string    `gorm:"not null" json:"title" form:"title" valid:"required~Title is required"`
	Caption  string    `json:"caption" form:"caption"`
	PhotoUrl string    `gorm:"not null" json:"photo_url" form:"photo_url" valid:"required~Photo url is required, url~Url photo not valid"`
	UserID   uint      `gorm:"not null" json:"user_id"`
	// User     User      `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}


func (p *Photo) TableName() string {
	return "tb_photo"
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(p)
	if err != nil {
		return
	}

	return
}

func (p *Photo) BeforeUpdate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(p)
	if err != nil {
		return
	}

	return
}
