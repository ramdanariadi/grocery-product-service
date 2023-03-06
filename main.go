package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	cart "github.com/ramdanariadi/grocery-product-service/main/cart/model"
	"github.com/ramdanariadi/grocery-product-service/main/category"
	"github.com/ramdanariadi/grocery-product-service/main/exception"
	"github.com/ramdanariadi/grocery-product-service/main/product"
	"github.com/ramdanariadi/grocery-product-service/main/setup"
	transaction "github.com/ramdanariadi/grocery-product-service/main/transaction/model"
	"github.com/ramdanariadi/grocery-product-service/main/utils"
	wishlist "github.com/ramdanariadi/grocery-product-service/main/wishlist/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	connection, err := setup.NewDbConnection()
	db, err := gorm.Open(postgres.New(postgres.Config{Conn: connection}))
	utils.PanicIfError(err)
	err = db.AutoMigrate(&category.Category{}, &product.Product{}, &wishlist.Wishlist{}, &cart.Cart{}, &transaction.Transaction{}, &transaction.TransactionDetail{})
	utils.LogIfError(err)

	router := gin.Default()
	router.Use(gin.CustomRecovery(exception.Handler))

	categoryRoute := router.Group("api/v1/category")
	{
		categoryController := category.NewCategoryController(db)
		categoryRoute.POST("/", categoryController.Save)
		categoryRoute.GET("/:id", categoryController.FindById)
		categoryRoute.GET("/", categoryController.FindAll)
		categoryRoute.PUT("/:id", categoryController.Update)
		categoryRoute.DELETE("/:id", categoryController.Delete)
	}

	productRoute := router.Group("api/v1/product")
	{
		productController := product.NewProductController(db)
		productRoute.POST("/", productController.Save)
		productRoute.GET("/:id", productController.FindById)
		productRoute.GET("/", productController.FindAll)
		productRoute.PUT("/:id", productController.Update)
		productRoute.DELETE("/:id", productController.Delete)
		productRoute.PUT("/top/:id", productController.SetTopProduct)
		productRoute.PUT("/recommendation/:id", productController.SetRecommendationProduct)
	}
	err = router.Run()
	utils.LogIfError(err)
}
