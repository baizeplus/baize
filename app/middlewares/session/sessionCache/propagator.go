package sessionCache

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/errgo.v2/errors"
	"strings"
)

const (
	TokenPrefix   = "Bearer" //令牌前缀
	Authorization = "Authorization"
)

type Propagator struct {
}

func NewPropagator() *Propagator {
	return &Propagator{}
}
func (p *Propagator) Extract(c *gin.Context) (string, error) {

	authHeader := c.Request.Header.Get(Authorization)
	// 按空格分割
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == TokenPrefix) {
		return "", errors.New("token获取失败")
	}
	value := parts[1]
	return value, nil
}
