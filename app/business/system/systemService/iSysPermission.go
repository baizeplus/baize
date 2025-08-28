package systemService

import (
	"baize/app/business/system/systemModels"
	"context"

	"github.com/gin-gonic/gin"
)

type ISysPermissionService interface {
	SelectPermissionList(c *gin.Context, permission *systemModels.SysPermissionDQL) (list []*systemModels.SysPermissionVo)
	SelectPermissionById(ctx context.Context, permissionId string) (Permission *systemModels.SysPermissionVo)
	InsertPermission(ctx context.Context, permission *systemModels.SysPermissionAdd)
	UpdatePermission(ctx context.Context, permission *systemModels.SysPermissionEdit)
	DeletePermissionById(ctx context.Context, permissionId string)
	HasChildByPermissionId(ctx context.Context, permissionId string) bool
	//CheckPermissionExistRole(ctx context.Context, permissionId int64) bool
	SelectPermissionListByRoleIds(ctx context.Context, roleIds []string) (list []*systemModels.SysPermissionVo)
}
