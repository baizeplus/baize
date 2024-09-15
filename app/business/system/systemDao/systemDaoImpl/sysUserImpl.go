package systemDaoImpl

import (
	"baize/app/business/system/systemModels"
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/baizeplus/sqly"
)

type SysUserDao struct {
}

func NewSysUserDao() *SysUserDao {
	return &SysUserDao{}
}

func (userDao *SysUserDao) SelectUserNameByUserName(ctx context.Context, db sqly.SqlyContext, userName []string) []string {
	query, i, err := sqly.In("select user_name from sys_user where user_name in(?)", userName)
	if err != nil {
		panic(err)
	}
	list := make([]string, 0)
	err = db.SelectContext(ctx, &list, query, i...)
	if err != nil {
		panic(err)
	}
	return list
}
func (userDao *SysUserDao) CheckUserNameUnique(ctx context.Context, db sqly.SqlyContext, userName string) int {
	var count = 0
	err := db.GetContext(ctx, &count, "SELECT EXISTS( SELECT 1 FROM sys_user WHERE user_name =?)", userName)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		panic(err)
	}
	return count
}
func (userDao *SysUserDao) CheckPhoneUnique(ctx context.Context, db sqly.SqlyContext, phonenumber string) int64 {
	var userId int64 = 0
	err := db.GetContext(ctx, &userId, "select user_id from sys_user where phonenumber = ?", phonenumber)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		panic(err)
	}
	return userId
}

func (userDao *SysUserDao) CheckEmailUnique(ctx context.Context, db sqly.SqlyContext, email string) int64 {
	var userId int64 = 0
	err := db.GetContext(ctx, &userId, "select user_id from sys_user where email = ?", email)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		panic(err)
	}
	return userId
}

func (userDao *SysUserDao) InsertUser(ctx context.Context, db sqly.SqlyContext, sysUser *systemModels.SysUserDML) {
	insertSQL := `insert into sys_user(user_id,user_name,nick_name,sex,password,data_scope,status,create_by,create_time,update_by,update_time %s)
					values(:user_id,:user_name,:nick_name,:sex,:password,:data_scope,:status,:create_by,now(),:update_by,now() %s)`
	key := ""
	value := ""
	if sysUser.DeptId != 0 {
		key += ",dept_id"
		value += ",:dept_id"
	}
	if sysUser.Email != "" {
		key += ",email"
		value += ",:email"
	}
	if sysUser.Avatar != "" {
		key += ",avatar"
		value += ",:avatar"
	}
	if sysUser.Phonenumber != "" {
		key += ",phonenumber"
		value += ",:phonenumber"
	}
	if sysUser.Remark != "" {
		key += ",remark"
		value += ",:remark"
	}
	insertStr := fmt.Sprintf(insertSQL, key, value)

	_, err := db.NamedExecContext(ctx, insertStr, sysUser)

	if err != nil {
		panic(err)
	}
}
func (userDao *SysUserDao) BatchInsertUser(ctx context.Context, db sqly.SqlyContext, sysUser []*systemModels.SysUserDML) {
	insertSQL := `insert into sys_user(user_id,user_name,nick_name,email,phonenumber,sex,password,data_scope,status,dept_id,create_by,create_time,update_by,update_time)
					values(:user_id,:user_name,:nick_name,:email,:phonenumber,:sex,:password,:data_scope,:status,:dept_id,:create_by,now(),:update_by,now())`
	_, err := db.NamedExecContext(ctx, insertSQL, sysUser)

	if err != nil {
		panic(err)
	}
}

func (userDao *SysUserDao) UpdateUser(ctx context.Context, db sqly.SqlyContext, sysUser *systemModels.SysUserDML) {
	updateSQL := `update sys_user set update_time = now() , update_by = :update_by`

	if sysUser.Email != "" {
		updateSQL += ",email = :email"
	}
	if sysUser.DeptId != 0 {
		updateSQL += ",dept_id = :dept_id"
	}

	if sysUser.DataScope != "" {
		updateSQL += ",data_scope = :data_scope"
	}
	if sysUser.Avatar != "" {
		updateSQL += ",avatar = :avatar"
	}
	if sysUser.Phonenumber != "" {
		updateSQL += ",phonenumber = :phonenumber"
	}
	if sysUser.Remark != "" {
		updateSQL += ",remark = :remark"
	}
	if sysUser.NickName != "" {
		updateSQL += ",nick_name = :nick_name"
	}
	if sysUser.Sex != "" {
		updateSQL += ",sex = :sex"
	}
	if sysUser.Status != "" {
		updateSQL += ",status = :status"
	}
	updateSQL += " where user_id = :user_id"

	_, err := db.NamedExecContext(ctx, updateSQL, sysUser)
	if err != nil {
		panic(err)
	}
}

func (userDao *SysUserDao) SelectUserByUserName(ctx context.Context, db sqly.SqlyContext, userName string) (loginUser *systemModels.User) {
	sqlStr := `select u.user_id, u.dept_id, u.user_name,  u.avatar, u.password, u.status, u.del_flag,u.data_scope
        from sys_user u
		where u.user_name = ?			
			`

	loginUser = new(systemModels.User)
	err := db.GetContext(ctx, loginUser, sqlStr, userName)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		panic(err)
	}
	return
}
func (userDao *SysUserDao) SelectUserById(ctx context.Context, db sqly.SqlyContext, userId int64) (sysUser *systemModels.SysUserVo) {
	sqlStr := `select u.user_id, u.dept_id, u.nick_name, u.user_name, u.email, u.avatar, u.phonenumber, u.sex, u.status, u.del_flag,  u.create_by, u.create_time, u.remark, d.dept_name, d.leader,  u.data_scope,  r.role_id
        from sys_user u
		    left join sys_dept d on u.dept_id = d.dept_id
		    left join sys_user_role ur on u.user_id = ur.user_id
		    left join sys_role r on r.role_id = ur.role_id		
			where u.user_id = ?
			`

	sysUser = new(systemModels.SysUserVo)
	err := db.GetContext(ctx, sysUser, sqlStr, userId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		panic(err)
	}
	return
}

