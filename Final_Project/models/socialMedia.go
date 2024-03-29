package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type SocialMedia struct {
	GormModel
	Name           string `gorm:"not null" json:"name" form:"name" valid:"required~Name is required"`
	SocialMediaUrl string `gorm:"not null" json:"social_media_url" form:"social_media_url" valid:"required~Social media url is required, url~Url social media not valid"`
	UserID         uint   `gorm:"not null" json:"user_id"`
	User           *User
}

func (s *SocialMedia) TableName() string {
	return "tb_social_media"
}

func (s *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(s)
	if err != nil {
		return
	}

	return
}

func (s *SocialMedia) BeforeUpdate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(s)
	if err != nil {
		return
	}

	return
}
