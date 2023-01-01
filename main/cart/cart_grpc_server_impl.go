package cart

import (
	"context"
	"database/sql"
	"github.com/ramdanariadi/grocery-product-service/main/cart/model"
	repository2 "github.com/ramdanariadi/grocery-product-service/main/cart/repository"
	"github.com/ramdanariadi/grocery-product-service/main/product/repository"
	"github.com/ramdanariadi/grocery-product-service/main/response"
	"github.com/ramdanariadi/grocery-product-service/main/setup"
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
	utils.PanicIfError(err)
	defer utils.CommitOrRollback(tx)
	productModel := server.ProductRepository.FindById(ctx, tx, cart.ProductId)
	if productModel == nil {
		status, message := setup.ResponseForQuerying(false)
		return &response.Response{
			Message: message,
			Status:  status,
		}, nil
	}

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

	existingCartRow := server.Repository.FindByUserAndProductId(ctx, tx, cart.UserId, cart.ProductId)

	existingCart := CartDetail{}
	err = existingCartRow.Scan(&existingCart.Id, &existingCart.Name, &existingCart.Price, &existingCart.Weight,
		&existingCart.Category, &existingCart.Total, &existingCart.PerUnit, &existingCart.ImageUrl, &existingCart.ProductId)
	utils.LogIfError(err)

	if err == nil {
		cartModel.Id = existingCart.Id
		cartModel.Total = cart.Total + existingCart.Total
		err = server.Repository.Update(ctx, tx, &cartModel)
		status, message := setup.ResponseForModifying(err == nil)
		return &response.Response{
			Message: message,
			Status:  status,
		}, nil
	}

	err = server.Repository.Save(ctx, tx, &cartModel)
	status, message := setup.ResponseForModifying(err == nil)
	return &response.Response{
		Message: message,
		Status:  status,
	}, nil
}

func (server CartServiceServerImpl) Delete(ctx context.Context, id *CartAndUserId) (*response.Response, error) {
	tx, err := server.Repository.DB.Begin()
	utils.PanicIfError(err)
	defer utils.CommitOrRollback(tx)
	err = server.Repository.Delete(ctx, tx, id.UserId, id.Id)
	status, message := setup.ResponseForModifying(err == nil)
	return &response.Response{Message: message, Status: status}, nil
}

func (server CartServiceServerImpl) FindByUserId(ctx context.Context, id *CartUserId) (*MultipleCartResponse, error) {
	tx, err := server.Repository.DB.Begin()
	utils.PanicIfError(err)
	defer utils.CommitOrRollback(tx)
	rows := server.Repository.FindByUserId(ctx, tx, id.Id)
	wishlist := fetchWishlist(rows)
	status, message := setup.ResponseForQuerying(true)
	return &MultipleCartResponse{Status: status, Message: message, Data: wishlist}, nil
}

func fetchWishlist(rows *sql.Rows) []*CartDetail {
	var carts []*CartDetail
	for rows.Next() {
		cart := CartDetail{}
		err := rows.Scan(&cart.Id, &cart.Name, &cart.Price, &cart.Weight, &cart.Category, &cart.Total, &cart.PerUnit, &cart.ImageUrl, &cart.ProductId)
		if err != nil {
			continue
		}
		carts = append(carts, &cart)
	}
	utils.LogIfError(rows.Close())
	return carts
}

func (server CartServiceServerImpl) mustEmbedUnimplementedCartServiceServer() {
	//TODO implement me
	panic("implement me")
}
