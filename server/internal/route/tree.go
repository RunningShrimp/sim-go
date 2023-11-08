package server

import (
	"reflect"
)

type Handler struct {
	*handlerFunc
	ParamValue map[string]string
}

var routeTreeRoot = &node{
	part:     "",
	isEnd:    false,
	children: make([]*node, 0),
}

func AddRoute(method, patten string, handler any) {
	if routeTreeRoot == nil {
		routeTreeRoot = &node{
			part:     "",
			isEnd:    false,
			children: make([]*node, 0),
		}
	}

	handlerType := reflect.TypeOf(handler)
	handlerValue := reflect.ValueOf(handler)
	handlerRouter(method, patten, handlerValue, handlerType)
}

func handlerRouter(method, patten string, handlerValue reflect.Value, handlerType reflect.Type) {
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
	routeTreeRoot.insert(patten, method, info)
}

func GetRoute(httpMethod, patten string) *Handler {
	if routeTreeRoot == nil {
		panic("请先添加路由")
	}
	n, paramVal := routeTreeRoot.search(patten)
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
