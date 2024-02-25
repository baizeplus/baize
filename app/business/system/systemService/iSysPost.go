package systemService

import (
	"baize/app/business/system/systemModels"
	"github.com/gin-gonic/gin"
)

type IPostService interface {
	PostExport(c *gin.Context, role *systemModels.SysPostDQL) (data []byte)
	SelectPostList(c *gin.Context, post *systemModels.SysPostDQL) (list []*systemModels.SysPostVo, count *int64)
	SelectPostById(c *gin.Context, postId int64) (Post *systemModels.SysPostVo)
	InsertPost(c *gin.Context, post *systemModels.SysPostVo)
	UpdatePost(c *gin.Context, post *systemModels.SysPostVo)
	DeletePostByIds(c *gin.Context, postId []int64)
	SelectUserPostGroupByUserId(c *gin.Context, userId int64) string
}
