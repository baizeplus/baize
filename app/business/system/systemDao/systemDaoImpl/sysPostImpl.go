package systemDaoImpl

import (
	"baize/app/business/system/systemModels"
	"context"
	"database/sql"
	"errors"
	"github.com/baizeplus/sqly"
)

type SysPostDao struct {
	postSql string
}

func NewSysPostDao() *SysPostDao {
	return &SysPostDao{
		postSql: `select post_id, post_code, post_name, post_sort, status, create_by, create_time, remark  from sys_post`,
	}
}

func (postDao *SysPostDao) SelectPostAll(ctx context.Context, db sqly.SqlyContext) (sysPost []*systemModels.SysPostVo) {
	sysPost = make([]*systemModels.SysPostVo, 0)
	err := db.SelectContext(ctx, &sysPost, postDao.postSql)
	if err != nil {
		panic(err)
	}
	return
}

func (postDao *SysPostDao) SelectPostListByUserId(ctx context.Context, db sqly.SqlyContext, userId int64) (list []int64) {
	sqlStr := `select p.post_id
		from sys_post p
		left join sys_user_post up on up.post_id = p.post_id
		left join sys_user u on u.user_id = up.user_id
		where u.user_id = ?`
	list = make([]int64, 0, 1)
	err := db.SelectContext(ctx, &list, sqlStr, userId)
	if err != nil {
		panic(err)
	}
	return
}

func (postDao *SysPostDao) SelectPostList(ctx context.Context, db sqly.SqlyContext, post *systemModels.SysPostDQL) (list []*systemModels.SysPostVo, total int64) {
	whereSql := ``
	if post.PostCode != "" {
		whereSql += " AND post_code like concat('%', :post_code, '%')"
	}
	if post.Status != "" {
		whereSql += " AND  status = :status"
	}
	if post.PostName != "" {
		whereSql += " AND post_name like concat('%', :post_name, '%')"
	}

	if whereSql != "" {
		whereSql = " where " + whereSql[4:]
	}
	err := db.NamedSelectPageContext(ctx, &list, &total, postDao.postSql+whereSql, post)
	if err != nil {
		panic(err)
	}
	return
}
func (postDao *SysPostDao) SelectPostListAll(ctx context.Context, db sqly.SqlyContext, post *systemModels.SysPostDQL) (list []*systemModels.SysPostVo) {
	whereSql := ``
	if post.PostCode != "" {
		whereSql += " AND post_code like concat('%', :post_code, '%')"
	}
	if post.Status != "" {
		whereSql += " AND  status = :status"
	}
	if post.PostName != "" {
		whereSql += " AND post_name like concat('%', :post_name, '%')"
	}

	if whereSql != "" {
		whereSql = " where " + whereSql[4:]
	}
	list = make([]*systemModels.SysPostVo, 0, 16)
	err := db.NamedSelectContext(ctx, &list, postDao.postSql+whereSql, post)
	if err != nil {
		panic(err)
	}
	return
}

func (postDao *SysPostDao) SelectPostById(ctx context.Context, db sqly.SqlyContext, postId int64) (dictData *systemModels.SysPostVo) {

	dictData = new(systemModels.SysPostVo)
	err := db.GetContext(ctx, dictData, postDao.postSql+" where post_id = ?", postId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		panic(err)
	}
	return
}

func (postDao *SysPostDao) InsertPost(ctx context.Context, db sqly.SqlyContext, post *systemModels.SysPostVo) {
	insertSQL := `insert into sys_post(post_id,post_code,post_name,post_sort,status,remark,create_by,create_time,update_by,update_time )
					values(:post_id,:post_code,:post_name,:post_sort,:status,:remark,:create_by,now(),:update_by,now() )`
	_, err := db.NamedExecContext(ctx, insertSQL, post)
	if err != nil {
		panic(err)
	}
	return
}

func (postDao *SysPostDao) UpdatePost(ctx context.Context, db sqly.SqlyContext, post *systemModels.SysPostVo) {
	updateSQL := `update sys_post set update_time = now() , update_by = :update_by`

	if post.PostCode != "" {
		updateSQL += ",post_code = :post_code"
	}

	if post.PostName != "" {
		updateSQL += ",post_name = :post_name"
	}
	if post.PostSort != 0 {
		updateSQL += ",post_sort = :post_sort"
	}
	if post.Status != "" {
		updateSQL += ",status = :status"
	}
	if post.Status != "" {
		updateSQL += ",remark = :remark"
	}

	updateSQL += " where post_id = :post_id"

	_, err := db.NamedExecContext(ctx, updateSQL, post)
	if err != nil {
		panic(err)
	}
	return
}

func (postDao *SysPostDao) DeletePostByIds(ctx context.Context, db sqly.SqlyContext, dictCodes []int64) {
	query, i, err := sqly.In("delete from sys_post where post_id in (?)", dictCodes)
	if err != nil {
		panic(err)
	}
	_, err = db.ExecContext(ctx, query, i...)
	if err != nil {
		panic(err)
	}
	return
}
func (postDao *SysPostDao) SelectPostNameListByUserId(ctx context.Context, db sqly.SqlyContext, userId int64) (list []string) {
	sqlStr := `select p.post_name
		from sys_post p
		left join sys_user_post up on up.post_id = p.post_id
		left join sys_user u on u.user_id = up.user_id
		where u.user_id = ?`
	list = make([]string, 0, 1)
	err := db.SelectContext(ctx, &list, sqlStr, userId)
	if err != nil {
		panic(err)
	}
	return
}
