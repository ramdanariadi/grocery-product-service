package cart

import (
	"context"
	"database/sql"
	"github.com/ramdanariadi/grocery-product-service/main/cart/model"
	repository2 "github.com/ramdanariadi/grocery-product-service/main/cart/repository"
	"github.com/ramdanariadi/grocery-product-service/main/helpers"
	"github.com/ramdanariadi/grocery-product-service/main/product/repository"
	"github.com/ramdanariadi/grocery-product-service/main/response"
	"github.com/ramdanariadi/grocery-product-service/main/utils"
)

type CartServiceServerImpl struct {
	Repository        repository2.CartRepositoryImpl
	ProductRepository repository.ProductRepositoryImpl
}

func NewCartServiceImpl(db *sql.DB) *CartServiceServerImpl {
	return &CartServiceServerImpl{
		Repository:        repository2.CartRepositoryImpl{DB: db},
		ProductRepository: repository.ProductRepositoryImpl{DB: db},
	}
}

func (server CartServiceServerImpl) Save(ctx context.Context, cart *Cart) (*response.Response, error) {
	tx, err := server.ProductRepository.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)
	productModel := server.ProductRepository.FindById(ctx, tx, cart.ProductId)
	cartModel := model.CartModel{
		ImageUrl:  productModel.ImageUrl,
		ProductId: productModel.Id,
		Name:      productModel.Name,
		Weight:    uint32(productModel.Weight),
		Category:  productModel.Category,
		Price:     productModel.Price,
		PerUnit:   uint64(productModel.PerUnit),
		UserId:    cart.UserId,
		Total:     cart.Total,
	}
	err = server.Repository.Save(ctx, tx, &cartModel)
	status, message := utils.ResponseForModifying(err == nil)
	return &response.Response{
		Message: message,
		Status:  status,
	}, nil
}

func (server CartServiceServerImpl) Delete(ctx context.Context, id *CartAndUserId) (*response.Response, error) {
	tx, err := server.Repository.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)
	err = server.Repository.Delete(ctx, tx, id.UserId, id.Id)
	status, message := utils.ResponseForModifying(err == nil)
	return &response.Response{Message: message, Status: status}, nil
}

func (server CartServiceServerImpl) FindByUserId(ctx context.Context, id *CartUserId) (*MultipleCartResponse, error) {
	tx, err := server.Repository.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)
	rows := server.Repository.FindByUserId(ctx, tx, id.Id)
	wishlist := fetchWishlist(rows)
	status, message := utils.ResponseForQuerying(true)
	return &MultipleCartResponse{Status: status, Message: message, Data: wishlist}, nil
}

func fetchWishlist(rows *sql.Rows) []*CartDetail {
	var carts []*CartDetail
	for rows.Next() {
		cart := CartDetail{}
		rows.Scan(&cart.Id, &cart.Name, &cart.Price, &cart.Weight, &cart.Category, &cart.Total, &cart.PerUnit, &cart.ImageUrl)
		carts = append(carts, &cart)
	}
	helpers.LogIfError(rows.Close())
	return carts
}

func (server CartServiceServerImpl) mustEmbedUnimplementedCartServiceServer() {
	//TODO implement me
	panic("implement me")
}
