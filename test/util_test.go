package test

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	helpers2 "github.com/ramdanariadi/grocery-product-service/main/helpers"
	"github.com/ramdanariadi/grocery-product-service/main/models"
	"github.com/ramdanariadi/grocery-product-service/main/repositories/product"
	utils2 "github.com/ramdanariadi/grocery-product-service/main/utils"
	"os"
	"path/filepath"
	"reflect"
	"sync"
	"testing"
	"time"
)

func Test_read_csv(t *testing.T) {
	abs, err := filepath.Abs("../others/product.csv")
	helpers2.PanicIfError(err)
	utils2.ProductsFromCSV(abs)
}

func Test_os_pwd(t *testing.T) {
	getwd, err := os.Getwd()
	helpers2.PanicIfError(err)

	fmt.Println(getwd)

	abs, err := filepath.Abs("../utils/product.csv")
	helpers2.PanicIfError(err)

	fmt.Println(abs)
}

func Test_insert_product_from_csv(t *testing.T) {
	db := utils2.GetDBConnection()

	productRepo := product.ProductRepositoryImpl{
		DB: db,
	}

	abs, err := filepath.Abs("../others/product.csv")
	if err != nil {
		fmt.Println("filepath abs error")
	}

	products := utils2.ProductsFromCSV(abs)

	group := sync.WaitGroup{}
	begin, err := productRepo.DB.Begin()

	if err != nil {
		fmt.Println("begin transaction error")
	}

	defer func(db *sql.DB) {
		helpers2.CommitOrRollback(begin)
		err := db.Close()
		if err != nil {
			fmt.Println("close db error")
		}
		fmt.Println("defer call", time.Now())
	}(db)

	for index, product := range products {
		id, _ := uuid.NewUUID()
		product.Id = id.String()
		go productRepo.SaveFromCSV(&group, context.Background(), begin, product, index)
		//if productRepo.SaveFromVCS(context.Background(), begin, product) {
		//	fmt.Println(product.Name, "success")
		//} else {
		//	fmt.Println(product.Name, "fail")
		//}
	}

	group.Wait()
	fmt.Println("done loop", time.Now())
}

func Test_insert_product_from_csv_with_channel(t *testing.T) {
	db := utils2.GetDBConnection()

	productRepo := product.ProductRepositoryImpl{
		DB: db,
	}

	abs, err := filepath.Abs("../others/product.csv")
	if err != nil {
		fmt.Println("filepath abs error")
	}
	productChanel := make(chan models.ProductModelCSV)

	group := sync.WaitGroup{}
	begin, err := productRepo.DB.Begin()

	if err != nil {
		fmt.Println("begin transaction error")
	}

	defer func(db *sql.DB) {
		helpers2.CommitOrRollback(begin)
		err := db.Close()
		if err != nil {
			fmt.Println("close db error")
		}
		close(productChanel)
		fmt.Println("defer call", time.Now())
	}(db)

	go productRepo.SaveFromCSVWithChannel(&group, context.Background(), begin, productChanel)
	utils2.ProductsFromCSVWithChannel(abs, productChanel)
	fmt.Println("done insert")
	group.Wait()
	fmt.Println("waitgroup done", time.Now())
}

func TestObjectEmpty(t *testing.T) {
	model := models.CategoryModel{}
	var n interface{}
	equal := reflect.DeepEqual(model, n)
	zero := reflect.ValueOf(model).IsZero()
	fmt.Println(zero)
	fmt.Println(equal)
}
