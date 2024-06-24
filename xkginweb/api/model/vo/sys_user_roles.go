package vo

type SysRolesVo struct {
	ID       uint   `json:"id"`
	RoleName string `json:"roleName"` // 角色名
	RoleCode string `json:"roleCode"` // 角色代号
}
