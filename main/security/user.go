package security

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"github.com/ramdanariadi/grocery-be-golang/main/helpers"
	"golang.org/x/crypto/bcrypt"
)

type UserModel struct {
	Username string `verified:"required"`
	Password string `verified:"required"`
}

type User interface {
	Login(user UserModel, context context.Context, tx *sql.Tx) bool
	SignUp(user UserModel, context context.Context, tx *sql.Tx) bool
	UserExist(username string, context context.Context, tx *sql.Tx) bool
}

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(DB *sql.DB) *UserRepository {
	return &UserRepository{DB: DB}
}

func (u UserRepository) Login(user UserModel, context context.Context, tx *sql.Tx) bool {
	var hashedPassword string
	query := "SELECT password FROM users where username = $1"
	row := tx.QueryRowContext(context, query, user.Username)

	err := row.Scan(&hashedPassword)
	if err != nil {
		return false
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(user.Password))
	helpers.PanicIfError(err)

	return true
}

func (u UserRepository) UserExist(username string, context context.Context, tx *sql.Tx) bool {
	query := "SELECT username FROM users where username = $1"
	row := tx.QueryRowContext(context, query, username)
	err := row.Scan(&username)
	return err != nil
}

func (u UserRepository) SignUp(user UserModel, context context.Context, tx *sql.Tx) bool {

	signUpSQL := "INSERT INTO users(id, username, password) values ($1,$2,$3)"

	newUUID, err := uuid.NewUUID()
	helpers.PanicIfError(err)

	execContext, insertErr := tx.ExecContext(context, signUpSQL, newUUID, user.Username, user.Password)
	helpers.PanicIfError(insertErr)

	affected, affectionErr := execContext.RowsAffected()
	helpers.PanicIfError(affectionErr)

	return affected > 0
}
