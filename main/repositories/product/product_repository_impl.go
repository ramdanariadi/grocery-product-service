package product

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/ramdanariadi/grocery-product-service/main/helpers"
	"github.com/ramdanariadi/grocery-product-service/main/models"
	"sync"
)

type ProductRepositoryImpl struct {
	DB *sql.DB
}

func (repository ProductRepositoryImpl) FindById(context context.Context, tx *sql.Tx, id string) models.ProductModel {
	query := "SELECT products.id, name, price, per_unit, weight, category, category_id, description, products.image_url  " +
		"FROM products " +
		"JOIN category ON products.category_id = category.id " +
		"WHERE products.id = $1"
	row := tx.QueryRowContext(context, query, id)
	product := models.ProductModel{}
	err := row.Scan(&product.Id, &product.Name, &product.Price, &product.PerUnit, &product.Weight,
		&product.Category, &product.CategoryId,
		&product.Description, &product.ImageUrl)
	if err != nil {
		return models.ProductModel{}
	}
	//helpers.PanicIfError(err)
	return product
}

func (repository ProductRepositoryImpl) FindAll(context context.Context, tx *sql.Tx) *sql.Rows {
	query := "SELECT products.id, name, price, per_unit, weight, category, category_id, description, products.image_url  " +
		"FROM products " +
		"JOIN category ON products.category_id = category.id"

	rows, err := tx.QueryContext(context, query)
	helpers.PanicIfError(err)
	return rows
}

func (repository ProductRepositoryImpl) FindByCategory(context context.Context, tx *sql.Tx, id string) *sql.Rows {
	query := "SELECT products.id, name, price, per_unit, weight, category, category_id, description, products.image_url  " +
		"FROM products " +
		"JOIN category ON products.category_id = category.id " +
		"WHERE products.category_id = $1"
	rows, err := tx.QueryContext(context, query, id)
	helpers.PanicIfError(err)
	return rows
}

func (repository ProductRepositoryImpl) Save(context context.Context, tx *sql.Tx, product models.ProductModel) bool {
	sql := "INSERT INTO products(id, name, weight, price, per_unit, category_id, description, image_url, deleted) " +
		"VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9)"
	result, err := tx.ExecContext(context, sql, product.Id, product.Name, product.Weight, product.Price,
		product.PerUnit, product.CategoryId, product.Description, product.ImageUrl, false)
	helpers.PanicIfError(err)
	affected, err := result.RowsAffected()
	helpers.PanicIfError(err)
	return affected > 0
}

func (repository ProductRepositoryImpl) SaveFromCSV(waitgroup *sync.WaitGroup, context context.Context, tx *sql.Tx, saveModel models.ProductModelCSV, index int) bool {

	defer waitgroup.Done()

	waitgroup.Add(1)
	sqlInsert := "INSERT INTO " +
		"product(id,deleted,price,weight,category_id,per_unit,description,image_url,name) " +
		"VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)"
	result, err := tx.ExecContext(context, sqlInsert, saveModel.Id, saveModel.Deleted, saveModel.Price, saveModel.Weight, saveModel.CategoryId, saveModel.PerUnit, saveModel.Description, saveModel.ImageUrl.(string), saveModel.Name)
	if err != nil {
		fmt.Println("err exec context")
	}

	affected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("error rows affected")
	}
	if index%100 == 0 {
		fmt.Println(index)
	}
	return affected > 0
}

func (repository ProductRepositoryImpl) SaveFromCSVWithChannel(waitgroup *sync.WaitGroup, context context.Context, tx *sql.Tx, channel chan models.ProductModelCSV) bool {
	for i := 0; i < 10000; i++ {
		go func(index int) {
			defer waitgroup.Done()
			waitgroup.Add(1)

			saveModel := <-channel
			id, _ := uuid.NewUUID()
			saveModel.Id = id.String()
			var outerError error
			for {
				func(outerError *error) {
					defer func() {
						err := recover()
						if err != nil {
							*outerError = fmt.Errorf("error %v", err)
						}
					}()

				}(&outerError)

				sqlInsert := "INSERT INTO " +
					"product(id,deleted,price,weight,category_id,per_unit,description,image_url,name) " +
					"VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)"
				_, err := tx.ExecContext(context, sqlInsert, saveModel.Id, saveModel.Deleted, saveModel.Price, saveModel.Weight, saveModel.CategoryId, saveModel.PerUnit, saveModel.Description, saveModel.ImageUrl.(string), saveModel.Name)
				if err != nil {
					fmt.Println("err exec context : ", err)
					//panic(err)
				}

				if index%100 == 0 {
					fmt.Println(index)
				}

				if outerError == nil {
					break
				}
				//break
			}
		}(i)
	}
	return true
}

func (repository ProductRepositoryImpl) Update(context context.Context, tx *sql.Tx, product models.ProductModel) bool {
	sql := "UPDATE products SET name=$1, price=$2, weight=$3, category_id=$4, per_unit=$5," +
		"description=$6, image_url=$7" +
		"WHERE id = $8"
	result, err := tx.ExecContext(context, sql, product.Name, product.Price, product.Weight,
		product.Category, product.PerUnit, product.Description, product.ImageUrl, product.Id)
	helpers.PanicIfError(err)
	affected, err := result.RowsAffected()
	helpers.PanicIfError(err)

	return affected > 0
}

func (repository ProductRepositoryImpl) Delete(context context.Context, tx *sql.Tx, id string) bool {
	sql := "DELETE FROM products WHERE id = $1"
	result, err := tx.ExecContext(context, sql, id)
	helpers.PanicIfError(err)
	affected, err := result.RowsAffected()
	helpers.PanicIfError(err)
	return affected > 0
}
