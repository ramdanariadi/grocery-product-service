package user

import "github.com/ramdanariadi/grocery-product-service/main/user/dto"

type Service interface {
	Login(dto *dto.LoginDTO) *dto.TokenDTO
	Register(dto *dto.RegisterDTO) *dto.TokenDTO
	Token(dto dto.TokenDTO) *dto.TokenDTO
}
