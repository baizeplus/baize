package systemRoutes

import (
	"baize/app/business/system/systemController"
	"baize/app/middlewares"
	"github.com/gin-gonic/gin"
)

func InitSysDictTypeRouter(router *gin.RouterGroup, dictType *systemController.DictType) {
	systemDictType := router.Group("/system/dict/type")
	systemDictType.GET("/list", middlewares.HasPermission("system:dict:list"), dictType.DictTypeList)
	systemDictType.POST("/export", middlewares.HasPermission("system:dict:export"), dictType.DictTypeExport)
	systemDictType.GET("/:dictId", middlewares.HasPermission("system:dict:query"), dictType.DictTypeGetInfo)
	systemDictType.POST("", middlewares.SetLog(dictTypeManagement, middlewares.Insert), middlewares.HasPermission("system:dict:add"), dictType.DictTypeAdd)
	systemDictType.PUT("", middlewares.SetLog(dictTypeManagement, middlewares.Update), middlewares.HasPermission("system:dict:edit"), dictType.DictTypeEdit)
	systemDictType.DELETE("/:dictIds", middlewares.SetLog(dictTypeManagement, middlewares.Delete), middlewares.HasPermission("system:dict:remove"), dictType.DictTypeRemove)
	systemDictType.DELETE("/refreshCache", middlewares.SetLog(dictTypeManagement, middlewares.Clear), middlewares.HasPermission("system:dict:remove"), dictType.DictTypeClearCache)
	systemDictType.GET("/optionSelect", dictType.DictTypeOptionSelect)

}
