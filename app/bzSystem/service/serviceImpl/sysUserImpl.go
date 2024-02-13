package serviceImpl

import (
	"baize/app/bzSystem/dao"
	"baize/app/bzSystem/dao/daoImpl"
	"baize/app/bzSystem/models"
	"baize/app/utils/IOFile"
	"baize/app/utils/bCryptPasswordEncoder"
	"baize/app/utils/baizeContext"
	"baize/app/utils/snowflake"
	"context"
	"github.com/baizeplus/sqly"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/util/gconv"
	"mime/multipart"
	"strconv"
)

type UserService struct {
	data        *sqly.DB
	userDao     dao.IUserDao
	userPostDao dao.IUserPostDao
	userRoleDao dao.IUserRoleDao
	roleDao     dao.IRoleDao
	postDao     dao.IPostDao
}

func NewUserService(data *sqly.DB, ud *daoImpl.SysUserDao, upd *daoImpl.SysUserPostDao, urd *daoImpl.SysUserRoleDao, rd *daoImpl.SysRoleDao, pd *daoImpl.SysPostDao) *UserService {
	return &UserService{
		data:        data,
		userDao:     ud,
		userPostDao: upd,
		userRoleDao: urd,
		roleDao:     rd,
		postDao:     pd,
	}
}

func (userService *UserService) SelectUserByUserName(c *gin.Context, userName string) *models.User {
	return userService.userDao.SelectUserByUserName(c, userService.data, userName)

}
func (userService *UserService) SelectUserList(c *gin.Context, user *models.SysUserDQL) (sysUserList []*models.SysUserVo, count *int64) {
	return userService.userDao.SelectUserList(c, userService.data, user)
}

//func (userService *UserService) UserExport(user *models.SysUserDQL) (data []byte) {
//	sysUserList, _ := userService.userDao.SelectUserList(userService.data.GetSlaveDb(), user)
//	return exceLize.SetRows(models.SysUserListToRows(sysUserList))
//}
//func (userService *UserService) ImportTemplate() (data []byte) {
//f := excelize.NewFile()
//template := models.SysUserImportTemplate()
//f.SetSheetRow("Sheet1", "A1", &template)
//buffer, _ := f.WriteToBuffer()
//return buffer.Bytes()

//}

