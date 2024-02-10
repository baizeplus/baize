package systemDao

import (
	"baize/app/bzSystem/models"
	"context"
	"github.com/baizeplus/sqly"
)

type IPostDao interface {
	SelectPostAll(ctx context.Context, db sqly.SqlyContext) (sysPost []*models.SysPostVo)
	SelectPostListByUserId(ctx context.Context, db sqly.SqlyContext, userId int64) (list []int64)
	SelectPostList(ctx context.Context, db sqly.SqlyContext, post *models.SysPostDQL) (list []*models.SysPostVo, total *int64)
	SelectPostById(ctx context.Context, db sqly.SqlyContext, postId int64) (dictData *models.SysPostVo)
	InsertPost(ctx context.Context, db sqly.SqlyContext, post *models.SysPostVo)
	UpdatePost(ctx context.Context, db sqly.SqlyContext, post *models.SysPostVo)
	DeletePostByIds(ctx context.Context, db sqly.SqlyContext, dictCodes []int64)
	SelectPostNameListByUserId(ctx context.Context, db sqly.SqlyContext, userId int64) (list []string)
}
