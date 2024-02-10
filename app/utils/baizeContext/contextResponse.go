package baizeContext

import (
	"baize/app/utils/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Success(c *gin.Context) {
	c.JSON(http.StatusOK, response.ResponseData{Code: response.Success, Msg: response.Success.Msg()})
}
func SuccessMsg(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, response.ResponseData{Code: response.Success, Msg: msg})
}

func SuccessData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, response.ResponseData{Code: response.Success, Msg: response.Success.Msg(), Data: data})
}
func SuccessListData(c *gin.Context, rows interface{}, total *int64) {
	c.JSON(http.StatusOK, response.ResponseData{Code: response.Success, Msg: response.Success.Msg(), Data: response.ListData{Rows: rows, Total: total}})
}

func Waring(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, response.ResponseData{Code: response.Waring, Msg: msg})
}

func ParameterError(c *gin.Context) {
	c.JSON(http.StatusOK, response.ResponseData{Code: response.Parameter, Msg: response.Parameter.Msg()})
}
func InvalidToken(c *gin.Context) {
	c.JSON(http.StatusOK, response.ResponseData{Code: response.Unauthorized, Msg: response.Unauthorized.Msg()})
}
func PermissionDenied(c *gin.Context) {
	c.JSON(http.StatusOK, response.ResponseData{Code: response.Forbidden, Msg: response.Forbidden.Msg()})
}
func DataPackageExcel(c *gin.Context, data []byte) {
	c.Header("Content-Type", "application/vnd.ms-excel")
	c.Header("Pragma", "public")
	c.Header("Cache-Control", "no-store")
	c.Header("Cache-Control", "max-age=0")
	c.Header("Content-Length", strconv.Itoa(len(data)))
	c.Data(http.StatusOK, "application/vnd.ms-excel", data)
}
func DataPackageZip(c *gin.Context, data []byte) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Expose-Headers", "Content-Disposition")
	c.Header("Content-Disposition", "attachment; filename=\"baize.zip\"")
	c.Header("Content-Length", strconv.Itoa(len(data)))
	c.Data(http.StatusOK, "application/octet-stream; charset=UTF-8", data)
}
