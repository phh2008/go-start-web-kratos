package model

type PageData[T any] struct {
	Count    int64 `json:"count"`    // 总记录数
	PageNo   int   `json:"pageNo"`   // 当前页码
	PageSize int   `json:"pageSize"` // 每页数量
	Data     []T   `json:"data"`     // 分页数据
}

func NewPageData[T any](pageNo int, pageSize int) *PageData[T] {
	return &PageData[T]{PageNo: pageNo, PageSize: pageSize}
}

func (a *PageData[T]) SetData(data []T) *PageData[T] {
	a.Data = data
	return a
}
