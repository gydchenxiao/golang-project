import request from '@/request/index.js'
import { C2B } from '@/utils/wordtransfer'

/**
 * 查询菜单权限列表并分页
 */
export const LoadTreeData = (data = { keyword: '' }) => {
  return request.post(`/sys/apis/tree?keyword=${data?.keyword}`, data) // 注意 ？ 的使用
}

/**
 * 根据id查询菜单权限信息
 */
export const GetById = (id) => {
  return request.post(`/sys/apis/get/${id}`)
}

/**
 * 保存菜单权限
 */
export const LoadRootData = () => {
  return request.post(`/sys/apis/root`)
}

/**
 * 保存菜单权限
 */
export const SaveData = (data) => {
  return request.post(`/sys/apis/save`, data)
}

/**
 * 更新菜单权限信息
 */
export const UpdateData = (data) => {
  return request.post(`/sys/apis/update`, data)
}

/**
 * 根据id删除菜单权限信息
 */
export const DelById = (id) => {
  return request.post(`/sys/apis/del/${id}`)
}

/**
 * 根据ids批量删除菜单权限信息
 */
export const DelByIds = (ids) => {
  return request.post(`/sys/apis/dels?ids=${ids}`)
}

/**
 * 菜单权限启用和未启用
 */
// export const UpdateStatusSysApis = (data) => {
//   return request.post(`/sys/apis/update/status`, data)
// }
export const UpdateStatus = (data) => {
  data.field = C2B(data.field)
  return request.post(`/sys/apis/update/status`, data)
}

// 复制数据
export const CopyData = (id) => {
  return request.post(`/sys/apis/copy/${id}`, {})
}
