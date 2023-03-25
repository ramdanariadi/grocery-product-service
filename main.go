package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/ramdanariadi/grocery-product-service/main/cart"
	"github.com/ramdanariadi/grocery-product-service/main/category"
	"github.com/ramdanariadi/grocery-product-service/main/exception"
	"github.com/ramdanariadi/grocery-product-service/main/product"
	"github.com/ramdanariadi/grocery-product-service/main/setup"
	transaction "github.com/ramdanariadi/grocery-product-service/main/transaction/model"
	"github.com/ramdanariadi/grocery-product-service/main/user"
	"github.com/ramdanariadi/grocery-product-service/main/utils"
	wishlist "github.com/ramdanariadi/grocery-product-service/main/wishlist"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	connection, err := setup.NewDbConnection()
	db, err := gorm.Open(postgres.New(postgres.Config{Conn: connection}))
	utils.PanicIfError(err)
	err = db.AutoMigrate(&category.Category{}, &product.Product{}, &wishlist.Wishlist{}, &cart.Cart{}, &transaction.Transaction{}, &transaction.TransactionDetail{}, &user.User{})
	utils.LogIfError(err)

	router := gin.Default()
	router.Use(gin.CustomRecovery(exception.Handler))

	userGroup := router.Group("api/v1/user")
	{
		userController := user.NewUserController(db)
		userGroup.POST("/register", userController.Register)
		userGroup.POST("/login", userController.Login)
		userGroup.POST("/token", userController.Token)
	}

	categoryRoute := router.Group("api/v1/category")
	{
		categoryController := category.NewCategoryController(db)
		categoryRoute.POST("/", user.Middleware, categoryController.Save)
		categoryRoute.GET("/:id", categoryController.FindById)
		categoryRoute.GET("/", categoryController.FindAll)
		categoryRoute.PUT("/:id", user.Middleware, categoryController.Update)
		categoryRoute.DELETE("/:id", user.Middleware, categoryController.Delete)
	}

	productRoute := router.Group("api/v1/product")
	{
		productController := product.NewProductController(db)
		productRoute.POST("/", user.Middleware, productController.Save)
		productRoute.GET("/:id", productController.FindById)
		productRoute.GET("/", productController.FindAll)
		productRoute.PUT("/:id", user.Middleware, productController.Update)
		productRoute.DELETE("/:id", user.Middleware, productController.Delete)
		productRoute.PUT("/top/:id", user.Middleware, productController.SetTopProduct)
		productRoute.PUT("/recommendation/:id", user.Middleware, productController.SetRecommendationProduct)
	}

	cartRoute := router.Group("api/v1/cart")
	{
		cartController := cart.NewController(db)
		cartRoute.POST("/:productId/:total", user.Middleware, cartController.Store)
		cartRoute.DELETE("/:id", user.Middleware, cartController.Destroy)
		cartRoute.GET("/", user.Middleware, cartController.Find)
	}

	wishlistRoute := router.Group("api/v1/wishlist")
	{
		wishlistController := wishlist.NewWishlistController(db)
		wishlistRoute.POST("/:productId", user.Middleware, wishlistController.Store)
		wishlistRoute.DELETE("/:id", user.Middleware, wishlistController.Destroy)
		wishlistRoute.GET("/", user.Middleware, wishlistController.Find)
	}

	err = router.Run()
	utils.LogIfError(err)
}
