package dao

import (
	"baize/app/bzSystem/models"
	"context"
	"github.com/baizeplus/sqly"
)

type IRoleDeptDao interface {
	DeleteRoleDept(ctx context.Context, db sqly.SqlyContext, ids []int64)
	DeleteRoleDeptByRoleId(ctx context.Context, db sqly.SqlyContext, id int64)
	BatchRoleDept(ctx context.Context, db sqly.SqlyContext, list []*models.SysRoleDept)
}
