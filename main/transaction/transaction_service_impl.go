package transaction

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/ramdanariadi/grocery-product-service/main/product"
	"github.com/ramdanariadi/grocery-product-service/main/transaction/dto"
	"github.com/ramdanariadi/grocery-product-service/main/transaction/model"
	"github.com/ramdanariadi/grocery-product-service/main/utils"
	"gorm.io/gorm"
	"log"
)

type TransactionServiceImpl struct {
	DB *gorm.DB
}

func (service TransactionServiceImpl) save(request *dto.AddTransactionDTO, userId string) {
	marshal, _ := json.Marshal(request)
	log.Println("request body " + string(marshal))
	err := service.DB.Transaction(func(tx *gorm.DB) error {
		var productIds []string
		for _, item := range request.Data {
			productIds = append(productIds, item.ProductId)
		}

		var products []*product.Product
		tx.Model(&product.Product{}).Where("id IN ?", productIds).Preload("Category").Find(&products)

		//if len(products) != len(request.Data) {
		//	panic(exception.ValidationException{Message: "INVALID_PRODUCT"})
		//}

		productMap := map[string]*product.Product{}
		var totalPrice uint64
		for _, p := range products {
			totalPrice += p.Price
			productMap[p.ID] = p
		}

		id, _ := uuid.NewUUID()
		transaction := model.Transaction{
			ID:         id.String(),
			UserId:     userId,
			TotalPrice: totalPrice,
		}

		if err := tx.Create(&transaction).Error; err != nil {
			return err
		}
		var transactionDetails []*model.TransactionDetail
		for _, d := range request.Data {
			p := productMap[d.ProductId]
			dtId, _ := uuid.NewUUID()
			detail := model.TransactionDetail{ID: dtId.String(), Transaction: transaction, Product: *p, Total: d.Total, Name: p.Name, Price: p.Price, ImageUrl: p.ImageUrl, Description: p.Description, PerUnit: p.PerUnit, Weight: p.Weight, CategoryId: p.CategoryId, Category: p.Category.Category}
			transactionDetails = append(transactionDetails, &detail)
		}
		log.Printf("request body data length %d", len(transactionDetails))
		if err := tx.Create(&transactionDetails).Error; err != nil {
			return err
		}
		log.Println("success save detail transaction")
		return nil
	})
	utils.PanicIfError(err)
}

func (service TransactionServiceImpl) find(param *dto.FindTransactionDTO) []*dto.TransactionDTO {
	var transactions []*model.Transaction
	tx := service.DB.Model(&model.Transaction{})
	tx.Joins("LEFT JOIN transaction_details td ON td.transaction_id = transactions.id")
	tx.Joins("LEFT JOIN products p ON td.product_id = p.id AND p.deleted_at IS NULL")
	tx.Preload("TransactionDetails.Product")
	if param.Search != nil {
		tx.Where("LOWER(p.name) ilike ?", "%"+*param.Search+"%")
	}
	tx.Limit(param.PageSize).Offset(param.PageIndex * param.PageSize).Find(&transactions)

	result := make([]*dto.TransactionDTO, 0)
	for _, t := range transactions {
		transactionDTO := dto.TransactionDTO{Id: t.ID, PriceTotal: 0}
		for _, td := range t.TransactionDetails {
			p := td.Product
			item := dto.TransactionItemDTO{ID: td.ID, Name: td.Name, Price: td.Price, PerUnit: td.PerUnit, Weight: td.Weight, ImageUrl: p.ImageUrl, Description: p.Description, Total: td.Total}
			transactionDTO.Items = append(transactionDTO.Items, &item)
			transactionDTO.PriceTotal += td.Price
		}
		result = append(result, &transactionDTO)
	}

	return result
}
