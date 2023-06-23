package model

type QueryPage struct {
	PageNo    int    `json:"pageNo" form:"pageNo"`       // 当前页码
	PageSize  int    `json:"pageSize" form:"pageSize"`   // 每页数量
	Sort      string `json:"sort" form:"sort"`           // 排序字段
	Direction string `json:"direction" form:"direction"` // 排序类型：asc、desc
}

func NewQueryPage(pageNo int, pageSize int) *QueryPage {
	return &QueryPage{PageNo: pageNo, PageSize: pageSize}
}

func (a QueryPage) GetPageNo() int {
	if a.PageNo <= 0 {
		return 1
	}
	return a.PageNo
}

func (a QueryPage) GetPageSize() int {
	if a.PageSize <= 0 {
		return 10
	}
	return a.PageSize
}
