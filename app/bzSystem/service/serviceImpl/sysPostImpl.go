package serviceImpl

import (
	"baize/app/bzSystem/dao"
	"baize/app/bzSystem/dao/daoImpl"
	"baize/app/bzSystem/models"
	"baize/app/utils/snowflake"
	"github.com/baizeplus/sqly"
	"github.com/gin-gonic/gin"
	"strings"
)

type PostService struct {
	data    *sqly.DB
	postDao dao.IPostDao
}

func NewPostService(data *sqly.DB, pd *daoImpl.SysPostDao) *PostService {
	return &PostService{
		data:    data,
		postDao: pd,
	}
}

func (postService *PostService) SelectPostList(c *gin.Context, post *models.SysPostDQL) (list []*models.SysPostVo, count *int64) {
	return postService.postDao.SelectPostList(c, postService.data, post)

}
func (postService *PostService) PostExport(c *gin.Context, post *models.SysPostDQL) (data []byte) {
	//list, _ := postService.postDao.SelectPostList(c,postService.data, post)
	//rows := models.SysPostListToRows(list)
	//return exceLize.SetRows(rows)
	return nil
}

func (postService *PostService) SelectPostById(c *gin.Context, postId int64) (Post *models.SysPostVo) {
	return postService.postDao.SelectPostById(c, postService.data, postId)

}

func (postService *PostService) InsertPost(c *gin.Context, post *models.SysPostVo) {
	post.PostId = snowflake.GenID()
	postService.postDao.InsertPost(c, postService.data, post)
}

func (postService *PostService) UpdatePost(c *gin.Context, post *models.SysPostVo) {
	postService.postDao.UpdatePost(c, postService.data, post)
}
func (postService *PostService) DeletePostByIds(c *gin.Context, postId []int64) {
	postService.postDao.DeletePostByIds(c, postService.data, postId)
	return
}
func (postService *PostService) SelectUserPostGroupByUserId(c *gin.Context, userId int64) string {

	return strings.Join(postService.postDao.SelectPostNameListByUserId(c, postService.data, userId), ",")

}
