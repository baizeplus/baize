package systemDao

import (
	"baize/app/business/system/systemModels"
	"context"
	"github.com/baizeplus/sqly"
)

type IUserPostDao interface {
	BatchUserPost(ctx context.Context, db sqly.SqlyContext, users []*systemModels.SysUserPost)
	DeleteUserPostByUserId(ctx context.Context, db sqly.SqlyContext, userId int64)
	DeleteUserPost(ctx context.Context, db sqly.SqlyContext, ids []int64)
}
