package response

// 响应参数
type XkVideoResp struct {
	PageNum  int         `form:"pageNum" json:"pageNum"`
	PageSize int         `form:"pageSize" json:"pageSize"`
	Total    int64       `json:"total"`
	List     interface{} `json:"list"`
}
