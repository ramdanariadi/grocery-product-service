package product

import (
	"context"
	"database/sql"
	"github.com/ramdanariadi/grocery-product-service/main/models"
)

type ProductRepository interface {
	FindById(context context.Context, tx *sql.Tx, id string) *models.ProductModel
	FindByIds(context context.Context, tx *sql.Tx, ids []string) []*models.ProductModel
	FindAll(context context.Context, tx *sql.Tx) *sql.Rows
	FindByCategory(context context.Context, tx *sql.Tx, id string) *sql.Rows
	FindWhere(context context.Context, tx *sql.Tx, where string, value ...any) *sql.Rows
	Save(context context.Context, tx *sql.Tx, product *models.ProductModel) error
	Update(context context.Context, tx *sql.Tx, product *models.ProductModel) error
	Delete(context context.Context, tx *sql.Tx, id string) error
}
