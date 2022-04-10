package security

import (
	_ "embed"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/ramdanariadi/grocery-be-golang/main/customresponses"
	"github.com/ramdanariadi/grocery-be-golang/main/helpers"
	"net/http"
	"strings"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

//go:embed JWTSECRET
var token_secret []byte

func (middleware AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	apiKey := request.Header.Get("Authorization")

	if strings.HasPrefix(apiKey, "Bearer ") {
		token := strings.SplitAfter(apiKey, "Bearer ")[1]
		decodedJwt, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
				return token_secret, nil
			}
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		})
		helpers.PanicIfError(err)
		if claims, ok := decodedJwt.Claims.(jwt.MapClaims); ok && decodedJwt.Valid {

			if len(claims) > 0 {
				if claims["role"] == "admin" {
					middleware.Handler.ServeHTTP(writer, request)
					return
				}
			}
			customresponses.SendResponse(writer, "token claims not valid", http.StatusUnauthorized)
			return
		}
	}

	customresponses.SendResponse(writer, "Not authenticated", http.StatusUnauthorized)
	return

}

func SecureHandler(handle http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		apiKey := request.Header.Get("Authorization")

		if strings.HasPrefix(apiKey, "Bearer ") {
			token := strings.SplitAfter(apiKey, "Bearer ")[1]
			decodedJwt, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
					return token_secret, nil
				}
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			})
			helpers.PanicIfError(err)
			if claims, ok := decodedJwt.Claims.(jwt.MapClaims); ok && decodedJwt.Valid {

				if len(claims) > 0 {
					if claims["role"] == "user" {
						handle(writer, request)
						return
					}
				}
				customresponses.SendResponse(writer, "token claims not valid", http.StatusUnauthorized)
				return
			}
		}
		customresponses.SendResponse(writer, "Not authenticated", http.StatusUnauthorized)
		return
	}
}
