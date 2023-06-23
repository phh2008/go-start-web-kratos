package result

import (
	"errors"
	"github.com/gin-gonic/gin"
	"helloword/pkg/exception"
	"net/http"
	"time"
)

const success = "200"
const fail = "500"

type Result[T any] struct {
	Code      string `json:"code"`      // 响应码，200表示成功，其它为失败
	Message   string `json:"message"`   // 错误信息
	TimeStamp int64  `json:"timeStamp"` // 时间戳
	Data      T      `json:"data"`      // 响应数据
}

func Success[T any]() *Result[T] {
	return &Result[T]{Code: success, Message: "success", TimeStamp: time.Now().UnixMilli()}
}

func Ok[T any](data T) *Result[T] {
	return &Result[T]{Code: success, Message: "success", TimeStamp: time.Now().UnixMilli(), Data: data}
}

func Fail[T any]() *Result[T] {
	//var zero T
	return Error[T](exception.SysError)
}

func Failure[T any](msg string) *Result[T] {
	return &Result[T]{Code: fail, Message: msg, TimeStamp: time.Now().UnixMilli()}
}

func Error[T any](err error) *Result[T] {
	// 解析具体的错误，获取相应错误码
	var ex exception.BizError
	if ok := errors.As(err, &ex); ok {
		var t T
		return New(ex.Code, ex.Message, t)
	}
	return Failure[T](err.Error())
}

func New[T any](code string, msg string, data T) *Result[T] {
	return &Result[T]{Code: code, Message: msg, Data: data, TimeStamp: time.Now().UnixMilli()}
}

func (a *Result[T]) IsSuccess() bool {
	return a.Code == success
}

func (a *Result[T]) SetCode(code string) *Result[T] {
	a.Code = code
	return a
}

func (a *Result[T]) SetMsg(msg string) *Result[T] {
	a.Message = msg
	return a
}

func (a *Result[T]) SetData(data T) *Result[T] {
	a.Data = data
	return a
}

func (a *Result[T]) Response(c *gin.Context) {
	c.JSON(http.StatusOK, a)
}
