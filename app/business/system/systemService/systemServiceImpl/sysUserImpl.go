package systemServiceImpl

import (
	"baize/app/baize"
	"baize/app/business/system/systemDao"
	"baize/app/business/system/systemDao/systemDaoImpl"
	"baize/app/business/system/systemModels"
	"baize/app/business/system/systemService"
	"baize/app/constant/dataScopeAspect"
	"baize/app/constant/sessionStatus"
	"baize/app/utils/IOFile"
	"baize/app/utils/bCryptPasswordEncoder"
	"baize/app/utils/baizeContext"
	"baize/app/utils/excel"
	"baize/app/utils/snowflake"
	"context"
	"github.com/baizeplus/sqly"
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
	"mime/multipart"
	"path/filepath"
	"strconv"
	"strings"
)

type UserService struct {
	data        *sqly.DB
	userDao     systemDao.IUserDao
	userPostDao systemDao.IUserPostDao
	userRoleDao systemDao.IUserRoleDao
	roleDao     systemDao.IRoleDao
	postDao     systemDao.IPostDao
	deptDao     systemDao.IDeptDao
	uds         systemDao.IUserDeptScopeDao
	cs          systemService.IConfigService
}

func NewUserService(data *sqly.DB, ud *systemDaoImpl.SysUserDao, upd *systemDaoImpl.SysUserPostDao, urd *systemDaoImpl.SysUserRoleDao, dd *systemDaoImpl.SysDeptDao, rd *systemDaoImpl.SysRoleDao, pd *systemDaoImpl.SysPostDao, uds *systemDaoImpl.SysUserDeptScopeDao, cs *ConfigService) *UserService {
	return &UserService{
		data:        data,
		userDao:     ud,
		userPostDao: upd,
		userRoleDao: urd,
		roleDao:     rd,
		postDao:     pd,
		uds:         uds,
		deptDao:     dd,
		cs:          cs,
	}
}

func (userService *UserService) SelectUserByUserName(c *gin.Context, userName string) *systemModels.User {
	return userService.userDao.SelectUserByUserName(c, userService.data, userName)

}
func (userService *UserService) SelectUserList(c *gin.Context, user *systemModels.SysUserDQL) (sysUserList []*systemModels.SysUserVo, total int64) {
	return userService.userDao.SelectUserList(c, userService.data, user)
}

func (userService *UserService) UserExport(c *gin.Context, user *systemModels.SysUserDQL) (data []byte) {
	sysUserList := userService.userDao.SelectUserListAll(c, userService.data, user)
	toExcel, err := excel.SliceToExcel(sysUserList)
	if err != nil {
		panic(err)
	}
	buffer, err := toExcel.WriteToBuffer()
	if err != nil {
		panic(err)
	}
	return buffer.Bytes()
}

func (userService *UserService) ImportTemplate(c *gin.Context) (data []byte) {
	f := excelize.NewFile()
	dept := new(systemModels.SysDeptDQL)
	dept.DataScope = baizeContext.GetDataScope(c, "d")
	list := userService.deptDao.SelectDeptList(c, userService.data, dept)
	all := systemModels.GetParentNameAll(list)
	sqref := "C2:C100"
	dvRange1 := excelize.NewDataValidation(true)
	dvRange1.Sqref = sqref
	dvRange1.SetDropList(all)
	dvRange1.ShowInputMessage = true
	dvRange1.ShowErrorMessage = true
	f.AddDataValidation("Sheet1", dvRange1)

	sqref2 := "F2:F100"
	dvRange2 := excelize.NewDataValidation(true)
	dvRange2.Sqref = sqref2
	dvRange2.SetDropList([]string{"男", "女", "未知"})
	dvRange2.ShowInputMessage = true
	dvRange2.ShowErrorMessage = true
	f.AddDataValidation("Sheet1", dvRange1)
	f.AddDataValidation("Sheet1", dvRange2)
	f.SetSheetRow("Sheet1", "A1", &[]string{"登录名称", "用户名", "部门", "邮箱", "手机号", "性别"})
	border := []excelize.Border{
		{Type: "top", Style: 1, Color: "cccccc"},
		{Type: "left", Style: 1, Color: "cccccc"},
		{Type: "right", Style: 1, Color: "cccccc"},
		{Type: "bottom", Style: 1, Color: "cccccc"},
	}
	// 定义标题行单元格样式
	headerStyle, err := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{Bold: true, Color: "ffffff"},
		Fill: excelize.Fill{
			Type: "pattern", Color: []string{"515151"}, Pattern: 1},
		Border: border},
	)
	// 为标题行设置样式
	f.SetCellStyle("Sheet1", "A1", "F1", headerStyle)
	f.SetColWidth("Sheet1", "A", "B", 15)
	f.SetColWidth("Sheet1", "C", "C", 30)
	f.SetColWidth("Sheet1", "D", "E", 20)
	f.SetColWidth("Sheet1", "F", "F", 10)

	buffer, err := f.WriteToBuffer()
	if err != nil {
		panic(err)
	}
	return buffer.Bytes()

}

