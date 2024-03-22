package models

import "time"

type LoginResponse struct {
	Token string `json:"token"`
}

type ResponseFailed struct {
	Message string `json:"message"`
}

type ResponseFailedUnauthorized struct {
	Message string `json:"message"`
}


type UserResponse struct {
	Email     string    `json:"email"`
	Username  string    `json:"username"`
}


type PhotoResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      UserResponse `json:"user" gorm:"foreignKey:UserID"`
}


