package controller

import (
	"baize/app/bzSystem/models"
	"baize/app/bzSystem/service"
	"baize/app/bzSystem/service/serviceImpl"
	"baize/app/utils/baizeContext"
	"github.com/gin-gonic/gin"
)

type User struct {
	us service.IUserService
	ps service.IPostService
	rs service.IRoleService
}

func NewUser(
	us *serviceImpl.UserService,
	ps *serviceImpl.PostService,
	rs *serviceImpl.RoleService,
) *User {
	return &User{
		us: us,
		ps: ps,
		rs: rs,
	}
}

// ChangeStatus 修改用户状态
// @Summary 修改用户状态
// @Description 修改用户状态
// @Tags 用户相关
// @Param  object body models.EditUserStatus true "用户信息"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData  "成功"
// @Router /system/user/changeStatus [put]
func (uc *User) ChangeStatus(c *gin.Context) {

	sysUser := new(models.EditUserStatus)
	if err := c.ShouldBindJSON(sysUser); err != nil {
		baizeContext.ParameterError(c)
		return
	}
	sysUser.SetUpdateBy(baizeContext.GetUserId(c))
	uc.us.UpdateUserStatus(c, sysUser)
	baizeContext.Success(c)
}

// ResetPwd 重置密码
// @Summary 重置密码
// @Description 重置密码
// @Tags 用户相关
// @Param  object body models.ResetPwd true "密码"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData  "成功"
// @Router /system/user/resetPwd [put]
func (uc *User) ResetPwd(c *gin.Context) {
	resetPwd := new(models.ResetPwd)
	if err := c.ShouldBindJSON(resetPwd); err != nil {
		baizeContext.ParameterError(c)
		return
	}
	uc.us.ResetPwd(c, resetPwd.UserId, resetPwd.Password)
	baizeContext.Success(c)
}

// UserEdit 修改用户
// @Summary 修改用户
// @Description 修改用户
// @Tags 用户相关
// @Param  object body models.SysUserDML true "用户信息"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData  "成功"
// @Router /system/user  [put]
func (uc *User) UserEdit(c *gin.Context) {

	sysUser := new(models.SysUserDML)
	_ = c.ShouldBindJSON(sysUser)

	if uc.us.CheckPhoneUnique(c, sysUser.UserId, sysUser.Phonenumber) {
		baizeContext.Waring(c, "新增用户'"+sysUser.Phonenumber+"'失败，手机号码已存在")
		return
	}
	if uc.us.CheckEmailUnique(c, sysUser.UserId, sysUser.Email) {
		baizeContext.Waring(c, "新增用户'"+sysUser.Email+"'失败，邮箱账号已存在")
		return
	}
	sysUser.SetUpdateBy(baizeContext.GetUserId(c))
	uc.us.UpdateUser(c, sysUser)
	baizeContext.Success(c)
}

// UserAdd 添加用户
// @Summary 添加用户
// @Description 添加用户
// @Tags 用户相关
// @Param  object body models.SysUserDML true "用户信息"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData  "成功"
// @Router /system/user  [post]
func (uc *User) UserAdd(c *gin.Context) {

	sysUser := new(models.SysUserDML)
	_ = c.ShouldBindJSON(sysUser)
	if sysUser.DeptId == 0 {
		sysUser.DeptId = baizeContext.GetDeptId(c)
	}
	if uc.us.CheckUserNameUnique(c, sysUser.UserName) {
		baizeContext.Waring(c, "新增用户'"+sysUser.UserName+"'失败，登录账号已存在")
		return
	}
	if uc.us.CheckPhoneUnique(c, sysUser.UserId, sysUser.Phonenumber) {
		baizeContext.Waring(c, "新增用户'"+sysUser.Phonenumber+"'失败，手机号码已存在")
		return
	}

	if uc.us.CheckEmailUnique(c, sysUser.UserId, sysUser.Email) {
		baizeContext.Waring(c, "新增用户'"+sysUser.Email+"'失败，邮箱账号已存在")
		return
	}
	sysUser.SetCreateBy(baizeContext.GetUserId(c))
	uc.us.InsertUser(c, sysUser)
	baizeContext.Success(c)
}

// UserList 查询用户列表
// @Summary 查询用户列表
// @Description 查询用户列表
// @Tags 用户相关
// @Param  object query models.SysUserDQL true "查询信息"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData{data=response.ResponseData{Rows=[]models.SysUserVo}}  "成功"
// @Router /system/user  [get]
func (uc *User) UserList(c *gin.Context) {

	user := new(models.SysUserDQL)
	_ = c.ShouldBind(user)
	user.DataScope = baizeContext.GetDataScope(c, "d")
	list, count := uc.us.SelectUserList(c, user)
	baizeContext.SuccessListData(c, list, count)

}

