package dto

type ProfileDTO struct {
	Name              string  `json:"name"`
	Username          string  `json:"username"`
	Email             string  `json:"email"`
	MobilePhoneNumber string  `json:"mobilePhoneNumber"`
	ProfileImageUrl   *string `json:"profileImageUrl"`
}
