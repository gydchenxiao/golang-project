package response

// 查询作者业务的数据信息---list
type AuthorResopnse struct {
	Id         int    `json:"id"`
	Nickname   string `json:"nickname"`
	Account    string `json:"account"`
	Avatar     string `json:"avatar"`
	AuthorName string `json:"author_name"`
}

// 响应参数----------------统一返回
type XkAuthorResp struct {
	PageNum  int         `form:"pageNum" json:"pageNum"`
	PageSize int         `form:"pageSize" json:"pageSize"`
	Total    int64       `json:"total"`
	List     interface{} `json:"list"`
}
