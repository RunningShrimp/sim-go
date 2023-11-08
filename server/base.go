package server

type IController interface {
}

type IRequest interface {
}

// BaseController 基础的Controller，
// 所有的实现自动注册的rest风格的api都应该继承BaseController
type BaseController struct {
}

// BaseRequest 基础的request参数，所有的请求处理handler的参数都应该继承BaseRequest
type BaseRequest struct {
}
