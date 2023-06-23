package exception

import "fmt"

var NoLogin = NewBizError("401", "未登录")
var Unauthorized = NewBizError("402", "无权限")
var SignVerifyError = NewBizError("403", "验签失败")
var NotFound = NewBizError("404", "资源未找到")
var SysError = NewBizError("500", "系统错误")
var ParamError = NewBizError("501", "参数错误")
var DBError = NewBizError("502", "数据系统错误")

// BizError 业务错误
type BizError struct {
	Code    string
	Message string
}

func NewBizError(code, message string) BizError {
	return BizError{Code: code, Message: message}
}

func (a BizError) Error() string {
	return fmt.Sprintf("code:%s, message:%s", a.Code, a.Message)
}
