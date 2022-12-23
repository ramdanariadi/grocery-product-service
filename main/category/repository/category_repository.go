package repository

import (
	"context"
	"database/sql"
	"github.com/ramdanariadi/grocery-product-service/main/category/model"
)

type CategoryRepository interface {
	FindById(context context.Context, tx *sql.Tx, id string) *model.CategoryModel
	FindAll(context context.Context, tx *sql.Tx) *sql.Rows
	Save(context context.Context, tx *sql.Tx, saveRequest *model.CategoryModel) error
	Update(context context.Context, tx *sql.Tx, updateRequest *model.CategoryModel, id string) error
	Delete(context context.Context, tx *sql.Tx, id string) error
}
