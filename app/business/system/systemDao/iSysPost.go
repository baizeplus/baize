package systemDao

import (
	"baize/app/business/system/systemModels"
	"context"
	"github.com/baizeplus/sqly"
)

type IPostDao interface {
	SelectPostAll(ctx context.Context, db sqly.SqlyContext) (sysPost []*systemModels.SysPostVo)
	SelectPostListByUserId(ctx context.Context, db sqly.SqlyContext, userId int64) (list []int64)
	SelectPostList(ctx context.Context, db sqly.SqlyContext, post *systemModels.SysPostDQL) (list []*systemModels.SysPostVo, total int64)
	SelectPostListAll(ctx context.Context, db sqly.SqlyContext, post *systemModels.SysPostDQL) (list []*systemModels.SysPostVo)
	SelectPostById(ctx context.Context, db sqly.SqlyContext, postId int64) (dictData *systemModels.SysPostVo)
	InsertPost(ctx context.Context, db sqly.SqlyContext, post *systemModels.SysPostVo)
	UpdatePost(ctx context.Context, db sqly.SqlyContext, post *systemModels.SysPostVo)
	DeletePostByIds(ctx context.Context, db sqly.SqlyContext, dictCodes []int64)
	SelectPostNameListByUserId(ctx context.Context, db sqly.SqlyContext, userId int64) (list []string)
}
