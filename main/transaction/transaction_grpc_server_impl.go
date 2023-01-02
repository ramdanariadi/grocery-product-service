package transaction

import (
	"context"
	"errors"
	"github.com/google/uuid"
	productModel "github.com/ramdanariadi/grocery-product-service/main/product/model"
	"github.com/ramdanariadi/grocery-product-service/main/response"
	"github.com/ramdanariadi/grocery-product-service/main/transaction/model"
	"github.com/ramdanariadi/grocery-product-service/main/utils"
	"gorm.io/gorm"
)

type TransactionServiceServerImpl struct {
	DB *gorm.DB
}

func NewTransactionServiceServer(db *gorm.DB) *TransactionServiceServerImpl {
	return &TransactionServiceServerImpl{
		DB: db,
	}
}

func (server TransactionServiceServerImpl) FindByTransactionId(_ context.Context, id *TransactionId) (*TransactionResponse, error) {
	var transaction model.Transaction
	tx := server.DB.Find(&transaction, "id = ?", id.Id)
	if tx.Error != nil {
		status, message := utils.QueryResponse(false)
		return &TransactionResponse{Status: status, Message: message, Data: nil}, nil
	}

	transactionData := Transaction{
		Id:         transaction.ID,
		TotalPrice: transaction.TotalPrice,
	}
	attachTransactionDetail(&transactionData, transaction.TransactionDetails)
	status, message := utils.QueryResponse(true)
	return &TransactionResponse{Status: status, Message: message, Data: &transactionData}, nil
}

func attachTransactionDetail(transaction *Transaction, detailTransaction []*model.TransactionDetail) {
	var transactionProductDetail []*TransactionProductDetail
	for _, dt := range detailTransaction {
		productDetail := TransactionProductDetail{
			Id:          dt.ID,
			Name:        dt.Name,
			ImageUrl:    dt.ImageUrl,
			ProductId:   dt.ProductId,
			Weight:      uint32(dt.Weight),
			Price:       dt.Price,
			PerUnit:     uint64(dt.PerUnit),
			Description: dt.Description,
			Total:       uint32(dt.Total),
		}
		transactionProductDetail = append(transactionProductDetail, &productDetail)
	}

	transaction.Products = transactionProductDetail
}

func (server TransactionServiceServerImpl) FindByUserId(_ context.Context, transactionUserId *TransactionUserId) (*MultipleTransactionResponse, error) {
	var transactionModels []*model.Transaction
	server.DB.Find(&transactionModels, "user_id = ?", transactionUserId.Id)
	status, message := utils.QueryResponse(true)
	result := MultipleTransactionResponse{
		Status:  status,
		Message: message,
	}

	for _, t := range transactionModels {
		tTemp := Transaction{
			Id:         t.ID,
			TotalPrice: t.TotalPrice,
		}
		attachTransactionDetail(&tTemp, t.TransactionDetails)
		result.Data = append(result.Data, &tTemp)
	}
	return &result, nil
}

func (server TransactionServiceServerImpl) Save(ctx context.Context, body *TransactionBody) (*response.Response, error) {
	var ids []string
	for _, d := range body.Products {
		ids = append(ids, d.ProductId)
	}

	var productModels []*productModel.Product
	server.DB.Find(&productModels, "id in ?", ids)

	transactionId, _ := uuid.NewUUID()
	transactionModel := model.Transaction{ID: transactionId.String(), UserId: body.UserId}
	var detailTransaction []*model.TransactionDetail
	var totalPrice uint64
	for _, pm := range productModels {
		total, err := findProductTotal(body.Products, pm.ID)
		if err != nil {
			continue
		}
		totalPrice += uint64(pm.Weight/pm.PerUnit) * pm.Price * uint64(total)
		detailTransactionProductModel := model.TransactionDetail{
			ProductId: pm.ID, Name: pm.Name, ImageUrl: pm.ImageUrl, Price: pm.Price,
			PerUnit: pm.PerUnit, Weight: pm.Weight, Category: pm.Category, Description: pm.Description,
			Total: uint(total), TransactionId: transactionId.String(),
		}
		detailTransaction = append(detailTransaction, &detailTransactionProductModel)
	}
	transactionModel.TransactionDetails = detailTransaction
	transactionModel.TotalPrice = totalPrice
	tx := server.DB.Save(&transactionModel)
	utils.LogIfError(tx.Error)

	status, message := utils.QueryResponse(tx.Error == nil)
	return &response.Response{Status: status, Message: message}, nil
}

func findProductTotal(products []*TransactionProduct, id string) (uint32, error) {
	for _, tp := range products {
		if tp.ProductId == id {
			return tp.Total, nil
		}
	}
	return 0, errors.New("PRODUCT_NOT_FOUND")
}

func (server TransactionServiceServerImpl) Delete(ctx context.Context, id *TransactionId) (*response.Response, error) {
	err := server.DB.Transaction(func(tx *gorm.DB) error {

		if error := tx.Delete(&model.Transaction{}, "id = ? ", id.Id).Error; error != nil {
			return error
		}

		if error := tx.Delete(&model.TransactionDetail{}, "transaction_id = ?", id.Id).Error; error != nil {
			return error
		}
		return nil
	})
	status, message := utils.ModifyingResponse(err == nil)
	return &response.Response{Status: status, Message: message}, nil
}

func (server TransactionServiceServerImpl) mustEmbedUnimplementedTransactionServiceServer() {
	//TODO implement me
	panic("implement me")
}
