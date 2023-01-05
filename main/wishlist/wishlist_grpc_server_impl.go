package wishlist

import (
	"database/sql"
	repository2 "github.com/ramdanariadi/grocery-product-service/main/product/repository"
	"github.com/ramdanariadi/grocery-product-service/main/response"
	"github.com/ramdanariadi/grocery-product-service/main/setup"
	"github.com/ramdanariadi/grocery-product-service/main/utils"
	"github.com/ramdanariadi/grocery-product-service/main/wishlist/model"
	"github.com/ramdanariadi/grocery-product-service/main/wishlist/repository"
	"golang.org/x/net/context"
	"log"
)

type WishlistServiceServerImpl struct {
	Repository        repository.WishlistRepository
	ProductRepository repository2.ProductRepository
	DB                *sql.DB
}

func NewWishlistServer(db *sql.DB) *WishlistServiceServerImpl {
	return &WishlistServiceServerImpl{
		Repository:        repository.WishlistRepositoryImpl{},
		ProductRepository: repository2.ProductRepositoryImpl{},
		DB:                db,
	}
}

func (server WishlistServiceServerImpl) Save(ctx context.Context, wishlist *Wishlist) (*response.Response, error) {
	tx, err := server.DB.Begin()
	utils.PanicIfError(err)
	defer utils.CommitOrRollback(tx)
	productModel := server.ProductRepository.FindById(ctx, tx, wishlist.ProductId)
	if productModel == nil {
		status, message := setup.ResponseForModifying(false)
		return &response.Response{Status: status, Message: message}, nil
	}

	check := server.Repository.FindByUserAndProductId(ctx, tx, wishlist.UserId, productModel.Id)
	if check != nil {
		status, message := setup.ResponseForModifying(true)
		return &response.Response{Status: status, Message: message}, nil
	}

	wishlistModel := model.WishlistModel{
		ImageUrl:  productModel.ImageUrl,
		Name:      productModel.Name,
		Weight:    uint32(productModel.Weight),
		Category:  productModel.Category,
		Price:     productModel.Price,
		PerUnit:   uint64(productModel.PerUnit),
		UserId:    wishlist.UserId,
		ProductId: productModel.Id,
	}
	err = server.Repository.Save(ctx, tx, &wishlistModel)
	status, message := setup.ResponseForModifying(err == nil)
	return &response.Response{Status: status, Message: message}, nil
}

func (server WishlistServiceServerImpl) Delete(ctx context.Context, id *UserAndProductId) (*response.Response, error) {
	tx, err := server.DB.Begin()
	utils.PanicIfError(err)
	defer utils.CommitOrRollback(tx)
	err = server.Repository.Delete(ctx, tx, id.UserId, id.ProductId)
	status, message := setup.ResponseForModifying(err == nil)
	return &response.Response{
		Status:  status,
		Message: message,
	}, nil
}

func (server WishlistServiceServerImpl) FindByUserId(ctx context.Context, id *WishlistUserId) (*MultipleWishlistResponse, error) {
	tx, err := server.DB.Begin()
	utils.PanicIfError(err)
	defer utils.CommitOrRollback(tx)
	wishlistModels := server.Repository.FindByUserId(ctx, tx, id.Id)
	wishlist := fetchWishlist(wishlistModels)
	status, message := setup.ResponseForQuerying(len(wishlist) > 0)
	return &MultipleWishlistResponse{
		Status:  status,
		Message: message,
		Data:    wishlist,
	}, nil
}
func (server WishlistServiceServerImpl) FindWishlistByProductId(ctx context.Context, id *UserAndProductId) (*WishlistResponse, error) {
	tx, err := server.DB.Begin()
	utils.PanicIfError(err)
	defer utils.CommitOrRollback(tx)

	wishlistModel := server.Repository.FindByUserAndProductId(ctx, tx, id.UserId, id.ProductId)
	status, message := setup.ResponseForQuerying(wishlistModel != nil)

	if wishlistModel == nil {
		return &WishlistResponse{Message: message, Status: status, Data: nil}, nil
	}

	return &WishlistResponse{Message: message, Status: status, Data: &WishlistDetail{
		Id:        wishlistModel.Id,
		Name:      wishlistModel.Name,
		Weight:    wishlistModel.Weight,
		ProductId: wishlistModel.ProductId,
		Price:     wishlistModel.Price,
		Category:  wishlistModel.Category,
		ImageUrl:  wishlistModel.ImageUrl,
		UserId:    wishlistModel.UserId,
		PerUnit:   wishlistModel.PerUnit,
	}}, nil
}

func fetchWishlist(rows *sql.Rows) []*WishlistDetail {
	var wishlists []*WishlistDetail
	for rows.Next() {
		wishlist := WishlistDetail{}
		var imageUrl sql.NullString
		err := rows.Scan(&wishlist.Id, &wishlist.Name, &wishlist.Price, &wishlist.Weight, &wishlist.Category, &wishlist.PerUnit, &imageUrl, &wishlist.ProductId, &wishlist.UserId)
		if err != nil {
			log.Println("scan error : " + err.Error())
			continue
		}

		if imageUrl.Valid {
			wishlist.ImageUrl = imageUrl.String
		}

		wishlists = append(wishlists, &wishlist)
	}
	utils.LogIfError(rows.Close())
	return wishlists
}

func (server WishlistServiceServerImpl) mustEmbedUnimplementedWishlistServiceServer() {
	//TODO implement me
	panic("implement me")
}
