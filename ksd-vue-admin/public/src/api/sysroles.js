import request from '@/request/index.js'

/**
 * 查询用户角色信息
 */
export const FindData = () => {
  return request.post(`/sys/role/list`, {})
}

/**
 * 查询系统角色列表并分页
 */
export const LoadData = (data) => {
  return request.post(`/sys/role/load`, data)
}

/**
 * 根据id查询系统角色信息
 */
export const GetById = (id) => {
  return request.post(`/sys/role/get/${id}`)
}

/**
 * 保存系统角色
 */
export const SaveData = (data) => {
  return request.post(`/sys/role/save`, data)
}

/**
 * 更新系统角色信息
 */
export const UpdateData = (data) => {
  return request.post(`/sys/role/update`, data)
}

/**
 * 根据id删除系统角色信息
 */
export const DelById = (id) => {
  return request.post(`/sys/role/del/${id}`)
}

/**
 * 根据ids批量删除系统角色信息
 */
export const DelByIds = (ids) => {
  return request.post(`/sys/role/dels?ids=${ids}`)
}

/**
 * 系统角色启用和未启用
 */
export const UpdateStatusSysRole = (data) => {
  return request.post(`/sys/role/update/status`, data)
}

/**
 * 查询角色对应的菜单
 */
export const SelectRoleMenus = (roleId) => {
  return request.post(`/sys/role/menu/list?roleId=${roleId}`, {})
}

/**
 * 保存角色对应的菜单
 */
export const SaveRoleMenus = (data = {}) => {
  return request.post(`/sys/role/menu/save`, data)
}

/**
 * 查询角色对应的apis
 */
export const SelectRoleApis = (roleId) => {
  return request.post(`/sys/role/api/list?roleId=${roleId}`, {})
}

/**
 * 保存角色对应的apis
 */
export const SaveRoleApis = (data = {}) => {
  return request.post(`/sys/role/api/save`, data)
}

/**
 * 改变角色
 */
export const ChangeRoleIdMenus = (data) => {
  return request.post(`/sys/role/api/change?roleId=${data.roleId}`, data)
}
