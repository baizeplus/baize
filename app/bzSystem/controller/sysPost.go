package controller

import (
	"baize/app/bzSystem/models"
	"baize/app/bzSystem/service"
	"baize/app/bzSystem/service/serviceImpl"
	"baize/app/utils/baizeContext"
	"github.com/gin-gonic/gin"
)

type PostController struct {
	ps service.IPostService
}

func NewPostController(ps *serviceImpl.PostService) *PostController {
	return &PostController{ps: ps}
}

// PostList 查询岗位列表查询
// @Summary 查询岗位列表查询
// @Description 查询岗位列表查询
// @Tags 岗位相关
// @Param  object query models.SysPostDQL true "查询信息"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData{data=response.ResponseData{Rows=[]models.SysPostVo}}  "成功"
// @Router /system/post/list  [get]
func (pc *PostController) PostList(c *gin.Context) {
	post := new(models.SysPostDQL)
	_ = c.ShouldBind(post)
	list, count := pc.ps.SelectPostList(c, post)
	baizeContext.SuccessListData(c, list, count)

}

func (pc *PostController) PostExport(c *gin.Context) {
	//bzc := baizeContext.NewBaiZeContext(c)
	//post := new(models.SysPostDQL)
	//_ = c.ShouldBind(post)
	//data := pc.ps.PostExport(post)
	//bzc.DataPackageExcel(data)
}

// PostGetInfo 根据岗位ID获取岗位信息
// @Summary 根据岗位ID获取岗位信息
// @Description 根据岗位ID获取岗位信息
// @Tags 岗位相关
// @Param id path string true "PostId"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData{data=models.SysPostVo}  "成功"
// @Router /system/post/{postId}  [get]
func (pc *PostController) PostGetInfo(c *gin.Context) {
	postId := baizeContext.ParamInt64(c, "postId")
	if postId == 0 {
		baizeContext.ParameterError(c)
		return
	}
	baizeContext.SuccessData(c, pc.ps.SelectPostById(c, postId))
}

// PostAdd 添加岗位
// @Summary 添加岗位
// @Description 添加岗位
// @Tags 岗位相关
// @Param  object body models.SysPostVo true "公司信息"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData "成功"
// @Router /system/post  [post]
func (pc *PostController) PostAdd(c *gin.Context) {
	sysPost := new(models.SysPostVo)
	if err := c.ShouldBindJSON(sysPost); err != nil {
		baizeContext.ParameterError(c)
		return
	}
	sysPost.SetCreateBy(baizeContext.GetUserId(c))
	pc.ps.InsertPost(c, sysPost)
	baizeContext.Success(c)
}

// PostEdit 修改岗位
// @Summary 修改岗位
// @Description 修改岗位
// @Tags 岗位相关
// @Param  object body models.SysPostVo true "公司信息"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData "成功"
// @Router /system/post  [put]
func (pc *PostController) PostEdit(c *gin.Context) {
	post := new(models.SysPostVo)
	if err := c.ShouldBindJSON(post); err != nil {
		baizeContext.ParameterError(c)
		return
	}
	post.SetUpdateBy(baizeContext.GetUserId(c))
	pc.ps.UpdatePost(c, post)
	baizeContext.Success(c)

}

// PostRemove 删除岗位
// @Summary 删除岗位
// @Description 删除岗位
// @Tags 岗位相关
// @Param ids path []string true "postId"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData "成功"
// @Router /system/post [delete]
func (pc *PostController) PostRemove(c *gin.Context) {
	pc.ps.DeletePostByIds(c, baizeContext.ParamInt64Array(c, "postIds"))
	baizeContext.Success(c)
}
