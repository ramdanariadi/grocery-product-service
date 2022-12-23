package transaction

import (
	"context"
	"database/sql"
	"errors"
	"github.com/ramdanariadi/grocery-product-service/main/helpers"
	repository2 "github.com/ramdanariadi/grocery-product-service/main/product/repository"
	"github.com/ramdanariadi/grocery-product-service/main/response"
	"github.com/ramdanariadi/grocery-product-service/main/transaction/model"
	"github.com/ramdanariadi/grocery-product-service/main/transaction/repository"
	"github.com/ramdanariadi/grocery-product-service/main/utils"
)

type TransactionServiceServerImpl struct {
	Repository        repository.TransactionRepositoryImpl
	ProductRepository repository2.ProductRepositoryImpl
}

func NewTransactionServiceServer(db *sql.DB) *TransactionServiceServerImpl {
	return &TransactionServiceServerImpl{
		Repository:        repository.TransactionRepositoryImpl{DB: db},
		ProductRepository: repository2.ProductRepositoryImpl{DB: db},
	}
}

func (transaction TransactionServiceServerImpl) FindByTransactionId(ctx context.Context, id *TransactionId) (*TransactionResponse, error) {
	tx, err := transaction.Repository.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)
	transactionModel := transaction.Repository.FindByTransactionId(ctx, tx, id.Id)
	var transactionData Transaction
	if transactionModel != nil {
		transactionData = Transaction{
			Id:         transactionModel.Id,
			TotalPrice: transactionModel.TotalPrice,
		}
		attachTransactionDetail(&transactionData, transactionModel.DetailTransaction)
	}

	status, message := utils.ResponseForQuerying(transactionModel != nil)
	return &TransactionResponse{Status: status, Message: message, Data: &transactionData}, nil
}

func attachTransactionDetail(transaction *Transaction, detailTransaction []*model.DetailTransactionProductModel) {
	var transactionProductDetail []*TransactionProductDetail
	for _, dt := range detailTransaction {
		productDetail := TransactionProductDetail{
			Id:          dt.Id,
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

func (transaction TransactionServiceServerImpl) FindByUserId(ctx context.Context, id *TransactionUserId) (*MultipleTransactionResponse, error) {
	tx, err := transaction.Repository.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)
	transactionModels := transaction.Repository.FindByUserId(ctx, tx, id.Id)
	status, message := utils.ResponseForQuerying(true)
	result := MultipleTransactionResponse{
		Status:  status,
		Message: message,
	}

	for _, t := range transactionModels {
		tTemp := Transaction{
			Id:         t.Id,
			TotalPrice: t.TotalPrice,
		}
		attachTransactionDetail(&tTemp, t.DetailTransaction)
		result.Data = append(result.Data, &tTemp)
	}
	return &result, nil
}

func (transaction TransactionServiceServerImpl) Save(ctx context.Context, body *TransactionBody) (*response.Response, error) {
	tx, err := transaction.Repository.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)
	var ids []string
	for _, d := range body.Products {
		ids = append(ids, d.ProductId)
	}

	productModels := transaction.ProductRepository.FindByIds(ctx, tx, ids)
	transactionModel := model.TransactionModel{UserId: body.UserId}
	var detailTransaction []*model.DetailTransactionProductModel
	var totalPrice uint64
	for _, pm := range productModels {
		total, err := findProductTotal(body.Products, pm.Id)
		if err != nil {
			continue
		}
		totalPrice += uint64(pm.Weight/pm.PerUnit) * pm.Price * uint64(total)
		detailTransactionProductModel := model.DetailTransactionProductModel{ProductId: pm.Id, Name: pm.Name, ImageUrl: pm.ImageUrl, Price: pm.Price,
			PerUnit: pm.PerUnit, Weight: pm.Weight, Category: pm.Category, Description: pm.Description, Total: uint(total)}
		detailTransaction = append(detailTransaction, &detailTransactionProductModel)
	}
	transactionModel.DetailTransaction = detailTransaction
	transactionModel.TotalPrice = totalPrice
	err = transaction.Repository.Save(ctx, tx, &transactionModel)
	status, message := utils.ResponseForQuerying(err == nil)
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

func (transaction TransactionServiceServerImpl) Delete(ctx context.Context, id *TransactionId) (*response.Response, error) {
	tx, err := transaction.Repository.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)
	err = transaction.Repository.Delete(ctx, tx, id.Id)
	status, message := utils.ResponseForModifying(err == nil)
	return &response.Response{Status: status, Message: message}, nil
}

func (transaction TransactionServiceServerImpl) mustEmbedUnimplementedTransactionServiceServer() {
	//TODO implement me
	panic("implement me")
}