func (userService *UserService) SelectUserAndAccreditById(c *gin.Context, userId int64) (sysUser *systemModels.UserAndAccredit) {
	uaa := new(systemModels.UserAndAccredit)
	uaa.User = userService.userDao.SelectUserById(c, userService.data, userId)
	uaa.Roles = userService.roleDao.SelectRoleAll(c, userService.data, new(systemModels.SysRoleDQL))
	uaa.Posts = userService.postDao.SelectPostAll(c, userService.data)
	rIds := userService.roleDao.SelectRoleListByUserId(c, userService.data, userId)
	pIds := userService.postDao.SelectPostListByUserId(c, userService.data, userId)
	uaa.RoleIds = make([]string, 0, len(rIds))
	for _, id := range rIds {
		uaa.RoleIds = append(uaa.RoleIds, strconv.FormatInt(id, 10))
	}
	uaa.PostIds = make([]string, 0, len(pIds))
	for _, id := range pIds {
		uaa.PostIds = append(uaa.PostIds, strconv.FormatInt(id, 10))
	}
	return uaa
}
func (userService *UserService) SelectAccredit(c *gin.Context) (sysUser *systemModels.Accredit) {
	ua := new(systemModels.Accredit)
	ua.Roles = userService.roleDao.SelectRoleAll(c, userService.data, new(systemModels.SysRoleDQL))
	ua.Posts = userService.postDao.SelectPostAll(c, userService.data)
	return ua
}

func (userService *UserService) InsertUser(c *gin.Context, sysUser *systemModels.SysUserDML) {
	sysUser.UserId = snowflake.GenID()
	sysUser.Password = bCryptPasswordEncoder.HashPassword(sysUser.Password)
	tx, err := userService.data.Beginx()
	if err != nil {
		panic(err)
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // re-throw panic after Rollback
		} else {
			err = tx.Commit() // err is nil; if Commit returns error update err
		}
	}()
	sysUser.DataScope = dataScopeAspect.NoDataScope
	userService.userDao.InsertUser(c, tx, sysUser)
	userService.insertUserPost(c, tx, sysUser.UserId, sysUser.PostIds)
	userService.insertUserRole(c, tx, sysUser.UserId, sysUser.RoleIds)
}

func (userService *UserService) UpdateUser(c *gin.Context, sysUser *systemModels.SysUserDML) {
	userId := sysUser.UserId
	tx, err := userService.data.Beginx()
	if err != nil {
		panic(err)
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // re-throw panic after Rollback
		} else {
			err = tx.Commit() // err is nil; if Commit returns error update err
		}
	}()
	userService.userPostDao.DeleteUserPostByUserId(c, tx, userId)
	userService.insertUserPost(c, tx, sysUser.UserId, sysUser.PostIds)
	userService.userRoleDao.DeleteUserRoleByUserId(c, tx, userId)
	userService.insertUserRole(c, tx, sysUser.UserId, sysUser.RoleIds)
	userService.userDao.UpdateUser(c, tx, sysUser)
}

func (userService *UserService) UpdateUserStatus(c *gin.Context, sysUser *systemModels.EditUserStatus) {
	s := new(systemModels.SysUserDML)
	s.UserId = sysUser.UserId
	s.Status = sysUser.Status
	s.BaseEntity = sysUser.BaseEntity
	userService.userDao.UpdateUser(c, userService.data, s)

}
func (userService *UserService) ResetPwd(c *gin.Context, userId int64, password string) {
	userService.userDao.ResetUserPwd(c, userService.data, userId, bCryptPasswordEncoder.HashPassword(password))

}

func (userService *UserService) insertUserPost(ctx context.Context, db sqly.SqlyContext, userId int64, posts []string) {
	if len(posts) != 0 {
		list := make([]*systemModels.SysUserPost, 0, len(posts))
		for _, postId := range posts {
			i, err := strconv.ParseInt(postId, 10, 64)
			if err != nil {
				panic(err)
			}
			post := systemModels.NewSysUserPost(userId, i)
			list = append(list, post)
		}
		userService.userPostDao.BatchUserPost(ctx, db, list)
	}

}

