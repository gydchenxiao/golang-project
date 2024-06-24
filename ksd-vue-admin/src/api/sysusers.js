import request from '@/request/index.js'

// 查询系统用户信息
// error
export const LoadData = (data) => {
  return request.post(`/sys/user/load`, data)
}
// 上面的写法是有问题的，下面的写法得加上 params: ,不然后端获取不到前端的分页请求参数page,pageSize,keyword....数据值
// export const LoadData = (data) => {
//   return request({
//     url: '/sys/user/load',
//     method: 'post',
//     params: data
//   })
// }

/**
 * 系统用户启用和未启用、删除和未删除
 */
export const UpdateStatusSysUser = (data) => {
  return request.post(`/sys/user/update/status`, data)
}

/**
 * 根据id查询系统用户信息
 */
export const GetById = (id) => {
  return request.post(`/sys/user/get/${id}`)
}

/**
 * 保存系统用户
 */
export const SaveData = (data) => {
  return request.post(`/sys/user/save`, data)
}

// 修改
export const UpdateData = (data) => {
  return request.post(`/sys/user/update`, data)
}

// 修改密码
export const ResetPassword = (data) => {
  return request.post(`/sys/user/updatePwd`, data)
}

/**
 * 系统用户授权角色
 */
export const SaveUserRole = (data) => {
  return request.post(`/sys/user/role/save`, data)
}
