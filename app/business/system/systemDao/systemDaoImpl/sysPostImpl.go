package systemDaoImpl

import (
	"baize/app/business/system/systemDao"
	"baize/app/business/system/systemModels"
	"context"
	"database/sql"
	"errors"

	"github.com/baizeplus/sqly"
)

type sysPostDao struct {
	ms      sqly.SqlyContext
	postSql string
}

func NewSysPostDao(ms sqly.SqlyContext) systemDao.IPostDao {
	return &sysPostDao{
		ms:      ms,
		postSql: `select post_id, post_code, post_name, post_sort, status, create_by, create_time, remark  from sys_post`,
	}
}

func (postDao *sysPostDao) SelectPostAll(ctx context.Context) (sysPost []*systemModels.SysPostVo) {
	sysPost = make([]*systemModels.SysPostVo, 0)
	err := postDao.ms.SelectContext(ctx, &sysPost, postDao.postSql+" order by post_sort")
	if err != nil {
		panic(err)
	}
	return
}

func (postDao *sysPostDao) SelectPostListByUserId(ctx context.Context, userId string) (list []string) {
	sqlStr := `select p.post_id
		from sys_post p
		left join sys_user_post up on up.post_id = p.post_id
		left join sys_user u on u.user_id = up.user_id
		where u.user_id = ?`
	list = make([]string, 0)
	err := postDao.ms.SelectContext(ctx, &list, sqlStr, userId)
	if err != nil {
		panic(err)
	}
	return
}

func (postDao *sysPostDao) SelectPostList(ctx context.Context, post *systemModels.SysPostDQL) (list []*systemModels.SysPostVo, total int64) {
	if post.OrderBy == "" {
		post.OrderBy = "post_sort"
	}
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
	err := postDao.ms.NamedSelectPageContext(ctx, &list, &total, postDao.postSql+whereSql, post)
	if err != nil {
		panic(err)
	}
	return
}
func (postDao *sysPostDao) SelectPostListAll(ctx context.Context, post *systemModels.SysPostDQL) (list []*systemModels.SysPostVo) {
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
	list = make([]*systemModels.SysPostVo, 0)
	err := postDao.ms.NamedSelectContext(ctx, &list, postDao.postSql+whereSql, post)
	if err != nil {
		panic(err)
	}
	return
}

func (postDao *sysPostDao) SelectPostById(ctx context.Context, postId string) (dictData *systemModels.SysPostVo) {

	dictData = new(systemModels.SysPostVo)
	err := postDao.ms.GetContext(ctx, dictData, postDao.postSql+" where post_id = ?", postId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		panic(err)
	}
	return
}

func (postDao *sysPostDao) InsertPost(ctx context.Context, post *systemModels.SysPostVo) {
	insertSQL := `insert into sys_post(post_id,post_code,post_name,post_sort,status,remark,create_by,create_time,update_by,update_time )
					values(:post_id,:post_code,:post_name,:post_sort,:status,:remark,:create_by,:create_time,:update_by,:update_time )`
	_, err := postDao.ms.NamedExecContext(ctx, insertSQL, post)
	if err != nil {
		panic(err)
	}
	return
}

func (postDao *sysPostDao) UpdatePost(ctx context.Context, post *systemModels.SysPostVo) {
	updateSQL := `update sys_post set update_time = :update_time , update_by = :update_by`

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

	_, err := postDao.ms.NamedExecContext(ctx, updateSQL, post)
	if err != nil {
		panic(err)
	}
	return
}

func (postDao *sysPostDao) DeletePostByIds(ctx context.Context, postId []string) {
	query, i, err := sqly.In("delete from sys_post where post_id in (?)", postId)
	if err != nil {
		panic(err)
	}
	_, err = postDao.ms.ExecContext(ctx, query, i...)
	if err != nil {
		panic(err)
	}
	return
}
func (postDao *sysPostDao) SelectPostNameListByUserId(ctx context.Context, userId string) (list []string) {
	sqlStr := `select p.post_name
		from sys_post p
		left join sys_user_post up on up.post_id = p.post_id
		left join sys_user u on u.user_id = up.user_id
		where u.user_id = ?`
	list = make([]string, 0, 1)
	err := postDao.ms.SelectContext(ctx, &list, sqlStr, userId)
	if err != nil {
		panic(err)
	}
	return
}
