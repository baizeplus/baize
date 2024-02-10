package service

import (
	"baize/app/bzSystem/models"
	"github.com/gin-gonic/gin"
)

type IPostService interface {
	PostExport(c *gin.Context, role *models.SysPostDQL) (data []byte)
	SelectPostList(c *gin.Context, post *models.SysPostDQL) (list []*models.SysPostVo, count *int64)
	SelectPostById(c *gin.Context, postId int64) (Post *models.SysPostVo)
	InsertPost(c *gin.Context, post *models.SysPostVo)
	UpdatePost(c *gin.Context, post *models.SysPostVo)
	DeletePostByIds(c *gin.Context, postId []int64)
	SelectUserPostGroupByUserId(c *gin.Context, userId int64) string
}
