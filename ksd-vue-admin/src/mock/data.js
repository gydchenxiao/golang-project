export const users = [
  {
    name: 'visitor',
    roleId: 'visitor',
    password: 'visitor'
  },
  {
    name: 'master',
    roleId: 'master',
    password: 'master'
  },
  {
    name: 'admin',
    roleId: 'admin',
    password: 'admin'
  }
]

// 模拟服务端角色对应菜单信息--超级管理员
export const menuTreeData = [
  {
    id: 4,
    parentId: 0,
    name: 'DashBoard',
    path: '/dashboard',
    icon: 'home-filled',
    children: []
  },
  {
    id: 2,
    parentId: 0,
    name: 'Sys',
    path: '/sys',
    icon: 'setting',
    children: [
      {
        id: 21,
        parentId: 2,
        name: 'SysUser',
        path: '/sys/User',
        icon: 'user-filled'
      },
      {
        id: 22,
        parentId: 2,
        name: 'SysMenu',
        path: '/sys/Menu',
        icon: 'user-filled'
      },
      {
        id: 23,
        parentId: 2,
        name: 'SysRole',
        path: '/sys/Role',
        icon: 'user-filled'
      },
      {
        id: 24,
        parentId: 2,
        name: 'SysPermission',
        path: '/sys/Permission',
        icon: 'user-filled'
      },
      {
        id: 25,
        parentId: 2,
        name: 'SysNotice',
        path: '/sys/Notice',
        icon: 'user-filled'
      }
    ]
  },
  {
    id: 1,
    parentId: 0,
    name: 'App',
    path: '/app',
    icon: 'menu',
    children: [
      {
        id: 11,
        parentId: 1,
        name: 'AppUser',
        path: '/app/User',
        icon: 'user'
      },
      {
        id: 12,
        parentId: 1,
        name: 'AppDept',
        path: '/app/Dept',
        icon: 'office-building'
      },
      {
        id: 13,
        parentId: 1,
        name: 'AppRole',
        path: '/app/Role',
        icon: 'avatar'
      },
      {
        id: 14,
        parentId: 1,
        name: 'AppResource',
        path: '/app/Resource',
        icon: 'avatar'
      }
    ]
  },
  {
    id: 3,
    parentId: 0,
    name: 'Logs',
    path: '/logs',
    icon: 'document',
    children: [
      {
        id: 31,
        parentId: 1,
        name: 'LogsVisit',
        path: '/logs/Visit',
        icon: 'tickets'
      },
      {
        id: 32,
        parentId: 1,
        name: 'LogsOperation',
        path: '/logs/Operation',
        icon: 'operation'
      }
    ]
  },
  {
    id: 6,
    parentId: 0,
    name: 'Course',
    path: '/course',
    icon: 'operation',
    children: [
      {
        id: 61,
        parentId: 1,
        name: 'CourseList',
        path: '/course/List',
        icon: 'operation'
      }
    ]
  },
  {
    id: 7,
    parentId: 0,
    name: 'Video',
    path: '/video',
    icon: 'document',
    children: []
  },
  {
    id: 8,
    parentId: 0,
    name: 'Order',
    path: '/order',
    icon: 'odometer',
    children: [
      {
        id: 81,
        parentId: 8,
        name: 'OrderList',
        path: '/order/List',
        icon: 'odometer'
      }
    ]
  }
]
