package e

const (
	Success       = 200
	Error         = 500
	InvalidParams = 400
)

var MsgFlags = map[int]string{
	Success:       "ok",
	Error:         "fail",
	InvalidParams: "请求参数错误",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[Error]
}
