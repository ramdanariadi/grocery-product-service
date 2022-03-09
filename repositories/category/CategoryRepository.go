package category

import (
	"context"
	"database/sql"
	"go-tunas/models/category"
)

type CategoryRepository interface {
	FindById(context context.Context, tx *sql.Tx, id string) category.CategoryModel
	FindAll(context context.Context, tx *sql.Tx) []category.CategoryModel
	Save(context context.Context, tx *sql.Tx, categoryModel category.CategoryModel) bool
	Update(context context.Context, tx *sql.Tx, model category.CategoryModel, id string) bool
	Delete(context context.Context, tx *sql.Tx, id string) bool
}
