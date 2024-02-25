package monitorController

import (
	"baize/app/business/monitor/monitorModels"
	"baize/app/utils/baizeContext"
	"github.com/gin-gonic/gin"
)

type InfoServer struct {
}

func NewInfoServer() *InfoServer {
	return &InfoServer{}
}

func (isc *InfoServer) GetInfoServer(c *gin.Context) {

	baizeContext.SuccessData(c, monitorModels.NewServer())
}
