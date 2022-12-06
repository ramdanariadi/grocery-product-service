package wishlist

import (
	"database/sql"
	"github.com/ramdanariadi/grocery-product-service/main/models"
	"github.com/ramdanariadi/grocery-product-service/main/repositories/product"
	"github.com/ramdanariadi/grocery-product-service/main/repositories/transactions"
	"github.com/ramdanariadi/grocery-product-service/main/service/response"
	"github.com/ramdanariadi/grocery-product-service/main/utils"
	"golang.org/x/net/context"
)

type WishlistServiceServerImpl struct {
	Repository        transactions.WishlistRepositoryImpl
	ProductRepository product.ProductRepositoryImpl
}

func NewWishlistServer(db *sql.DB) *WishlistServiceServerImpl {
	return &WishlistServiceServerImpl{
		Repository:        transactions.WishlistRepositoryImpl{DB: db},
		ProductRepository: product.ProductRepositoryImpl{DB: db},
	}
}

func (server WishlistServiceServerImpl) Save(ctx context.Context, wishlist *Wishlist) (*response.Response, error) {
	tx, _ := server.Repository.DB.Begin()
	productModel := server.ProductRepository.FindById(ctx, tx, wishlist.Id)
	wishlistModel := models.WishlistModel{
		ImageUrl:  productModel.ImageUrl,
		Name:      productModel.Name,
		Weight:    productModel.Weight,
		Category:  productModel.Category,
		Price:     productModel.Price,
		PerUnit:   productModel.PerUnit,
		UserId:    wishlist.UserId,
		ProductId: productModel.Id,
	}
	save := server.Repository.Save(ctx, tx, wishlistModel)
	status, message := utils.FetchResponseForModifying(save)
	return &response.Response{Status: status, Message: message}, nil
}

func (server WishlistServiceServerImpl) Delete(ctx context.Context, id *UserAndWishlistId) (*response.Response, error) {
	tx, _ := server.Repository.DB.Begin()
	deleted := server.Repository.Delete(ctx, tx, id.UserId, id.WishlistId)
	status, message := utils.FetchResponseForModifying(deleted)
	return &response.Response{
		Status:  status,
		Message: message,
	}, nil
}

func (server WishlistServiceServerImpl) FindByUserId(ctx context.Context, id *WishlistUserId) (*MultipleWishlistResponse, error) {
	tx, _ := server.Repository.DB.Begin()
	wishlistModels := server.Repository.FindByUserId(ctx, tx, id.Id)
	wishlist := fetchWishlist(wishlistModels)
	status, message := utils.FetchResponseForQuerying(len(wishlist) > 0)
	return &MultipleWishlistResponse{
		Status:  status,
		Message: message,
		Data:    wishlist,
	}, nil
}

func fetchWishlist(rows *sql.Rows) []*WishlistDetail {
	var wishlists []*WishlistDetail
	for rows.Next() {
		wishlist := WishlistDetail{}
		rows.Scan(&wishlist.Id, &wishlist.Name, &wishlist.Price, &wishlist.Category, &wishlist.PerUnit, &wishlist.ImageUrl)
		wishlists = append(wishlists, &wishlist)
	}
	return wishlists
}

func (server WishlistServiceServerImpl) mustEmbedUnimplementedWishlistServiceServer() {
	//TODO implement me
	panic("implement me")
}
