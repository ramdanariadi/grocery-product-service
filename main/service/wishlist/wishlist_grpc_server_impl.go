package wishlist

import (
	"database/sql"
	"github.com/ramdanariadi/grocery-product-service/main/helpers"
	"github.com/ramdanariadi/grocery-product-service/main/models"
	"github.com/ramdanariadi/grocery-product-service/main/repositories/product"
	"github.com/ramdanariadi/grocery-product-service/main/repositories/transactions"
	"github.com/ramdanariadi/grocery-product-service/main/service/response"
	"github.com/ramdanariadi/grocery-product-service/main/utils"
	"golang.org/x/net/context"
	"log"
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
	tx, err := server.Repository.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)
	productModel := server.ProductRepository.FindById(ctx, tx, wishlist.ProductId)
	if utils.IsStructEmpty(productModel) {
		status, message := utils.ResponseForModifying(false)
		return &response.Response{Status: status, Message: message}, nil
	}
	wishlistModel := models.WishlistModel{
		ImageUrl:  productModel.ImageUrl,
		Name:      productModel.Name,
		Weight:    uint32(productModel.Weight),
		Category:  productModel.Category,
		Price:     productModel.Price,
		PerUnit:   uint64(productModel.PerUnit),
		UserId:    wishlist.UserId,
		ProductId: productModel.Id,
	}
	save := server.Repository.Save(ctx, tx, wishlistModel)
	status, message := utils.ResponseForModifying(save)
	return &response.Response{Status: status, Message: message}, nil
}

func (server WishlistServiceServerImpl) Delete(ctx context.Context, id *UserAndWishlistId) (*response.Response, error) {
	tx, err := server.Repository.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)
	deleted := server.Repository.Delete(ctx, tx, id.UserId, id.WishlistId)
	status, message := utils.ResponseForModifying(deleted)
	return &response.Response{
		Status:  status,
		Message: message,
	}, nil
}

func (server WishlistServiceServerImpl) FindByUserId(ctx context.Context, id *WishlistUserId) (*MultipleWishlistResponse, error) {
	tx, err := server.Repository.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)
	wishlistModels := server.Repository.FindByUserId(ctx, tx, id.Id)
	log.Println("wishlist user id " + id.Id)
	wishlist := fetchWishlist(wishlistModels)
	status, message := utils.ResponseForQuerying(len(wishlist) > 0)
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
		var imageUrl sql.NullString
		err := rows.Scan(&wishlist.Id, &wishlist.Name, &wishlist.Price, &wishlist.Weight, &wishlist.Category, &wishlist.PerUnit, &imageUrl)
		if err != nil {
			log.Println("scan error : " + err.Error())
			continue
		}

		if imageUrl.Valid {
			wishlist.ImageUrl = imageUrl.String
		}

		wishlists = append(wishlists, &wishlist)
	}
	rows.Close()
	return wishlists
}

func (server WishlistServiceServerImpl) mustEmbedUnimplementedWishlistServiceServer() {
	//TODO implement me
	panic("implement me")
}
