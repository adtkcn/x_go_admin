package model

type BaseQuery struct {
	Key      string `form:"key"`
	PageNum  int    `form:"pageNum"`
	PageSize int    `form:"pageSize"`
}
