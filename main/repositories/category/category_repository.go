package category

import (
	"context"
	"database/sql"
	"github.com/ramdanariadi/grocery-be-golang/main/models"
	"github.com/ramdanariadi/grocery-be-golang/main/requestBody"
)

type CategoryRepository interface {
	FindById(context context.Context, tx *sql.Tx, id string) models.CategoryModel
	FindAll(context context.Context, tx *sql.Tx) []models.CategoryModel
	Save(context context.Context, tx *sql.Tx, saveRequest requestBody.CategorySaveRequest) bool
	Update(context context.Context, tx *sql.Tx, updateRequest requestBody.CategorySaveRequest, id string) bool
	Delete(context context.Context, tx *sql.Tx, id string) bool
}
