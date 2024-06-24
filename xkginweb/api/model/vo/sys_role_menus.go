package vo

type SysMenusVo struct {
	ID       uint   `json:"id"`
	Path     string `json:"path"`     // 路径
	Icon     string `json:"icon"`     // 图标
	Name     string `json:"name"`     // 名字
	Title    string `json:"title"`    // 标题
	ParentId uint   `json:"parentId"` // 标题
}
