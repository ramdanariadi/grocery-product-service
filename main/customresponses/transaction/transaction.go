package transaction

type TransactionCustomResponse struct {
	Id                string               `json:"id"`
	TotalPrice        int                  `json:"totalPrice"`
	TransactionDate   string               `json:"transactionDate"`
	UserName          string               `json:"userName"`
	UserMobile        interface{}          `json:"userMobile"`
	UserEmail         interface{}          `json:"userEmail"`
	DetailTransaction []ProductTransaction `json:"detailTransaction"`
}

type ProductTransaction struct {
	Name          string      `json:"name"`
	Id            string      `json:"id"`
	ImageUrl      interface{} `json:"imageUrl"`
	ProductId     string      `json:"productId"`
	Price         int         `json:"price"`
	Weight        int         `json:"weight"`
	PerUnit       int         `json:"perUnit"`
	Total         int         `json:"total"`
	TransactionId string      `json:"transactionId"`
}
