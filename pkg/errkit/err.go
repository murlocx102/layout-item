package errkit

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// 如果需要判断是否包含某个错误请使用 errors.Is
// 如果判断是否包含的某种类型的错误,请使用 errors.As

// 包含错误码和错误消息的接口,适用于应用程序错误
type Err interface {
	error
	Code() int
	Message() string
}

// 添加data数据接口
type ErrData interface {
	Err
	AddData(any)
	Data() any
}

// HttpErr http错误
type HttpErr struct {
	data    any    //数据
	code    int    //错误码
	message string //错误信息
}

func (h *HttpErr) Code() int {
	return h.code
}

func (h *HttpErr) Message() string {
	return h.message
}

func (h *HttpErr) AddData(data any) {
	h.data = data
}

func (h *HttpErr) Data() any {
	return h.data
}

func (h *HttpErr) Error() string {
	return fmt.Sprintf("Err-code:%d,message:%s", h.code, h.message)
}

// New 给定的错误码,返回一个定义错误信息的Err
func New(code HttpCode, err error) *HttpErr {
	var message = ""

	if msg, ok := HttpCodeMsg[code]; ok {
		message = msg
	}

	if err != nil {
		message = strings.Join([]string{message, err.Error()}, ":")
	}

	return &HttpErr{
		code:    code.Int(),
		message: message,
	}
}

// grpcerr http错误
type grpcerr struct {
	s       *status.Status
	code    int    //错误码
	message string //错误信息
}

func (g *grpcerr) Code() codes.Code {
	if g == nil || g.s == nil {
		return codes.OK
	}

	return codes.Code(g.s.Code())
}

func (g *grpcerr) Message() string {
	if g == nil || g.s == nil {
		return ""
	}

	return g.s.Message()
}

// 错误添加自定义信息.只返回错误信息及自定义信息
func ErrWithMsg(err error, message string) error {
	if err != nil {
		return errors.WithMessage(err, message)
	}

	return errors.New(message)
}
