package systemRoutes

import (
	"baize/app/business/system/systemController"
	"baize/app/middlewares"
	"github.com/gin-gonic/gin"
)

func InitSysDictDataRouter(router *gin.RouterGroup, dictData *systemController.DictData) {
	systemDictData := router.Group("/system/dict/data")
	systemDictData.GET("/list", middlewares.HasPermission("system:dict:list"), dictData.DictDataList)
	systemDictData.GET("/export", middlewares.HasPermission("system:dict:export"), dictData.DictDataExport)
	systemDictData.GET("/:dictCode", middlewares.HasPermission("system:dict:query"), dictData.DictDataGetInfo)
	systemDictData.GET("/type/:dictType", dictData.DictDataType)
	systemDictData.POST("", middlewares.HasPermission("system:dict:add"), dictData.DictDataAdd)
	systemDictData.PUT("", middlewares.HasPermission("system:dict:edit"), dictData.DictDataEdit)
	systemDictData.DELETE("/:dictCodes", middlewares.HasPermission("system:dict:remove"), dictData.DictDataRemove)

}
