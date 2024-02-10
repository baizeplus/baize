package daoImpl

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewSysDeptDao, NewSysDictDataDao, NewSysDictTypeDao, NewSysMenuDao,
	NewSysPostDao, NewSysRoleDeptDao, NewSysRoleDao, NewSysRoleMenuDao, NewSysUserPostDao, NewSysUserRoleDao, NewSysUserDao)
