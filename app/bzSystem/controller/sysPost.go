package controller

import (
	"github.com/gin-gonic/gin"
)

type PostController struct {
}

func NewPostController() *PostController {
	return &PostController{}
}

// PostList 查询岗位列表查询
// @Summary 查询岗位列表查询
// @Description 查询岗位列表查询
// @Tags 岗位相关
// @Param  object query models.SysPostDQL true "查询信息"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData{data=response.ResponseData{Rows=[]models.SysPostVo}}  "成功"
// @Router /bzSystem/post/list  [get]
func (pc *PostController) PostList(c *gin.Context) {
	//bzc := baizeContext.NewBaiZeContext(c)
	//post := new(models.SysPostDQL)
	//_ = c.ShouldBind(post)
	//list, count := pc.ps.SelectPostList(post)
	//bzc.SuccessListData(list, count)

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
// @Router /bzSystem/post/{postId}  [get]
func (pc *PostController) PostGetInfo(c *gin.Context) {
	//bzc := baizeContext.NewBaiZeContext(c)
	//postId := bzc.ParamInt64("postId")
	//if postId == 0 {
	//	bzc.ParameterError()
	//	return
	//}
	//bzc.SuccessData(pc.ps.SelectPostById(postId))
}

// PostAdd 添加岗位
// @Summary 添加岗位
// @Description 添加岗位
// @Tags 岗位相关
// @Param  object body models.SysPostVo true "公司信息"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData "成功"
// @Router /bzSystem/post  [post]
func (pc *PostController) PostAdd(c *gin.Context) {
	//bzc := baizeContext.NewBaiZeContext(c)
	//sysPost := new(models.SysPostAdd)
	//if err := c.ShouldBindJSON(sysPost); err != nil {
	//	bzc.ParameterError()
	//	return
	//}
	//sysPost.SetCreateBy(bzc.GetUserId())
	//pc.ps.InsertPost(sysPost)
	//bzc.Success()
}

// PostEdit 修改岗位
// @Summary 修改岗位
// @Description 修改岗位
// @Tags 岗位相关
// @Param  object body models.SysPostVo true "公司信息"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData "成功"
// @Router /bzSystem/post  [put]
func (pc *PostController) PostEdit(c *gin.Context) {
	//bzc := baizeContext.NewBaiZeContext(c)
	//post := new(models.SysPostEdit)
	//if err := c.ShouldBindJSON(post); err != nil {
	//	bzc.ParameterError()
	//	return
	//}
	//post.SetUpdateBy(bzc.GetUserId())
	//pc.ps.UpdatePost(post)
	//bzc.Success()

}

// PostRemove 删除岗位
// @Summary 删除岗位
// @Description 删除岗位
// @Tags 岗位相关
// @Param ids path []string true "postId"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData "成功"
// @Router /bzSystem/post [delete]
func (pc *PostController) PostRemove(c *gin.Context) {
	//bzc := baizeContext.NewBaiZeContext(c)
	//pc.ps.DeletePostByIds(bzc.ParamInt64Array("postIds"))
	//bzc.Success()
}
