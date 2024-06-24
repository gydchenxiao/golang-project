package request

// 作者查询参数
type AuthorPageReq struct {
	// 分页信息
	PageNum  int `form:"pageNum" json:"pageNum"`
	PageSize int `form:"pageSize" json:"pageSize"`
	// 搜索关键词---nickname,account,author_name
	Keyword string `from:"keyword" json:"keyword"`
	// 根据用户id查询
	UserId int `from:"userId" json:"userId"`
}
