package category

import (
	"context"
	"database/sql"
	"go-tunas/models/category"
)

type CategoryRepositoryImpl struct {
	DB *sql.DB
}

func (repository CategoryRepositoryImpl) FindById(context context.Context, tx *sql.Tx, id string) category.CategoryModel {
	//TODO implement me
	panic("implement me")
}

func (repository CategoryRepositoryImpl) FindAll(context context.Context, tx *sql.Tx) []category.CategoryModel {
	query := "select id, category, image_url from category"
	result, err := tx.QueryContext(context, query)
	if err != nil {
		panic("query error")
	}
	var categoriesModel []category.CategoryModel
	for result.Next() {
		cm := category.CategoryModel{}

		err := result.Scan(&cm.Id, &cm.Category, &cm.ImageUrl)
		if err != nil {
			panic("scan error")
		}
		cm.Deleted = false

		categoriesModel = append(categoriesModel, cm)

	}
	return categoriesModel
}

func (repository CategoryRepositoryImpl) Save(context context.Context, tx *sql.Tx, categoryModel category.CategoryModel) bool {
	//TODO implement me
	panic("implement me")
}

func (repository CategoryRepositoryImpl) Update(context context.Context, tx *sql.Tx, model category.CategoryModel, id string) bool {
	//TODO implement me
	panic("implement me")
}

func (repository CategoryRepositoryImpl) Delete(context context.Context, tx *sql.Tx, id string) bool {
	//TODO implement me
	panic("implement me")
}
