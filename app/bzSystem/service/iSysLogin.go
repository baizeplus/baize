package service

import (
	monitorModels "baize/app/bzMonitor/models"
	"baize/app/bzSystem/models"
	"github.com/gin-gonic/gin"
)

type ILoginService interface {
	Login(c *gin.Context, user *models.User, lb *monitorModels.Logininfor) string
	RecordLoginInfo(c *gin.Context, loginUser *monitorModels.Logininfor)
	GenerateCode(c *gin.Context) (m *models.CaptchaVo)
	VerityCaptcha(c *gin.Context, id, base64 string) bool
	ForceLogout(c *gin.Context, token string)
	GetInfo(c *gin.Context) *models.GetInfo
}
