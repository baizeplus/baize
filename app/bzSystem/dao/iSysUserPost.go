package systemDao

import (
	"baize/app/bzSystem/models"
	"context"
	"github.com/baizeplus/sqly"
)

type IUserPostDao interface {
	BatchUserPost(ctx context.Context, db sqly.SqlyContext, users []*models.SysUserPost)
	DeleteUserPostByUserId(ctx context.Context, db sqly.SqlyContext, userId int64)
	DeleteUserPost(ctx context.Context, db sqly.SqlyContext, ids []int64)
}
