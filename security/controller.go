package security

import (
	"context"
	"database/sql"
	_ "embed"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"github.com/julienschmidt/httprouter"
	"go-tunas/customresponses"
	"go-tunas/helpers"
	"golang.org/x/crypto/bcrypt"
	"io"
	"net/http"
)

type SecurityController struct {
	Repository UserRepository
	Validator  *validator.Validate
}

func NewSecurityController(db *sql.DB) *SecurityController {
	return &SecurityController{
		Repository: UserRepository{
			DB: db,
		},
		Validator: validator.New(),
	}
}

//go:embed JWTSECRET
var jwtSecret []byte

func (controller SecurityController) Login(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	user := UserModel{}
	userByte, err := io.ReadAll(r.Body)
	helpers.PanicIfError(err)

	err = json.Unmarshal(userByte, &user)
	helpers.PanicIfError(err)

	tx, err := controller.Repository.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)

	if !controller.Repository.Login(user, context.Background(), tx) {
		customresponses.SendResponse(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"role": "user",
	})

	signedString, err := token.SignedString(jwtSecret)
	if err != nil {
		return
	}

	tokens := Token{
		AccessToken:  signedString,
		RefreshToken: "",
	}

	customresponses.SendResponse(w, tokens, http.StatusOK)
}

func (controller SecurityController) SignUp(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	user := UserModel{}
	userByte, err := io.ReadAll(r.Body)
	helpers.PanicIfError(err)

	err = json.Unmarshal(userByte, &user)
	helpers.PanicIfError(err)

	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	helpers.PanicIfError(err)

	user.Password = string(password)
	tx, err := controller.Repository.DB.Begin()
	helpers.PanicIfError(err)

	defer helpers.CommitOrRollback(tx)

	if controller.Repository.SignUp(user, context.Background(), tx) {
		customresponses.SendResponse(w, "Created", http.StatusCreated)
	}
}
