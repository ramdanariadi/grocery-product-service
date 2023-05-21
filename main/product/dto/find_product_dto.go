package dto

type FindProductRequest struct {
	Search           *string `form:"search"`
	IsTop            *bool   `form:"isTop"`
	IsRecommendation *bool   `form:"isRecommendation"`
	CategoryId       *string `form:"categoryId"`
	PageIndex        int     `form:"pageIndex"`
	PageSize         int     `form:"pageSize"`
}

type FindProductResponse struct {
	Data            []*ProductDTO `json:"data"`
	RecordsTotal    int64         `json:"recordsTotal"`
	RecordsFiltered int           `json:"recordsFiltered"`
}
