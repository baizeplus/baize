package systemDao

import (
	"baize/app/business/system/systemModels"
	"context"
)

type IUserPostDao interface {
	BatchUserPost(ctx context.Context, users []*systemModels.SysUserPost)
	DeleteUserPostByUserId(ctx context.Context, userId string)
	DeleteUserPost(ctx context.Context, ids []string)
}
