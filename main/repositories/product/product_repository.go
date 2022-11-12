package product

import (
	"context"
	"database/sql"
	"github.com/ramdanariadi/grocery-be-golang/main/models"
	"github.com/ramdanariadi/grocery-be-golang/main/proto/product"
	"sync"
)

type ProductRepository interface {
	FindById(context context.Context, tx *sql.Tx, id string) models.ProductModel
	FindAll(context context.Context, tx *sql.Tx) *sql.Rows
	FindByCategory(context context.Context, tx *sql.Tx, id string) *sql.Rows
	Save(context context.Context, tx *sql.Tx, product product.Product) bool
	SaveFromCSV(waitgroup *sync.WaitGroup, context context.Context, tx *sql.Tx, saveModel models.ProductModelCSV, index int) bool
	SaveFromCSVWithChannel(waitgroup *sync.WaitGroup, context context.Context, tx *sql.Tx, product chan models.ProductModelCSV) bool
	Update(context context.Context, tx *sql.Tx, product product.Product) bool
	Delete(context context.Context, tx *sql.Tx, id string) bool
}

//type TopProductRepository interface {
//	FindById(context context.Context, tx *sql.Tx, id string) models.ProductModel
//	FindAll(context context.Context, tx *sql.Tx) []models.ProductModel
//	Save(context context.Context, tx *sql.Tx, saveRequest requestBody.TopProductSaveRequest) bool
//	Delete(context context.Context, tx *sql.Tx, id string) bool
//}
//
//type RcmdProductRepository interface {
//	FindById(context context.Context, tx *sql.Tx, id string) models.ProductModel
//	FindAll(context context.Context, tx *sql.Tx) []models.ProductModel
//	Save(context context.Context, tx *sql.Tx, saveRequest requestBody.RcmdProductSaveRequest) bool
//	Delete(context context.Context, tx *sql.Tx, id string) bool
//}
