package dto

type RegisterDTO struct {
	Email             string `json:"email"`
	Password          string `json:"password"`
	Username          string `json:"username"`
	MobilePhoneNumber string `json:"mobilePhoneNumber"`
}
