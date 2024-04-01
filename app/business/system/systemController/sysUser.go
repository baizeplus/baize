package systemController

import (
	"baize/app/business/system/systemModels"
	"baize/app/business/system/systemService"
	"baize/app/business/system/systemService/systemServiceImpl"
	"baize/app/constant/dataScopeAspect"
	"baize/app/utils/baizeContext"
	"baize/app/utils/response"
	"github.com/gin-gonic/gin"
)

type User struct {
	us systemService.IUserService
	ps systemService.IPostService
	rs systemService.IRoleService
}

func NewUser(
	us *systemServiceImpl.UserService,
	ps *systemServiceImpl.PostService,
	rs *systemServiceImpl.RoleService,
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
// @Param  object body systemModels.EditUserStatus true "用户信息"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData  "成功"
// @Router /system/user/changeStatus [put]
func (uc *User) ChangeStatus(c *gin.Context) {

	sysUser := new(systemModels.EditUserStatus)
	if err := c.ShouldBindJSON(sysUser); err != nil {
		baizeContext.ParameterError(c)
		return
	}
	if sysUser.UserId == baizeContext.GetUserId(c) {
		baizeContext.Waring(c, response.ForbiddenOperation)
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
// @Param  object body systemModels.ResetPwd true "密码"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData  "成功"
// @Router /system/user/resetPwd [put]
func (uc *User) ResetPwd(c *gin.Context) {
	resetPwd := new(systemModels.ResetPwd)
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
// @Param  object body systemModels.SysUserDML true "用户信息"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData  "成功"
// @Router /system/user  [put]
func (uc *User) UserEdit(c *gin.Context) {

	sysUser := new(systemModels.SysUserDML)
	_ = c.ShouldBindJSON(sysUser)
	if sysUser.UserId == baizeContext.GetUserId(c) {
		baizeContext.Waring(c, response.ForbiddenOperation)
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
	sysUser.SetUpdateBy(baizeContext.GetUserId(c))
	uc.us.UpdateUser(c, sysUser)
	baizeContext.Success(c)
}

// UpdateUserDataScope 修改数据权限
// @Summary 修改数据权限
// @Description 修改数据权限
// @Tags 用户相关
// @Param  object body systemModels.SysUserDataScope true "用户权限信息"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData  "成功"
// @Router /system/user/dataScope  [put]
func (uc *User) UpdateUserDataScope(c *gin.Context) {
	uds := new(systemModels.SysUserDataScope)
	if err := c.ShouldBindJSON(uds); err != nil {
		baizeContext.ParameterError(c)
		return
	}
	if uds.DataScope == dataScopeAspect.DataScopeAll && baizeContext.GetDataScopeAspect(c) != dataScopeAspect.DataScopeAll {
		baizeContext.Waring(c, "数据权限范围不能大于自己")
		return
	}
	if uds.UserId == baizeContext.GetUserId(c) {
		baizeContext.Waring(c, response.ForbiddenOperation)
		return
	}
	uc.us.UpdateUserDataScope(c, uds)
	baizeContext.Success(c)
}

// SelectUserDataScope 查询用户数据权限
// @Summary 查询用户数据权限
// @Description 查询用户数据权限
// @Tags 用户相关
// @Param id path int64 true "userId"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData  "成功"
// @Router /system/user/dataScope/{userId}   [get]
func (uc *User) SelectUserDataScope(c *gin.Context) {
	userId := baizeContext.ParamInt64(c, "userId")
	if userId == 0 {
		baizeContext.ParameterError(c)
		return
	}
	baizeContext.SuccessData(c, uc.us.SelectUserDataScope(c, userId))
}

// UserAdd 添加用户
// @Summary 添加用户
// @Description 添加用户
// @Tags 用户相关
// @Param  object body systemModels.SysUserDML true "用户信息"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData  "成功"
// @Router /system/user  [post]
func (uc *User) UserAdd(c *gin.Context) {

	sysUser := new(systemModels.SysUserDML)
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
// @Param  object query systemModels.SysUserDQL true "查询信息"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData{data=response.ListData{Rows=[]systemModels.SysUserVo}}  "成功"
// @Router /system/user  [get]
func (uc *User) UserList(c *gin.Context) {
	user := new(systemModels.SysUserDQL)
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
// @Success 200 {object}  response.ResponseData{data=systemModels.Accredit}  "成功"
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
// // @Success 200 {object}  response.ResponseData{data=systemModels.UserAndRoles}  "成功"
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
// @Success 200 {object}  response.ResponseData{data=systemModels.UserAndAccredit}  "成功"
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
	array := baizeContext.ParamInt64Array(c, "userIds")
	for _, i := range array {
		if i == baizeContext.GetUserId(c) {
			baizeContext.Waring(c, response.ForbiddenOperation)
			return
		}
	}
	uc.us.DeleteUserByIds(c, array)
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
	fileHeader, err := c.FormFile("file")
	if err != nil {
		baizeContext.ParameterError(c)
	}
	data, num := uc.us.UserImportData(c, fileHeader)

	if num > 0 {
		baizeContext.Waring(c, data)
		return
	}
	baizeContext.SuccessMsg(c, data)
}

// UserExport 导出用户
// @Summary 导出用户
// @Description 导出用户
// @Tags 系统用户
// @Param  object query systemModels.SysUserDQL true "查询信息"
// @Security BearerAuth
// @Produce application/octet-stream
// @Success 200 {object} []byte
// @Router /system/user/export [post]
func (uc *User) UserExport(c *gin.Context) {
	user := new(systemModels.SysUserDQL)
	_ = c.ShouldBind(user)
	user.DataScope = baizeContext.GetDataScope(c, "d")
	baizeContext.DataPackageExcel(c, uc.us.UserExport(c, user))
	return
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
	data := uc.us.ImportTemplate(c)
	baizeContext.DataPackageExcel(c, data)
	return
}

// InsertAuthRole 授权角色
// @Summary 授权角色
// @Description 授权角色
// @Tags 用户相关
// @Param  string query string true "角色id"
// @Param  string query string true "用户id"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData{data=response.ListData{Rows=[]systemModels.SysUserVo}}  "成功"
// @Router /system/user/authRole  [put]
func (uc *User) InsertAuthRole(c *gin.Context) {
	userId := baizeContext.QueryInt64(c, "userId")
	if userId == baizeContext.GetUserId(c) {
		baizeContext.Waring(c, response.ForbiddenOperation)
		return
	}
	array := baizeContext.QueryInt64Array(c, "roleIds")
	uc.us.InsertUserAuth(c, userId, array)
	baizeContext.Success(c)
}
