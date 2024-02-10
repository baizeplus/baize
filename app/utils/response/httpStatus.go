package response

type ResCode int64

const (
	Success      ResCode = 200 //成功
	Unauthorized ResCode = 401 //token失效
	Forbidden    ResCode = 403 //没有权限
	Parameter    ResCode = 412 //参数错误
	Error        ResCode = 500 //系统异常
	Waring       ResCode = 600 //详情看msg
)

var codeMsgMap = map[ResCode]string{
	Success:      "success",
	Unauthorized: "无效的令牌",
	Forbidden:    "没有权限，请联系管理员授权",
	Parameter:    "参数错误",
	Error:        "系统异常",
}

func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[Error]
	}
	return msg
}