func (userService *UserService) insertUserRole(ctx context.Context, db sqly.SqlyContext, userId int64, roles []string) {
	if len(roles) != 0 {
		list := make([]*systemModels.SysUserRole, 0, len(roles))
		for _, roleId := range roles {
			i, err := strconv.ParseInt(roleId, 10, 64)
			if err != nil {
				panic(err)
			}
			role := systemModels.NewSysUserRole(userId, i)
			list = append(list, role)
		}
		userService.userRoleDao.BatchUserRole(ctx, db, list)
	}

}

func (userService *UserService) CheckUserNameUnique(c *gin.Context, userName string) bool {
	return userService.userDao.CheckUserNameUnique(c, userService.data, userName) > 0

}

func (userService *UserService) CheckPhoneUnique(c *gin.Context, id int64, phonenumber string) bool {
	if phonenumber == "" {
		return false
	}
	userId := userService.userDao.CheckPhoneUnique(c, userService.data, phonenumber)
	if userId == id || userId == 0 {
		return false
	}
	return true
}

func (userService *UserService) CheckEmailUnique(c *gin.Context, id int64, email string) bool {
	if email == "" {
		return false
	}
	userId := userService.userDao.CheckEmailUnique(c, userService.data, email)
	if userId == id || userId == 0 {
		return false
	}
	return true
}

func (userService *UserService) DeleteUserByIds(c *gin.Context, ids []int64) {
	tx, err := userService.data.Beginx()
	if err != nil {
		panic(err)
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else {
			err = tx.Commit()
		}
	}()
	userService.userRoleDao.DeleteUserRole(c, tx, ids)
	userService.userPostDao.DeleteUserPost(c, tx, ids)
	userService.userDao.DeleteUserByIds(c, tx, ids)

}

func (userService *UserService) UserImportData(c *gin.Context, fileHeader *multipart.FileHeader) (msg string, failureNum int) {
	file, err := fileHeader.Open()
	if err != nil {
		panic(err)
	}
	defer file.Close()
	excelFile, err := excelize.OpenReader(file)
	if err != nil {
		panic(err)
	}
	rows, err := excelFile.GetRows("Sheet1")
	if err != nil {
		panic(err)
	}

	deptSet := make(map[string]int64)
	userNameSet := baize.NewSet([]string{})
	for _, row := range rows {
		deptSet[(row[2])] = 0
		if userNameSet.Contains(row[0]) {
			failureNum++
			msg += "<br/>账号 " + row[0] + " 重复"
		} else {
			userNameSet.Add(row[0])
		}
	}
	dept := new(systemModels.SysDeptDQL)
	dept.DataScope = baizeContext.GetDataScope(c, "d")
	dl := userService.deptDao.SelectDeptList(c, userService.data, dept)
	ids := systemModels.GetParentNameAndIds(dl)
	list := make([]*systemModels.SysUserDML, 0)
	password := bCryptPasswordEncoder.HashPassword(userService.cs.SelectConfigValueByKey(c, "sys.account.initPassword"))
	list, msg, failureNum = systemModels.RowsToSysUserDMLList(rows, msg, failureNum, ids, password, baizeContext.GetUserId(c))
	names := userService.userDao.SelectUserNameByUserName(c, userService.data, userNameSet.ToSlice())
	for _, name := range names {
		failureNum++
		msg += "<br/>账号 " + name + " 已存在"
	}
	if failureNum > 0 {
		msg = "很抱歉，导入失败！共 " + strconv.Itoa(failureNum) + " 条数据格式不正确，错误如下：" + msg
		return msg, failureNum
	}
	userService.userDao.BatchInsertUser(c, userService.data, list)
	return "恭喜您，数据已全部导入成功！共 " + strconv.Itoa(len(list)) + " 条。", 0
}
func (userService *UserService) UpdateLoginInformation(c *gin.Context, userId int64, ip string) {
	userService.userDao.UpdateLoginInformation(c, userService.data, userId, ip)
}

func (userService *UserService) UpdateUserAvatar(c *gin.Context, file *multipart.FileHeader) string {
	userId := baizeContext.GetUserId(c)
	open, err := file.Open()
	defer open.Close()
	if err != nil {
		panic(err)
	}
	name := IOFile.GetRandomName(userId, filepath.Ext(file.Filename))
	avatar, err := IOFile.GetConfig().PublicUploadFile(c, open, name)
	if err != nil {
		panic(err)
	}
	userService.userDao.UpdateUserAvatar(c, userService.data, userId, avatar)
	baizeContext.GetSession(c).Set(c, sessionStatus.Avatar, avatar)
	return avatar
}

