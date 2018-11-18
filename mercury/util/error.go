package util

const (
	ErrCodeSuccess    = 0
	ErrCodeParameter  = 1001
	ErrCodeUserExist  = 1002
	ErrCodeServerBusy = 1003
)

func GetMessage(code int) (message string) {
	switch code {
	case ErrCodeSuccess:
		message = "success"
	case ErrCodeParameter:
		message = "参数错误"
	case ErrCodeUserExist:
		message = "用户名已经存在"
	case ErrCodeServerBusy:
		message = "服务器繁忙"
	default:
		message = "未知错误"
	}
	return
}
