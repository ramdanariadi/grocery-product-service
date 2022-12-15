package transaction

import (
	"context"
	"database/sql"
	"errors"
	"github.com/ramdanariadi/grocery-product-service/main/helpers"
	"github.com/ramdanariadi/grocery-product-service/main/models"
	"github.com/ramdanariadi/grocery-product-service/main/repositories/product"
	"github.com/ramdanariadi/grocery-product-service/main/repositories/transactions"
	"github.com/ramdanariadi/grocery-product-service/main/service/response"
	"github.com/ramdanariadi/grocery-product-service/main/utils"
)

type TransactionServiceServerImpl struct {
	Repository        transactions.TransactionRepositoryImpl
	ProductRepository product.ProductRepositoryImpl
}

func NewTransactionServiceServer(db *sql.DB) *TransactionServiceServerImpl {
	return &TransactionServiceServerImpl{
		Repository:        transactions.TransactionRepositoryImpl{DB: db},
		ProductRepository: product.ProductRepositoryImpl{DB: db},
	}
}

func (transaction TransactionServiceServerImpl) FindByTransactionId(ctx context.Context, id *TransactionId) (*Transaction, error) {
	tx, err := transaction.Repository.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)
	transactionModel := transaction.Repository.FindByTransactionId(ctx, tx, id.Id)
	transactionData := Transaction{
		Id:         transactionModel.Id,
		TotalPrice: transactionModel.TotalPrice,
	}
	attachTransactionDetail(&transactionData, transactionModel.DetailTransaction)
	return &transactionData, nil
}

func attachTransactionDetail(transaction *Transaction, detailTransaction []*models.DetailTransactionProductModel) {
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

func (transaction TransactionServiceServerImpl) FindByUserId(ctx context.Context, id *TransactionUserId) (*Transactions, error) {
	tx, err := transaction.Repository.DB.Begin()
	helpers.PanicIfError(err)
	defer helpers.CommitOrRollback(tx)
	transactionModels := transaction.Repository.FindByUserId(ctx, tx, id.Id)
	var transactionData Transactions

	for _, t := range transactionModels {
		tTemp := Transaction{
			Id:         t.Id,
			TotalPrice: t.TotalPrice,
		}
		attachTransactionDetail(&tTemp, t.DetailTransaction)
		transactionData.Transactions = append(transactionData.Transactions, &tTemp)
	}

	return &transactionData, nil
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
	transactionModel := models.TransactionModel{UserId: body.UserId}
	var detailTransaction []*models.DetailTransactionProductModel
	var totalPrice uint64
	for _, pm := range productModels {
		total, err := findProductTotal(body.Products, pm.Id)
		if err != nil {
			continue
		}
		totalPrice += uint64(pm.Weight/pm.PerUnit) * pm.Price * uint64(total)
		detailTransactionProductModel := models.DetailTransactionProductModel{ProductId: pm.Id, Name: pm.Name, ImageUrl: pm.ImageUrl, Price: pm.Price,
			PerUnit: pm.PerUnit, Weight: pm.Weight, Category: pm.Category, Description: pm.Description, Total: uint(total)}
		detailTransaction = append(detailTransaction, &detailTransactionProductModel)
	}
	transactionModel.DetailTransaction = detailTransaction
	transactionModel.TotalPrice = totalPrice
	transaction.Repository.Save(ctx, tx, transactionModel)
	status, message := utils.ResponseForQuerying(true)
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
	transaction.Repository.Delete(ctx, tx, id.Id)
	status, message := utils.ResponseForModifying(true)
	return &response.Response{Status: status, Message: message}, nil
}

func (transaction TransactionServiceServerImpl) mustEmbedUnimplementedTransactionServiceServer() {
	//TODO implement me
	panic("implement me")
}