func (userService *UserService) SelectUserById(c *gin.Context, userId int64) (sysUser *models.SysUserVo) {
	return userService.userDao.SelectUserById(c, userService.data, userId)

}
func (userService *UserService) SelectUserAndAccreditById(c *gin.Context, userId int64) (sysUser *models.UserAndAccredit) {
	uaa := new(models.UserAndAccredit)
	uaa.User = userService.userDao.SelectUserById(c, userService.data, userId)
	uaa.Roles = userService.roleDao.SelectRoleAll(c, userService.data)
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
func (userService *UserService) SelectAccredit(c *gin.Context) (sysUser *models.Accredit) {
	ua := new(models.Accredit)
	ua.Roles = userService.roleDao.SelectRoleAll(c, userService.data)
	ua.Posts = userService.postDao.SelectPostAll(c, userService.data)
	return ua
}

func (userService *UserService) InsertUser(c *gin.Context, sysUser *models.SysUserDML) {
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
	userService.userDao.InsertUser(c, tx, sysUser)
	userService.insertUserPost(c, tx, sysUser.UserId, sysUser.PostIds)
	userService.insertUserRole(c, tx, sysUser.UserId, sysUser.RoleIds)

}

func (userService *UserService) UpdateUser(c *gin.Context, sysUser *models.SysUserDML) {
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

func (userService *UserService) UpdateUserStatus(c *gin.Context, sysUser *models.EditUserStatus) {
	s := new(models.SysUserDML)
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
		list := make([]*models.SysUserPost, 0, len(posts))
		for _, postId := range posts {
			post := models.NewSysUserPost(userId, gconv.Int64(postId))
			list = append(list, post)
		}
		userService.userPostDao.BatchUserPost(ctx, db, list)
	}

}

func (userService *UserService) insertUserRole(ctx context.Context, db sqly.SqlyContext, userId int64, roles []string) {
	if len(roles) != 0 {
		list := make([]*models.SysUserRole, 0, len(roles))
		for _, roleId := range roles {
			role := models.NewSysUserRole(userId, gconv.Int64(roleId))
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

func (userService *UserService) UserImportData(c *gin.Context, rows [][]string, userId int64, deptId int64) (msg string, failureNum int) {
	successNum := 0
	list, failureMsg, failureNum := models.RowsToSysUserDMLList(rows)
	password := bCryptPasswordEncoder.HashPassword("123456")
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
	for _, user := range list {
		unique := userService.userDao.CheckUserNameUnique(c, tx, user.UserName)
		if unique < 1 {
			user.DeptId = deptId
			user.Password = password
			//user.SetCreateBy(userId)
			userService.userDao.InsertUser(c, tx, user)
			successNum++
		} else {
			failureNum++
			failureMsg += "<br/>账号 " + user.UserName + " 已存在"
		}
	}
	if failureNum > 0 {
		failureMsg = "很抱歉，导入失败！共 " + strconv.Itoa(failureNum) + " 条数据格式不正确，错误如下：" + failureMsg
		return failureMsg, failureNum
	}
	return "恭喜您，数据已全部导入成功！共 " + strconv.Itoa(successNum) + " 条。", 0
}
func (userService *UserService) UpdateLoginInformation(c *gin.Context, userId int64, ip string) {
	userService.userDao.UpdateLoginInformation(c, userService.data, userId, ip)
}

func (userService *UserService) UpdateUserAvatar(c *gin.Context, file *multipart.FileHeader) string {
	userId := baizeContext.GetUserId(c)
	open, err := file.Open()
	if err != nil {
		panic(err)
	}
	avatar, err := IOFile.GetConfig().PublicUploadFile(IOFile.NewFileParamsRandomName(IOFile.GetTenantRandomName(userId, IOFile.GetExtension(file)), open))
	if err != nil {
		panic(err)
	}
	userService.userDao.UpdateUserAvatar(c, userService.data, userId, avatar)
	return avatar
}

func (userService *UserService) ResetUserPwd(c *gin.Context, userId int64, password string) {
	userService.userDao.ResetUserPwd(c, userService.data, userId, bCryptPasswordEncoder.HashPassword(password))
}
func (userService *UserService) UpdateUserProfile(c *gin.Context, sysUser *models.SysUserDML) {
	userService.userDao.UpdateUser(c, userService.data, sysUser)

}
func (userService *UserService) MatchesPassword(c *gin.Context, rawPassword string, userId int64) bool {

	return bCryptPasswordEncoder.CheckPasswordHash(rawPassword, userService.userDao.SelectPasswordByUserId(c, userService.data, userId))
}
func (userService *UserService) InsertUserAuth(c *gin.Context, userId int64, roleIds []int64) {
	userService.userRoleDao.DeleteUserRoleByUserId(c, userService.data, userId)
	if len(roleIds) != 0 {
		list := make([]*models.SysUserRole, 0, len(roleIds))
		for _, roleId := range roleIds {
			role := models.NewSysUserRole(userId, roleId)
			list = append(list, role)
		}
		userService.userRoleDao.BatchUserRole(c, userService.data, list)
	}
}
func (userService *UserService) GetUserAuthRole(c *gin.Context, userId int64) *models.UserAndRoles {
	uar := new(models.UserAndRoles)
	uar.User = userService.userDao.SelectUserById(c, userService.data, userId)
	uar.Roles = userService.roleDao.SelectRoleAll(c, userService.data)
	rIds := userService.roleDao.SelectRoleListByUserId(c, userService.data, userId)
	uar.RoleIds = make([]string, 0, len(rIds))
	for _, id := range rIds {
		uar.RoleIds = append(uar.RoleIds, strconv.FormatInt(id, 10))
	}
	return uar
}
