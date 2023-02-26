package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	cartModel "github.com/ramdanariadi/grocery-product-service/main/cart/model"
	categoryModel "github.com/ramdanariadi/grocery-product-service/main/category"
	"github.com/ramdanariadi/grocery-product-service/main/product/model"
	"github.com/ramdanariadi/grocery-product-service/main/setup"
	transactionModel "github.com/ramdanariadi/grocery-product-service/main/transaction/model"
	"github.com/ramdanariadi/grocery-product-service/main/utils"
	wishlistModel "github.com/ramdanariadi/grocery-product-service/main/wishlist/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	connection, err := setup.NewDbConnection()
	db, err := gorm.Open(postgres.New(postgres.Config{Conn: connection}))
	utils.PanicIfError(err)
	err = db.AutoMigrate(&categoryModel.Category{}, &model.Product{}, &wishlistModel.Wishlist{}, &cartModel.Cart{}, &transactionModel.Transaction{}, &transactionModel.TransactionDetail{})
	utils.LogIfError(err)

	router := gin.Default()

	categoryRoute := router.Group("api/v1/category")
	{
		categoryController := categoryModel.NewCategoryController(db)
		categoryRoute.GET("/", categoryController.FindAll)
		categoryRoute.GET("/:id", categoryController.FindById)
		categoryRoute.POST("/", categoryController.Save)
		categoryRoute.PUT("/:id", categoryController.Update)
		categoryRoute.DELETE("/:id", categoryController.Delete)
	}

	router.Run()
}
