package systemDaoImpl

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewSysDeptDao, NewSysDictDataDao, NewSysDictTypeDao, NewSysConfigDao, NewSysNoticeDao, NewSysPermissionDao,
	NewSysPostDao, NewSysUserDeptScopeDao, NewSysRoleDao, NewSysRoleMenuDao, NewSysUserPostDao, NewSysUserRoleDao, NewSysUserDao)
