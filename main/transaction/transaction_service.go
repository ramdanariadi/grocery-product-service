package transaction

import "github.com/ramdanariadi/grocery-product-service/main/transaction/dto"

type Service interface {
	save(request *dto.AddTransactionDTO, userId string)
	find(param *dto.FindTransactionDTO) []*dto.TransactionDTO
}
