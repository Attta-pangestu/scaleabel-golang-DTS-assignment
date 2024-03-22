package models

// RequestUserLogin represents the model for an User Login
type RequestUserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RequestUserRegister struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Age      uint   `json:"age"`
}


type RequestUserUpdate struct {
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
}

type RequestSocialMedia struct {
	Name           string `json:"name"`
	SocialMediaUrl string `json:"social_media_url"`
}

type RequestPhoto struct {
    Title    string `json:"title" form:"title" valid:"required~Title is required"`
    Caption  string `json:"caption" form:"caption"`
    PhotoUrl string `json:"photo_url" form:"photo_url" valid:"required~Photo url is required, url~Url photo not valid"`
}

type RequestComment struct {
	 Message string `json:"message"`
	 PhotoID uint   `json:"photo_id"`
}
