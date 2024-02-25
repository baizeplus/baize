package systemDaoImpl

import (
	"baize/app/business/system/systemModels"
	"context"
	"github.com/baizeplus/sqly"
)

type SysUserPostDao struct {
}

func NewSysUserPostDao() *SysUserPostDao {
	return &SysUserPostDao{}
}

func (sysUserPostDao *SysUserPostDao) BatchUserPost(ctx context.Context, db sqly.SqlyContext, users []*systemModels.SysUserPost) {

	_, err := db.NamedExecContext(ctx, "insert into sys_user_post(user_id, post_id) values (:user_id,:post_id)", users)
	if err != nil {
		panic(err)
	}
}

func (sysUserPostDao *SysUserPostDao) DeleteUserPostByUserId(ctx context.Context, db sqly.SqlyContext, userId int64) {

	_, err := db.ExecContext(ctx, "delete from sys_user_post where user_id= ?", userId)
	if err != nil {
		panic(err)
	}
}

func (sysUserPostDao *SysUserPostDao) DeleteUserPost(ctx context.Context, db sqly.SqlyContext, ids []int64) {
	query, i, err := sqly.In("delete from sys_user_post where user_id in(?)", ids)
	if err != nil {
		panic(err)
	}
	_, err = db.ExecContext(ctx, query, i...)
	if err != nil {
		panic(err)
	}
}
