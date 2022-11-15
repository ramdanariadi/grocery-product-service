package category

import (
	"context"
	"database/sql"
	"github.com/ramdanariadi/grocery-product-service/main/models"
)

type CategoryRepository interface {
	FindById(context context.Context, tx *sql.Tx, id string) models.CategoryModel
	FindAll(context context.Context, tx *sql.Tx) *sql.Rows
	Save(context context.Context, tx *sql.Tx, saveRequest models.CategoryModel) bool
	Update(context context.Context, tx *sql.Tx, updateRequest models.CategoryModel, id string) bool
	Delete(context context.Context, tx *sql.Tx, id string) bool
}
