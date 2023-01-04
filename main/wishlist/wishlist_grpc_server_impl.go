package wishlist

import (
	"github.com/google/uuid"
	productModel "github.com/ramdanariadi/grocery-product-service/main/product/model"
	"github.com/ramdanariadi/grocery-product-service/main/response"
	"github.com/ramdanariadi/grocery-product-service/main/utils"
	"github.com/ramdanariadi/grocery-product-service/main/wishlist/model"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type WishlistServiceServerImpl struct {
	DB *gorm.DB
}

func NewWishlistServer(db *gorm.DB) *WishlistServiceServerImpl {
	return &WishlistServiceServerImpl{
		DB: db,
	}
}

func (server WishlistServiceServerImpl) Save(_ context.Context, wishlist *Wishlist) (*response.Response, error) {
	var productRef productModel.Product
	first := server.DB.First(&productRef, "id = ?", wishlist.ProductId)
	if first.RowsAffected == 0 {
		status, message := utils.QueryResponse(false)
		return &response.Response{
			Message: message,
			Status:  status,
		}, nil
	}

	var wishlistModel model.Wishlist
	tx := server.DB.First(&wishlistModel, "product_id = ? and user_id = ?", wishlist.ProductId, wishlist.UserId)

	if tx.RowsAffected > 0 {
		status, message := utils.ModifyingResponse(true)
		return &response.Response{Status: status, Message: message}, nil
	}

	id, _ := uuid.NewUUID()
	wishlistModel = model.Wishlist{
		ID:        id.String(),
		ImageUrl:  productRef.ImageUrl,
		Name:      productRef.Name,
		Weight:    uint32(productRef.Weight),
		Category:  productRef.Category.Category,
		Price:     productRef.Price,
		PerUnit:   uint64(productRef.PerUnit),
		UserId:    wishlist.UserId,
		ProductId: productRef.ID,
	}
	save := server.DB.Save(&wishlistModel)
	status, message := utils.ModifyingResponse(save.Error == nil)
	return &response.Response{Status: status, Message: message}, nil
}

func (server WishlistServiceServerImpl) Delete(_ context.Context, userAndProductId *UserAndProductId) (*response.Response, error) {
	tx := server.DB.Delete(&model.Wishlist{UserId: userAndProductId.UserId, ProductId: userAndProductId.ProductId})
	status, message := utils.ModifyingResponse(tx.Error == nil)
	return &response.Response{
		Status:  status,
		Message: message,
	}, nil
}

func (server WishlistServiceServerImpl) FindByUserId(_ context.Context, id *WishlistUserId) (*MultipleWishlistResponse, error) {
	var wishlistModels []*model.Wishlist
	tx := server.DB.Find(&wishlistModels, "user_id = ?", id.Id)
	utils.LogIfError(tx.Error)
	wishlist := fetchWishlist(wishlistModels)
	status, message := utils.QueryResponse(true)
	return &MultipleWishlistResponse{
		Status:  status,
		Message: message,
		Data:    wishlist,
	}, nil
}

func (server WishlistServiceServerImpl) FindWishlistByProductId(_ context.Context, id *UserAndProductId) (*WishlistResponse, error) {
	var wishlist model.Wishlist
	tx := server.DB.Find(&wishlist, "user_id = ? and product_id = ?", id.UserId, id.ProductId)
	status, message := utils.QueryResponse(tx.RowsAffected > 0)
	if tx.RowsAffected == 0 {
		return &WishlistResponse{
			Message: message,
			Status:  status,
			Data:    nil,
		}, nil
	}

	return &WishlistResponse{Message: message, Status: status, Data: &WishlistDetail{
		Id:        wishlist.ID,
		Name:      wishlist.Name,
		Weight:    wishlist.Weight,
		ProductId: wishlist.ProductId,
		Price:     wishlist.Price,
		Category:  wishlist.Category,
		ImageUrl:  wishlist.ImageUrl,
		UserId:    wishlist.UserId,
		PerUnit:   wishlist.PerUnit,
	}}, nil
}

func fetchWishlist(wishlistModels []*model.Wishlist) []*WishlistDetail {
	var wishlists []*WishlistDetail
	for _, w := range wishlistModels {
		wishlist := WishlistDetail{}
		wishlist.Id = w.ID
		wishlist.Name = w.Name
		wishlist.Price = w.Price
		wishlist.Weight = w.Weight
		wishlist.Category = w.Category
		wishlist.PerUnit = w.PerUnit
		wishlist.ImageUrl = w.ImageUrl
		wishlist.ProductId = w.ProductId
		wishlist.UserId = w.UserId
		wishlists = append(wishlists, &wishlist)
	}
	return wishlists
}

func (server WishlistServiceServerImpl) mustEmbedUnimplementedWishlistServiceServer() {
	//TODO implement me
	panic("implement me")
}
