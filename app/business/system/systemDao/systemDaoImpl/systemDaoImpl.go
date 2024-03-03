package systemDaoImpl

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewSysDeptDao, NewSysDictDataDao, NewSysDictTypeDao, NewSysMenuDao,
	NewSysPostDao, NewSysUserDeptScopeDao, NewSysRoleDao, NewSysRoleMenuDao, NewSysUserPostDao, NewSysUserRoleDao, NewSysUserDao)