func (userDao *SysUserDao) SelectUserList(ctx context.Context, db sqly.SqlyContext, user *systemModels.SysUserDQL) (list []*systemModels.SysUserVo, total int64) {
	sql := `select u.user_id, u.dept_id, u.nick_name, u.user_name, u.email, u.avatar, u.phonenumber, u.sex, u.status, u.del_flag, u.create_by, u.create_time, u.remark, d.dept_name, d.leader
			 from sys_user u left join sys_dept d on u.dept_id = d.dept_id where u.del_flag = '0'`
	if user.UserName != "" {
		sql += " AND u.user_name like concat('%', :user_name, '%')"
	}
	if user.Status != "" {
		sql += " AND  u.status = :status"
	}
	if user.Phonenumber != "" {
		sql += " AND u.phonenumber like concat('%', :phonenumber, '%')"
	}
	if user.BeginTime != "" {
		sql += " AND date_format(u.create_time,'%y%m%d') >= date_format(:begin_time,'%y%m%d')"
	}
	if user.EndTime != "" {
		sql += " AND date_format(u.create_time,'%y%m%d') <= date_format(:end_time,'%y%m%d')"
	}
	if user.DeptId != 0 {
		sql += " AND (u.dept_id = :dept_id OR u.dept_id IN ( SELECT t.dept_id FROM sys_dept t WHERE find_in_set(:dept_id, ancestors) ))"
	}
	if user.DataScope != "" {
		sql += " AND " + user.DataScope
	}
	err := db.NamedSelectPageContext(ctx, &list, &total, sql, user)
	if err != nil {
		panic(err)
	}
	return

}

func (userDao *SysUserDao) SelectUserListAll(ctx context.Context, db sqly.SqlyContext, user *systemModels.SysUserDQL) (list []*systemModels.SysUserVo) {
	sql := `select u.user_id, u.dept_id, u.nick_name, u.user_name, u.email, u.avatar, u.phonenumber, u.sex, u.status, u.del_flag, u.create_by, u.create_time, u.remark, d.dept_name, d.leader
			 from sys_user u left join sys_dept d on u.dept_id = d.dept_id where u.del_flag = '0'`
	if user.UserName != "" {
		sql += " AND u.user_name like concat('%', :user_name, '%')"
	}
	if user.Status != "" {
		sql += " AND  u.status = :status"
	}
	if user.Phonenumber != "" {
		sql += " AND u.phonenumber like concat('%', :phonenumber, '%')"
	}
	if user.BeginTime != "" {
		sql += " AND date_format(u.create_time,'%y%m%d') >= date_format(:begin_time,'%y%m%d')"
	}
	if user.EndTime != "" {
		sql += " AND date_format(u.create_time,'%y%m%d') <= date_format(:end_time,'%y%m%d')"
	}
	if user.DeptId != 0 {
		sql += " AND (u.dept_id = :dept_id OR u.dept_id IN ( SELECT t.dept_id FROM sys_dept t WHERE find_in_set(:dept_id, ancestors) ))"
	}

	err := db.NamedSelectContext(ctx, &list, sql, user)
	if err != nil {
		panic(err)
	}
	return

}

func (userDao *SysUserDao) DeleteUserByIds(ctx context.Context, db sqly.SqlyContext, ids []int64) {
	query, i, err := sqly.In("update sys_user set del_flag = '2' where user_id in(?)", ids)
	if err != nil {
		panic(err)
	}
	_, err = db.ExecContext(ctx, query, i...)
	if err != nil {
		panic(err)
	}
	return
}
func (userDao *SysUserDao) UpdateLoginInformation(ctx context.Context, db sqly.SqlyContext, userId int64, ip string) {
	_, err := db.ExecContext(ctx, `update sys_user set login_date = now() , login_ip = ?  where user_id = ?`, ip, userId)
	if err != nil {
		panic(err)
	}
}
func (userDao *SysUserDao) UpdateUserAvatar(ctx context.Context, db sqly.SqlyContext, userId int64, avatar string) {
	_, err := db.ExecContext(ctx, `update sys_user set avatar = ?  where user_id = ?`, avatar, userId)
	if err != nil {
		panic(err)
	}
}
func (userDao *SysUserDao) ResetUserPwd(ctx context.Context, db sqly.SqlyContext, userId int64, password string) {
	_, err := db.ExecContext(ctx, `update sys_user set password = ?  where user_id = ?`, password, userId)
	if err != nil {
		panic(err)
	}
}
func (userDao *SysUserDao) SelectPasswordByUserId(ctx context.Context, db sqly.SqlyContext, userId int64) string {
	sqlStr := `select password
        from sys_user 
			where user_id = ?			
			`

	password := new(string)
	err := db.GetContext(ctx, password, sqlStr, userId)
	if err != nil {
		panic(err)
	}
	return *password
}

func (userDao *SysUserDao) SelectUserIdsByDeptIds(ctx context.Context, db sqly.SqlyContext, deptIds []int64) []int64 {
	query, i, err := sqly.In("select user_id from sys_user where dept_id in(?)", deptIds)
	if err != nil {
		panic(err)
	}
	list := make([]int64, 0)
	err = db.SelectContext(ctx, &list, query, i...)
	if err != nil {
		panic(err)
	}
	return list
}
