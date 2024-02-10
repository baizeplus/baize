package controller

import (
	"baize/app/bzSystem/models"
	"baize/app/bzSystem/service"
	"baize/app/bzSystem/service/serviceImpl"
	"baize/app/utils/baizeContext"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	us service.IUserService
	ps service.IPostService
	rs service.IRoleService
}

func NewUserController(
	us *serviceImpl.UserService,
	ps *serviceImpl.PostService,
	rs *serviceImpl.RoleService,
) *UserController {
	return &UserController{
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
func (uc *UserController) ChangeStatus(c *gin.Context) {
	//bzc := baizeContext.NewBaiZeContext(c)
	//
	//sysUser := new(models.EditUserStatus)
	//if err := c.ShouldBindJSON(sysUser); err != nil {
	//	bzc.ParameterError()
	//	return
	//}
	//sysUser.SetUpdateBy(bzc.GetUserId())
	//uc.us.UpdateUserStatus(sysUser)
	//bzc.Success()
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
func (uc *UserController) ResetPwd(c *gin.Context) {
	//bzc := baizeContext.NewBaiZeContext(c)
	//
	//resetPwd := new(models.ResetPwd)
	//if err := c.ShouldBindJSON(resetPwd); err != nil {
	//	bzc.ParameterError()
	//	return
	//}
	//uc.us.ResetPwd(resetPwd.UserId, resetPwd.Password)
	//bzc.Success()

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
func (uc *UserController) UserEdit(c *gin.Context) {
	//bzc := baizeContext.NewBaiZeContext(c)
	//
	//sysUser := new(models.SysUserEdit)
	//if err := c.ShouldBindJSON(sysUser); err != nil {
	//	fmt.Println(err)
	//	bzc.ParameterError()
	//	return
	//}
	//if uc.us.CheckPhoneUnique(sysUser.UserId, sysUser.Phonenumber) {
	//	bzc.Waring("新增用户'" + sysUser.Phonenumber + "'失败，手机号码已存在")
	//	return
	//}
	//if uc.us.CheckEmailUnique(sysUser.UserId, sysUser.Email) {
	//	bzc.Waring("新增用户'" + sysUser.Email + "'失败，邮箱账号已存在")
	//	return
	//}
	//sysUser.SetUpdateBy(bzc.GetUserId())
	//uc.us.UpdateUser(sysUser)
	//bzc.Success()
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
func (uc *UserController) UserAdd(c *gin.Context) {
	//bzc := baizeContext.NewBaiZeContext(c)
	//
	//user := bzc.GetUser()
	//sysUser := new(models.SysUserAdd)
	//if err := c.ShouldBindJSON(sysUser); err != nil {
	//	bzc.ParameterError()
	//	return
	//}
	//if sysUser.DeptId == nil {
	//	sysUser.DeptId = user.DeptId
	//}
	//if uc.us.CheckUserNameUnique(sysUser.UserName) {
	//	bzc.Waring("新增用户'" + sysUser.UserName + "'失败，登录账号已存在")
	//	return
	//}
	//if uc.us.CheckPhoneUnique(sysUser.UserId, sysUser.Phonenumber) {
	//	bzc.Waring("新增用户'" + sysUser.Phonenumber + "'失败，手机号码已存在")
	//	return
	//}
	//
	//if uc.us.CheckEmailUnique(sysUser.UserId, sysUser.Email) {
	//	bzc.Waring("新增用户'" + sysUser.Email + "'失败，邮箱账号已存在")
	//	return
	//}
	//sysUser.SetCreateBy(user.UserId)
	//uc.us.InsertUser(sysUser)
	//bzc.Success()
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
func (uc *UserController) UserList(c *gin.Context) {

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
func (uc *UserController) UserGetInfo(c *gin.Context) {
	baizeContext.SuccessData(c, uc.us.SelectAccredit(c))
}

// UserAuthRole 根据用户编号获取授权角色
// @Summary 根据用户编号获取授权角色
// @Description 根据用户编号获取授权角色
// @Tags 用户相关
// @Param id path string true "userId"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData{data=models.Auth}  "成功"
// @Router /system/user/authRole/{userId}  [get]
func (uc *UserController) UserAuthRole(c *gin.Context) {
	//bzc := baizeContext.NewBaiZeContext(c)
	//userId := bzc.ParamInt64("userId")
	//if userId == 0 {
	//	bzc.ParameterError()
	//}
	//ar := new(models.Auth)
	//ar.User = uc.us.SelectUserById(userId)
	//role := new(models.SysRoleDQL)
	//user := bzc.GetUser()
	//role.SetDataScope(user, "d", "")
	//roles := uc.rs.SelectRoleAll(role)
	//if !utils.IsAdmin(user.UserId) {
	//	for i, role := range roles {
	//		if role.RoleId == 1 {
	//			roles = append(roles[:i], roles[i+1:]...)
	//			break
	//		}
	//	}
	//}
	//ar.Roles = roles
	//ar.RoleIds = gconv.Strings(uc.rs.SelectRoleListByUserId(userId))
	//bzc.SuccessData(ar)
}

// UserGetInfoById 根据用户ID获取用户信息
// @Summary 根据用户ID获取用户信息
// @Description 根据用户ID获取用户信息
// @Tags 用户相关
// @Param id path string true "userId"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData{data=models.Auth}  "成功"
// @Router /system/user/{userId}  [get]
func (uc *UserController) UserGetInfoById(c *gin.Context) {
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
// @Param userIds body  []string true "userIds"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object} response.ResponseData
// @Router /system/user [delete]
func (uc *UserController) UserRemove(c *gin.Context) {
	//bzc := baizeContext.NewBaiZeContext(c)
	//uc.us.DeleteUserByIds(bzc.ParamInt64Array("userIds"))
	//bzc.Success()
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
func (uc *UserController) UserImportData(c *gin.Context) {
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
func (uc *UserController) UserExport(c *gin.Context) {
	//bzc := baizeContext.NewBaiZeContext(c)
	//user := new(models.SysUserDQL)
	//_ = c.ShouldBind(user)
	//user.SetDataScope(bzc.GetUser(), "d", "u")
	//bzc.DataPackageExcel(uc.us.UserExport(user))
	//return
}

// ImportTemplate 获取导入末班
// @Summary 导出用户
// @Description 导出用户
// @Tags 系统用户
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object} []byte
// @Router /system/user/input [post]
func (uc *UserController) ImportTemplate(c *gin.Context) {
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
func (uc *UserController) InsertAuthRole(c *gin.Context) {
	//bzc := baizeContext.NewBaiZeContext(c)
	//array := bzc.QueryInt64Array("roleIds")
	//uc.us.InsertUserAuth(bzc.QueryInt64("userId"), array)
	//bzc.Success()
	//return
}
