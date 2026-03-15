package model

type PageResult struct {
	List     interface{} `json:"list"`
	Total    int64       `json:"total"`
	PageNum  int         `json:"pageNum"`
	PageSize int         `json:"pageSize"`
}

func NewPageResult(list interface{}, total int64, pageNum, pageSize int) *PageResult {
	return &PageResult{
		List:     list,
		Total:    total,
		PageNum:  pageNum,
		PageSize: pageSize,
	}
}
