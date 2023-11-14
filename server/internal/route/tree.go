package server

import (
	"reflect"
)

type RouterObject struct {
	RouterRoot *node
}
type Handler struct {
	*handlerFunc
	ParamValue map[string]string
}

func (r *RouterObject) AddRoute(method, patten string, handler any) {

	handlerType := reflect.TypeOf(handler)
	handlerValue := reflect.ValueOf(handler)
	r.handlerRouter(method, patten, handlerValue, handlerType)
}

func (r *RouterObject) handlerRouter(method, patten string, handlerValue reflect.Value, handlerType reflect.Type) {
	if handlerType.Kind() != reflect.Func {
		panic("请添加方法")
	}

	argInNum := handlerType.NumIn()
	argOutNum := handlerType.NumOut()
	info := &handlerFunc{
		In:    make([]*reflect.Type, argInNum),
		Out:   make([]*reflect.Type, argOutNum),
		value: handlerValue,
	}
	for i := 0; i < argInNum; i++ {
		in := handlerType.In(i)

		info.In[i] = &in
	}
	for i := 0; i < argOutNum; i++ {
		out := handlerType.Out(i)
		info.Out[i] = &out
	}
	r.RouterRoot.insert(patten, method, info)
}

func (r *RouterObject) GetRoute(httpMethod, patten string) *Handler {
	if r.RouterRoot == nil {
		panic("请先添加路由")
	}
	n, paramVal := r.RouterRoot.search(patten)
	if n == nil {
		return nil
	}
	h, ok := n.routes[httpMethod]
	if !ok && h == nil {
		return nil
	}
	return &Handler{
		handlerFunc: h,
		ParamValue:  paramVal,
	}
}
