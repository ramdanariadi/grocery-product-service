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
	tx, _ := transaction.Repository.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	transactionRow, detailTransactionRows := transaction.Repository.FindByTransactionId(ctx, tx, id.Id)
	transactionData := Transaction{}
	transactionRow.Scan(transactionData.Id, transactionData.TotalPrice, transactionData.TransactionDate)
	var transactionProductDetail []*TransactionProductDetail
	transactionData.Products = transactionProductDetail
	for detailTransactionRows.Next() {
		productDetail := TransactionProductDetail{}
		detailTransactionRows.Scan(&productDetail.Name, &productDetail.Id, &productDetail.ImageUrl,
			&productDetail.ProductId, &productDetail.Price, &productDetail.Weight, &productDetail.PerUnit, &productDetail.Total)
		transactionProductDetail = append(transactionProductDetail, &productDetail)
	}
	detailTransactionRows.Close()
	return &transactionData, nil
}

func (transaction TransactionServiceServerImpl) FindByUserId(ctx context.Context, id *TransactionUserId) (*Transactions, error) {
	tx, _ := transaction.Repository.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	transactionRows, detailTransactionRows := transaction.Repository.FindByUserId(ctx, tx, id.Id)
	var transactionData Transactions
	for transactionRows.Next() {
		t := Transaction{}
		transactionRows.Scan(t.Id, t.TotalPrice, t.TransactionDate)
		transactionData.Transactions = append(transactionData.Transactions, &t)
	}
	transactionRows.Close()
	for detailTransactionRows.Next() {
		productDetail := TransactionProductDetail{}
		var transactionId string
		detailTransactionRows.Scan(&productDetail.Name, &productDetail.Id, &productDetail.ImageUrl,
			&productDetail.ProductId, &productDetail.Price, &productDetail.Weight, &productDetail.PerUnit, &productDetail.Total, &transactionId)
		attachDetailProductToTransaction(&transactionData, &productDetail, transactionId)
	}
	detailTransactionRows.Close()
	return &transactionData, nil
}

func attachDetailProductToTransaction(transactions *Transactions, detail *TransactionProductDetail, transactionId string) {
	for _, dt := range transactions.Transactions {
		if dt.Id == transactionId {
			dt.Products = append(dt.Products, detail)
		}
	}
}

func (transaction TransactionServiceServerImpl) Save(ctx context.Context, body *TransactionBody) (*response.Response, error) {
	tx, _ := transaction.Repository.DB.Begin()
	defer helpers.CommitOrRollback(tx)
	helpers.CommitOrRollback(tx)
	var ids []string
	for _, d := range body.Products {
		ids = append(ids, d.ProductId)
	}

	productModels := transaction.ProductRepository.FindByIds(ctx, tx, ids)
	transactionModel := models.TransactionModel{UserId: body.UserId}
	var detailTransaction []models.DetailTransactionProductModel
	var totalPrice uint64
	for _, pm := range productModels {
		total, err := findProductTotal(body.Products, pm.Id)
		if err != nil {
			continue
		}
		totalPrice += pm.Price * uint64(total)
		detailTransactionProductModel := models.DetailTransactionProductModel{ProductId: pm.Id, Name: pm.Name, ImageUrl: pm.ImageUrl, Price: pm.Price,
			PerUnit: pm.PerUnit, Weight: pm.Weight, Category: pm.Category, Description: pm.Description, Total: uint(total)}
		detailTransaction = append(detailTransaction, detailTransactionProductModel)
	}
	transactionModel.DetailTransaction = detailTransaction
	transactionModel.TotalPrice = totalPrice
	transactionModel.UserId = ""
	transaction.Repository.Save(ctx, tx, transactionModel)
	status, message := utils.FetchResponseForQuerying(true)
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
	//TODO implement me
	panic("implement me")
}

func (transaction TransactionServiceServerImpl) mustEmbedUnimplementedTransactionServiceServer() {
	//TODO implement me
	panic("implement me")
}
