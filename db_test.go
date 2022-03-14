package main

import (
	"database/sql"
	_ "embed"
	"fmt"
	"github.com/golang-jwt/jwt"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

var secret []byte

func Test_signup(t *testing.T) {
	pass := "123qweasdzxc"
	password, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return
	}

	fmt.Println(string(password))
}

func Test_login(t *testing.T) {
	hashedPass := "$2a$10$8eYbu.CueBJ8vRoCFXw6MuJfLsUiG5/McNWMUY9g4gA358uK9U0za"
	realPass := "123qweasdzxc"
	err := bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(realPass))
	if err != nil {
		fmt.Println("pass miss match")
		return
	}
	fmt.Println("pass correct")
}

func Test_jwt(t *testing.T) {
	secret = []byte("secret")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"role": "user"})
	signedString, err := token.SignedString(secret)
	if err != nil {
		return
	}
	fmt.Println(signedString)
}

func Test_jwt_decode(t *testing.T) {
	secret = []byte("secret")

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIifQ.dtxWM6MIcgoeMgH87tGvsNDY6cHWL6MGW4LeYvnm1JA"

	parse, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})

	if err != nil {
		return
	}

	if claims, ok := parse.Claims.(jwt.MapClaims); ok && parse.Valid {
		fmt.Println(claims["foo"])
	} else {
		fmt.Println("error")
	}
}

func Test_nil(t *testing.T) {
	var n interface{}
	n = nil
	if n != nil {
		fmt.Println("not nil")
	}
}

func Test_db(t *testing.T) {
	connStr := "postgres://postgres:secret@localhost/DBTunasGrocery?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
	}

	rows, err := db.Query("select * from category")

	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		var category string
		var id string
		var image_url interface{}
		var deleted bool

		rowerr := rows.Scan(&category, &id, &deleted, &image_url)

		if rowerr != nil {
			panic(rowerr)
		}

		fmt.Println(image_url)
	}
}

//go:embed security/JWTSECRET
var jwtSecret []byte

func Test_remember_jwt(t *testing.T) {
	fmt.Println(string(jwtSecret))
}
