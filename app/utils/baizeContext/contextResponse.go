package baizeContext

import (
	"baize/app/constant/sessionStatus"
	"baize/app/utils/response"
	"bytes"
	"compress/gzip"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var (
	success          = &response.ResponseData{Code: response.Success, Msg: response.Success.Msg()}
	parameter        = &response.ResponseData{Code: response.Parameter, Msg: response.Parameter.Msg()}
	unauthorized     = &response.ResponseData{Code: response.Unauthorized, Msg: response.Unauthorized.Msg()}
	forbidden        = &response.ResponseData{Code: response.Forbidden, Msg: response.Forbidden.Msg()}
	successGzip      []byte
	parameterGzip    []byte
	unauthorizedGzip []byte
	forbiddenGzip    []byte
)

func init() {
	marshal, err := json.Marshal(success)
	if err != nil {
		panic(err)
	}
	var successBuf bytes.Buffer
	gz, err := gzip.NewWriterLevel(&successBuf, gzip.BestCompression)
	if err != nil {
		panic(err)
	}
	if _, err = gz.Write(marshal); err != nil {
		panic(err)
	}
	if err = gz.Close(); err != nil {
		panic(err)
	}
	successGzip = successBuf.Bytes()

	marshal, err = json.Marshal(parameter)
	if err != nil {
		panic(err)
	}
	var parameterBuf bytes.Buffer
	gz, err = gzip.NewWriterLevel(&parameterBuf, gzip.BestCompression)
	if err != nil {
		panic(err)
	}
	if _, err = gz.Write(marshal); err != nil {
		panic(err)
	}
	if err = gz.Close(); err != nil {
		panic(err)
	}
	parameterGzip = parameterBuf.Bytes()

	marshal, err = json.Marshal(unauthorized)
	if err != nil {
		panic(err)
	}
	var unauthorizedBuf bytes.Buffer
	gz, err = gzip.NewWriterLevel(&unauthorizedBuf, gzip.BestCompression)
	if err != nil {
		panic(err)
	}
	if _, err = gz.Write(marshal); err != nil {
		panic(err)
	}
	if err = gz.Close(); err != nil {
		panic(err)
	}
	unauthorizedGzip = unauthorizedBuf.Bytes()

	marshal, err = json.Marshal(forbidden)
	if err != nil {
		panic(err)
	}
	var forbiddenBuf bytes.Buffer
	gz, err = gzip.NewWriterLevel(&forbiddenBuf, gzip.BestCompression)
	if err != nil {
		panic(err)
	}
	if _, err = gz.Write(marshal); err != nil {
		panic(err)
	}
	if err = gz.Close(); err != nil {
		panic(err)
	}
	forbiddenGzip = forbiddenBuf.Bytes()
}

func Success(c *gin.Context) {
	c.Writer.Header().Set("Content-Encoding", "gzip")
	c.Data(http.StatusOK, "application/json", successGzip)
	c.Set(sessionStatus.MsgKey, success)
}
func SuccessMsg(c *gin.Context, msg string) {
	rd := &response.ResponseData{Code: response.Success, Msg: msg}
	c.JSON(http.StatusOK, rd)
	c.Set(sessionStatus.MsgKey, rd)
}

func SuccessData(c *gin.Context, data interface{}) {
	rd := &response.ResponseData{Code: response.Success, Msg: response.Success.Msg(), Data: data}
	c.JSON(http.StatusOK, rd)
	c.Set(sessionStatus.MsgKey, rd)
}
func SuccessListData(c *gin.Context, rows interface{}, total int64) {
	c.JSON(http.StatusOK, response.ResponseData{Code: response.Success, Msg: response.Success.Msg(), Data: response.ListData{Rows: rows, Total: total}})
}
func SuccessGzip(c *gin.Context, gzipData []byte) {
	c.Writer.Header().Set("Content-Encoding", "gzip")
	c.Data(http.StatusOK, "application/json", gzipData)
}

func Waring(c *gin.Context, msg string) {
	rd := &response.ResponseData{Code: response.Waring, Msg: msg}
	c.JSON(http.StatusOK, rd)
	c.Set(sessionStatus.MsgKey, msg)
}

func ParameterError(c *gin.Context) {
	c.Writer.Header().Set("Content-Encoding", "gzip")
	c.Data(http.StatusOK, "application/json", parameterGzip)
	c.Set(sessionStatus.MsgKey, parameter)
}
func InvalidToken(c *gin.Context) {
	c.Writer.Header().Set("Content-Encoding", "gzip")
	c.Data(http.StatusOK, "application/json", unauthorizedGzip)
}
func PermissionDenied(c *gin.Context) {
	c.Writer.Header().Set("Content-Encoding", "gzip")
	c.Data(http.StatusOK, "application/json", forbiddenGzip)
	c.Set(sessionStatus.MsgKey, forbidden)
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
