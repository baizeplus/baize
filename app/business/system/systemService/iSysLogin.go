package systemService

import (
	"baize/app/business/monitor/monitorModels"
	"baize/app/business/system/systemModels"
	"github.com/gin-gonic/gin"
)

type ILoginService interface {
	Login(c *gin.Context, user *systemModels.User, lb *monitorModels.Logininfor) string
	Register(c *gin.Context, user *systemModels.LoginBody)
	RecordLoginInfo(c *gin.Context, loginUser *monitorModels.Logininfor)
	GenerateCode(c *gin.Context) (m *systemModels.CaptchaVo)
	VerityCaptcha(c *gin.Context, id, base64 string) bool
	ForceLogout(c *gin.Context, token string)
	GetInfo(c *gin.Context) *systemModels.GetInfo
}
