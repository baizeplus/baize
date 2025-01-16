package systemDaoImpl

import (
	"baize/app/business/system/systemDao"
	"baize/app/business/system/systemModels"
	"context"
	"github.com/baizeplus/sqly"
)

type sysUserPostDao struct {
	ms sqly.SqlyContext
}

func NewSysUserPostDao(ms sqly.SqlyContext) systemDao.IUserPostDao {
	return &sysUserPostDao{ms: ms}
}

func (sysUserPostDao *sysUserPostDao) BatchUserPost(ctx context.Context, users []*systemModels.SysUserPost) {

	_, err := sysUserPostDao.ms.NamedExecContext(ctx, "insert into sys_user_post(user_id, post_id) values (:user_id,:post_id)", users)
	if err != nil {
		panic(err)
	}
}

func (sysUserPostDao *sysUserPostDao) DeleteUserPostByUserId(ctx context.Context, userId int64) {

	_, err := sysUserPostDao.ms.ExecContext(ctx, "delete from sys_user_post where user_id= ?", userId)
	if err != nil {
		panic(err)
	}
}

func (sysUserPostDao *sysUserPostDao) DeleteUserPost(ctx context.Context, ids []int64) {
	query, i, err := sqly.In("delete from sys_user_post where user_id in(?)", ids)
	if err != nil {
		panic(err)
	}
	_, err = sysUserPostDao.ms.ExecContext(ctx, query, i...)
	if err != nil {
		panic(err)
	}
}
