package dto

type AllCategories struct {
	Data            []*CategoryDTO `json:"data"`
	RecordsTotal    int64          `json:"recordsTotal"`
	RecordsFiltered int            `json:"recordsFiltered"`
}

type PaginationDTO struct {
	PageIndex int `form:"pageIndex"`
	PageSize  int `form:"pageSize"`
}
