package systemDao

import (
	"baize/app/business/system/systemModels"
	"context"
)

type IPostDao interface {
	SelectPostAll(ctx context.Context) (sysPost []*systemModels.SysPostVo)
	SelectPostListByUserId(ctx context.Context, userId string) (list []int64)
	SelectPostList(ctx context.Context, post *systemModels.SysPostDQL) (list []*systemModels.SysPostVo, total int64)
	SelectPostListAll(ctx context.Context, post *systemModels.SysPostDQL) (list []*systemModels.SysPostVo)
	SelectPostById(ctx context.Context, postId string) (dictData *systemModels.SysPostVo)
	InsertPost(ctx context.Context, post *systemModels.SysPostVo)
	UpdatePost(ctx context.Context, post *systemModels.SysPostVo)
	DeletePostByIds(ctx context.Context, postId []string)
	SelectPostNameListByUserId(ctx context.Context, userId string) (list []string)
}
