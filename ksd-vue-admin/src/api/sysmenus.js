import request from '@/request/index.js'
import { C2B } from '@/utils/wordtransfer'

/**
 * 查询菜单列表并分页
 */
export const LoadTreeData = (data = { keyword: '' }) => {
  return request.post(`/sys/menus/tree?keyword=${data?.keyword}`, data) // 注意 ？ 的使用
}

/**
 * 查询父菜单
 */
export const LoadRootData = () => {
  return request.post(`/sys/menus/root`, {})
}

/**
 * 根据id查询菜单信息
 */
export const GetById = (id) => {
  return request.post(`/sys/menus/get/${id}`)
}

/**
 * 保存菜单
 */
export const SaveData = (data) => {
  return request.post(`/sys/menus/save`, data)
}

/**
 * 更新菜单信息
 */
export const UpdateData = (data) => {
  return request.post(`/sys/menus/update`, data)
}

/**
 * 根据id删除菜单信息
 */
export const DelById = (id) => {
  return request.post(`/sys/menus/del/${id}`)
}

/**
 * 根据ids批量删除菜单信息
 */
export const DelByIds = (ids) => {
  return request.post(`/sys/menus/dels?ids=${ids}`)
}

/**
 * 菜单启用和未启用
 */
export const UpdateStatus = (data) => {
  data.field = C2B(data.field) // wordtransfer 工具函数
  return request.post(`/sys/menus/update/status`, data)
}

// 复制数据
export const CopyData = (id) => {
  return request.post(`/sys/menus/copy/${id}`, {})
}
