package vo

type SysApisVo struct {
	ID     uint   `json:"id"`
	Path   string `json:"path"`   // 路径
	Title  string `json:"title"`  // 标题
	MenuId string `json:"menuId"` // 菜单id
	Method string `json:"method"` // 请求方式
}
