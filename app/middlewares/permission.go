package middlewares

import (
	"baize/app/utils/baizeContext"
	"github.com/gin-gonic/gin"
)

func HasPermissions(permissionSlice []string) func(c *gin.Context) {
	return func(c *gin.Context) {
		permissions := baizeContext.GetPermission(c)
		if permissions == nil || len(permissions) == 0 {
			baizeContext.PermissionDenied(c)
			c.Abort()
			return
		}
		if !hasPermissionsSlice(permissions, permissionSlice) {
			baizeContext.PermissionDenied(c)
			c.Abort()
			return
		}
		c.Next()
	}
}

func HasPermission(permission string) func(c *gin.Context) {
	return func(c *gin.Context) {
		permissions := baizeContext.GetPermission(c)
		if permissions == nil || len(permissions) == 0 {
			baizeContext.PermissionDenied(c)
			c.Abort()
			return
		}
		if !hasPermissions(permissions, permission) {
			baizeContext.PermissionDenied(c)
			c.Abort()
			return
		}
		c.Next()
	}
}

func hasPermissions(permissions []string, permission string) bool {
	for _, item := range permissions {
		if item == permission {
			return true
		}
	}
	return false
}
func hasPermissionsSlice(permissions []string, p []string) bool {
	for _, item := range permissions {
		for _, s := range p {
			if item == s {
				return true
			}
		}
	}
	return false
}