func (userService *UserService) ResetUserPwd(c *gin.Context, userId int64, password string) {
	userService.userDao.ResetUserPwd(c, userService.data, userId, bCryptPasswordEncoder.HashPassword(password))
}
func (userService *UserService) UpdateUserProfile(c *gin.Context, sysUser *systemModels.SysUserDML) {
	userService.userDao.UpdateUser(c, userService.data, sysUser)

}
func (userService *UserService) MatchesPassword(c *gin.Context, rawPassword string, userId int64) bool {

	return bCryptPasswordEncoder.CheckPasswordHash(rawPassword, userService.userDao.SelectPasswordByUserId(c, userService.data, userId))
}
func (userService *UserService) InsertUserAuth(c *gin.Context, userId int64, roleIds []int64) {
	userService.userRoleDao.DeleteUserRoleByUserId(c, userService.data, userId)
	if len(roleIds) != 0 {
		list := make([]*systemModels.SysUserRole, 0, len(roleIds))
		for _, roleId := range roleIds {
			role := systemModels.NewSysUserRole(userId, roleId)
			list = append(list, role)
		}
		userService.userRoleDao.BatchUserRole(c, userService.data, list)
	}
}
func (userService *UserService) GetUserAuthRole(c *gin.Context, userId int64) *systemModels.UserAndRoles {
	uar := new(systemModels.UserAndRoles)
	uar.User = userService.userDao.SelectUserById(c, userService.data, userId)
	s := new(systemModels.SysRoleDQL)
	if !baizeContext.IsAdmin(c) {
		s.CreateBy = baizeContext.GetUserId(c)
	}
	uar.Roles = userService.roleDao.SelectRoleAll(c, userService.data, s)
	rIds := userService.roleDao.SelectRoleListByUserId(c, userService.data, userId)
	uar.RoleIds = make([]string, 0, len(rIds))
	for _, id := range rIds {
		uar.RoleIds = append(uar.RoleIds, strconv.FormatInt(id, 10))
	}
	return uar
}

func (userService *UserService) GetUserProfile(c *gin.Context) *systemModels.UserProfile {
	userId := baizeContext.GetUserId(c)
	up := new(systemModels.UserProfile)
	up.User = userService.userDao.SelectUserById(c, userService.data, userId)
	roles := userService.roleDao.SelectBasicRolesByUserId(c, userService.data, userId)
	roleNames := make([]string, 0, len(roles))
	for _, role := range roles {
		roleNames = append(roleNames, role.RoleName)
	}
	up.RoleGroup = strings.Join(roleNames, ",")
	up.PostGroup = strings.Join(userService.postDao.SelectPostNameListByUserId(c, userService.data, userId), ",")

	return up
}

func (userService *UserService) UpdateUserDataScope(c *gin.Context, uds *systemModels.SysUserDataScope) {
	userId := uds.UserId
	su := new(systemModels.SysUserDML)
	su.SetUpdateBy(baizeContext.GetUserId(c))
	su.UserId = userId
	su.DataScope = uds.DataScope
	tx, err := userService.data.Beginx()
	if err != nil {
		panic(err)
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p) // re-throw panic after Rollback
		} else {
			err = tx.Commit() // err is nil; if Commit returns error update err
		}
	}()
	userService.uds.DeleteUserDeptScopeByUserId(c, tx, userId)
	userService.userDao.UpdateUser(c, tx, su)
	if uds.DataScope == dataScopeAspect.DataScopeCustom && len(uds.DeptIds) != 0 {
		scopes := make([]*systemModels.SysUserDeptScope, 0, len(uds.DeptIds))
		for _, id := range uds.DeptIds {
			i, err := strconv.ParseInt(id, 10, 64)
			if err != nil {
				panic(err)
			}
			scopes = append(scopes, &systemModels.SysUserDeptScope{UserId: userId, DeptId: i})
		}
		userService.uds.BatchUserDeptScope(c, tx, scopes)
	}
}

func (userService *UserService) SelectUserDataScope(c *gin.Context, userId int64) *systemModels.SysUserDataScope {
	s := new(systemModels.SysUserDataScope)
	s.UserId = userId
	s.DataScope = userService.userDao.SelectUserById(c, userService.data, userId).DataScope
	if s.DataScope == dataScopeAspect.DataScopeCustom {
		s.DeptIds = userService.uds.SelectUserDeptScopeDeptIdByUserId(c, userService.data, userId)
	} else {
		s.DeptIds = make([]string, 0)
	}
	return s
}
