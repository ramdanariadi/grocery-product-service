package dto

type FindCartDTO struct {
	Search    *string `form:"search"`
	PageIndex int     `form:"pageIndex"`
	PageSize  int     `form:"pageSize"`
}
