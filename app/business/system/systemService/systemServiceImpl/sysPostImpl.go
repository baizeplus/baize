package systemServiceImpl

import (
	"baize/app/business/system/systemDao"
	"baize/app/business/system/systemDao/systemDaoImpl"
	"baize/app/business/system/systemModels"
	"baize/app/utils/excel"
	"baize/app/utils/snowflake"
	"github.com/baizeplus/sqly"
	"github.com/gin-gonic/gin"
)

type PostService struct {
	data    *sqly.DB
	postDao systemDao.IPostDao
}

func NewPostService(data *sqly.DB, pd *systemDaoImpl.SysPostDao) *PostService {
	return &PostService{
		data:    data,
		postDao: pd,
	}
}

func (postService *PostService) SelectPostList(c *gin.Context, post *systemModels.SysPostDQL) (list []*systemModels.SysPostVo, total int64) {
	return postService.postDao.SelectPostList(c, postService.data, post)

}
func (postService *PostService) PostExport(c *gin.Context, post *systemModels.SysPostDQL) (data []byte) {
	list := postService.postDao.SelectPostListAll(c, postService.data, post)
	toExcel, err := excel.SliceToExcel(list)
	if err != nil {
		panic(err)
	}
	buffer, err := toExcel.WriteToBuffer()
	if err != nil {
		panic(err)
	}
	return buffer.Bytes()
}

func (postService *PostService) SelectPostById(c *gin.Context, postId int64) (Post *systemModels.SysPostVo) {
	return postService.postDao.SelectPostById(c, postService.data, postId)

}

func (postService *PostService) InsertPost(c *gin.Context, post *systemModels.SysPostVo) {
	post.PostId = snowflake.GenID()
	postService.postDao.InsertPost(c, postService.data, post)
}

func (postService *PostService) UpdatePost(c *gin.Context, post *systemModels.SysPostVo) {
	postService.postDao.UpdatePost(c, postService.data, post)
}
func (postService *PostService) DeletePostByIds(c *gin.Context, postId []int64) {
	postService.postDao.DeletePostByIds(c, postService.data, postId)
	return
}
