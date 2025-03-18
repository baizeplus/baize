package systemServiceImpl

import (
	"baize/app/business/system/systemDao"
	"baize/app/business/system/systemModels"
	"baize/app/business/system/systemService"
	"baize/app/utils/snowflake"
	"context"
	"github.com/gin-gonic/gin"
)

type PermissionService struct {
	pd systemDao.IPermissionDao
}

func NewPermissionService(pd systemDao.IPermissionDao) systemService.ISysPermissionService {
	return &PermissionService{pd: pd}
}
func (ps *PermissionService) SelectPermissionList(c *gin.Context, permission *systemModels.SysPermissionDQL) (list []*systemModels.SysPermissionVo) {

	list = ps.pd.SelectPermissionList(c, permission)
	return
}

func (ps *PermissionService) SelectPermissionById(ctx context.Context, permissionId int64) (Permission *systemModels.SysPermissionVo) {
	return ps.pd.SelectPermissionById(ctx, permissionId)
}

func (ps *PermissionService) SelectPermissionListByRoleIds(ctx context.Context, roleIds []int64) (list []*systemModels.SysPermissionVo) {
	return ps.pd.SelectPermissionListByRoleIds(ctx, roleIds)
}

func (ps *PermissionService) InsertPermission(ctx context.Context, permission *systemModels.SysPermissionAdd) {
	permission.PermissionId = snowflake.GenID()
	permission.Status = "0"
	ps.pd.InsertPermission(ctx, permission)
}

func (ps *PermissionService) UpdatePermission(ctx context.Context, permission *systemModels.SysPermissionEdit) {
	ps.pd.UpdatePermission(ctx, permission)
}

func (ps *PermissionService) DeletePermissionById(ctx context.Context, permissionId int64) {
	ps.pd.DeletePermissionById(ctx, permissionId)
}

func (ps *PermissionService) HasChildByPermissionId(ctx context.Context, permissionId int64) bool {
	return ps.pd.HasChildByPermissionId(ctx, permissionId) > 0
}

//func (ps *PermissionService) CheckPermissionExistRole(ctx context.Context, permissionId int64) bool {
//	return ps.rd.CheckPermissionExistRole(ctx, permissionId) > 0
//}
