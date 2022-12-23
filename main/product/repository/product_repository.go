package repository

import (
	"context"
	"database/sql"
	"github.com/ramdanariadi/grocery-product-service/main/product/model"
)

type ProductRepository interface {
	FindById(context context.Context, tx *sql.Tx, id string) *model.ProductModel
	FindByIds(context context.Context, tx *sql.Tx, ids []string) []*model.ProductModel
	FindAll(context context.Context, tx *sql.Tx) *sql.Rows
	FindByCategory(context context.Context, tx *sql.Tx, id string) *sql.Rows
	FindWhere(context context.Context, tx *sql.Tx, where string, value ...any) *sql.Rows
	Save(context context.Context, tx *sql.Tx, product *model.ProductModel) error
	Update(context context.Context, tx *sql.Tx, product *model.ProductModel) error
	Delete(context context.Context, tx *sql.Tx, id string) error
}
