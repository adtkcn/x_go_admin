package model

type BaseQuery struct {
	Key      string `form:"key"`
	Page     int    `form:"page"`
	PageSize int    `form:"pageSize"`
}
