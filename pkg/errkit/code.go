package errkit

import "errors"

type HttpCode int

func (h HttpCode) Int() int {
	return int(h)
}

func (h HttpCode) String() string {
	return HttpCodeMsg[h]
}

func (h HttpCode) ToErr() error {
	return errors.New(HttpCodeMsg[h])
}

const (
	// ok
	HttpOK HttpCode = 200
)

// 基本请求处理code
const (
	// 请求路由不存在
	HttpNoRouter HttpCode = iota + 1000
	// 请求数据不正确
	HttpReqDataFail
	// 请求数据解析失败
	HttpReqUnmarshal
	// 请求处理失败
	HttpReqProcessFail
	// 请求处理完成,存在失败
	HttpReqProcessPartFail
	// 服务器内部错误
	HttpServerErr
	// 服务器内部处理超时
	HttpTimeout
)

// 业务相关code
const (
	// 用户token获取失败
	UserNoToken HttpCode = iota + 2000
	// 用户未登录
	UserUnLogin
	// 用户登录参数无效
	UserSignParamInvalid
	// 用户token处理失败
	UserTokenProcessFail
	// 用户访问权限不足
	UserReqNotPermission
	// 用户账号或密码错误
	UserLoginFail
	// 参数验证失败
	ParamInvalid
)

var HttpCodeMsg = map[HttpCode]string{
	HttpOK: "ok",

	HttpNoRouter:           "请求路由不存在",
	HttpReqDataFail:        "请求数据不正确",
	HttpReqUnmarshal:       "请求数据解析失败",
	HttpReqProcessFail:     "请求处理失败",
	HttpReqProcessPartFail: "请求处理完成,存在失败",
	HttpServerErr:          "服务器内部错误",
	HttpTimeout:            "服务器内部处理超时",

	UserNoToken:          "用户token获取失败,请重新登录",
	UserUnLogin:          "用户登录失效,请重新登录",
	UserSignParamInvalid: "用户登录参数无效",
	UserTokenProcessFail: "用户token处理失败",
	UserReqNotPermission: "用户访问权限不足",
	UserLoginFail:        "用户账号或密码错误",
	ParamInvalid:         "参数验证失败",
}