// UserGetInfo 获取当前用户信息
// @Summary 获取当前用户信息
// @Description 获取当前用户信息
// @Tags 用户相关
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData{data=models.Accredit}  "成功"
// @Router /system/user/  [get]
func (uc *User) UserGetInfo(c *gin.Context) {
	baizeContext.SuccessData(c, uc.us.SelectAccredit(c))
}

// UserAuthRole 根据用户编号获取授权角色
// @Summary 根据用户编号获取授权角色
// @Description 根据用户编号获取授权角色
// @Tags 用户相关
// @Param id path string true "userId"
// @Security BearerAuth
// @Produce application/json
// // @Success 200 {object}  response.ResponseData{data=models.UserAndRoles}  "成功"
// @Router /system/user/authRole/{userId}  [get]
func (uc *User) UserAuthRole(c *gin.Context) {
	userId := baizeContext.ParamInt64(c, "userId")
	if userId == 0 {
		baizeContext.ParameterError(c)
	}
	baizeContext.SuccessData(c, uc.us.GetUserAuthRole(c, userId))
}

// UserGetInfoById 根据用户ID获取用户信息
// @Summary 根据用户ID获取用户信息
// @Description 根据用户ID获取用户信息
// @Tags 用户相关
// @Param id path int64 true "userId"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData{data=models.UserAndAccredit}  "成功"
// @Router /system/user/{userId}  [get]
func (uc *User) UserGetInfoById(c *gin.Context) {
	userId := baizeContext.ParamInt64(c, "userId")
	if userId == 0 {
		baizeContext.ParameterError(c)
		return
	}

	baizeContext.SuccessData(c, uc.us.SelectUserAndAccreditById(c, userId))

}

// UserRemove 删除用户
// @Summary 删除用户
// @Description 删除用户
// @Tags 系统用户
// @Param userIds path []int64 true "userIds"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object} response.ResponseData
// @Router /system/user/:userIds [delete]
func (uc *User) UserRemove(c *gin.Context) {
	uc.us.DeleteUserByIds(c, baizeContext.ParamInt64Array(c, "userIds"))
	baizeContext.Success(c)
}

// UserImportData 导入用户
// @Summary 导入用户
// @Description 导入用户
// @Tags 系统用户
// @Param file formData file true "file"
// @Accept multipart/form-data
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object} response.ResponseData
// @Router /system/user/importData [post]
func (uc *User) UserImportData(c *gin.Context) {
	//bzc := baizeContext.NewBaiZeContext(c)
	//
	//fileHeader, err := c.FormFile("file")
	//if err != nil {
	//	panic(err)
	//}
	//file, _ := fileHeader.Open()
	//defer file.Close()
	//excelFile, _ := excelize.OpenReader(file)
	//rows := excelFile.GetRows("Sheet1")
	//loginUser := bzc.GetUser()
	//data, num := uc.us.UserImportData(rows, loginUser.UserId, loginUser.DeptId)
	//if num > 0 {
	//	bzc.Waring(data)
	//	return
	//}
	//bzc.SuccessMsg(data)
}

// UserExport 导出用户
// @Summary 导出用户
// @Description 导出用户
// @Tags 系统用户
// @Param  object query models.SysUserDQL true "查询信息"
// @Security BearerAuth
// @Produce application/octet-stream
// @Success 200 {object} []byte
// @Router /system/user/export [post]
func (uc *User) UserExport(c *gin.Context) {
	//bzc := baizeContext.NewBaiZeContext(c)
	//user := new(models.SysUserDQL)
	//_ = c.ShouldBind(user)
	//user.SetDataScope(bzc.GetUser(), "d", "u")
	//bzc.DataPackageExcel(uc.us.UserExport(user))
	//return
}

// ImportTemplate 获取导入模版
// @Summary 导出用户
// @Description 导出用户
// @Tags 系统用户
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object} []byte
// @Router /system/user/input [post]
func (uc *User) ImportTemplate(c *gin.Context) {
	//bzc := baizeContext.NewBaiZeContext(c)
	//data := uc.us.ImportTemplate()
	//bzc.DataPackageExcel(data)
	//return
}

// InsertAuthRole 授权角色
// @Summary 授权角色
// @Description 授权角色
// @Tags 用户相关
// @Param  string query string true "角色id"
// @Param  string query string true "用户id"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData{data=response.ResponseData{Rows=[]models.SysUserVo}}  "成功"
// @Router /system/user/authRole  [put]
func (uc *User) InsertAuthRole(c *gin.Context) {
	array := baizeContext.QueryInt64Array(c, "roleIds")
	uc.us.InsertUserAuth(c, baizeContext.QueryInt64(c, "userId"), array)
	baizeContext.Success(c)
}
