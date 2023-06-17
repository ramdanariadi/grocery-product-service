package dto

type TokenDTO struct {
	AccessToken  string      `json:"accessToken"`
	RefreshToken string      `json:"refreshToken"`
	User         *ProfileDTO `json:"user"`
}
