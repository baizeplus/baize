package systemServiceImpl

import (
	"baize/app/business/system/systemDao"
	"baize/app/business/system/systemModels"
	"baize/app/business/system/systemService"
	"baize/app/utils/excel"
	"baize/app/utils/snowflake"
	"github.com/gin-gonic/gin"
)

type PostService struct {
	postDao systemDao.IPostDao
}

func NewPostService(pd systemDao.IPostDao) systemService.IPostService {
	return &PostService{
		postDao: pd,
	}
}

func (postService *PostService) SelectPostList(c *gin.Context, post *systemModels.SysPostDQL) (list []*systemModels.SysPostVo, total int64) {
	return postService.postDao.SelectPostList(c, post)

}
func (postService *PostService) PostExport(c *gin.Context, post *systemModels.SysPostDQL) (data []byte) {
	list := postService.postDao.SelectPostListAll(c, post)
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
	return postService.postDao.SelectPostById(c, postId)

}

func (postService *PostService) InsertPost(c *gin.Context, post *systemModels.SysPostVo) {
	post.PostId = snowflake.GenID()
	postService.postDao.InsertPost(c, post)
}

func (postService *PostService) UpdatePost(c *gin.Context, post *systemModels.SysPostVo) {
	postService.postDao.UpdatePost(c, post)
}
func (postService *PostService) DeletePostByIds(c *gin.Context, postId []int64) {
	postService.postDao.DeletePostByIds(c, postId)
	return
}
