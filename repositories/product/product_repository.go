package product

import (
	"context"
	"database/sql"
	"go-tunas/models"
	"go-tunas/requestBody"
	"sync"
)

type ProductRepository interface {
	FindById(context context.Context, tx *sql.Tx, id string) models.ProductModel
	FindAll(context context.Context, tx *sql.Tx) []models.ProductModel
	Save(context context.Context, tx *sql.Tx, saveRequest requestBody.ProductSaveRequest) bool
	SaveFromCSV(waitgroup *sync.WaitGroup, context context.Context, tx *sql.Tx, saveModel models.ProductModelCSV, index int) bool
	SaveFromCSVWithChannel(waitgroup *sync.WaitGroup, context context.Context, tx *sql.Tx, product chan models.ProductModelCSV) bool
	Update(context context.Context, tx *sql.Tx, updateRequest requestBody.ProductSaveRequest, id string) bool
	Delete(context context.Context, tx *sql.Tx, id string) bool
}

type TopProductRepository interface {
	FindById(context context.Context, tx *sql.Tx, id string) models.ProductModel
	FindAll(context context.Context, tx *sql.Tx) []models.ProductModel
	Save(context context.Context, tx *sql.Tx, saveRequest requestBody.TopProductSaveRequest) bool
	Update(context context.Context, tx *sql.Tx, updateRequest requestBody.TopProductSaveRequest, id string) bool
	Delete(context context.Context, tx *sql.Tx, id string) bool
}

type RcmdProductRepository interface {
	FindById(context context.Context, tx *sql.Tx, id string) models.ProductModel
	FindAll(context context.Context, tx *sql.Tx) []models.ProductModel
	Save(context context.Context, tx *sql.Tx, saveRequest requestBody.RcmdProductSaveRequest) bool
	Update(context context.Context, tx *sql.Tx, updateRequest requestBody.RcmdProductSaveRequest, id string) bool
	Delete(context context.Context, tx *sql.Tx, id string) bool
}
