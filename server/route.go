package server

import (
	"net/http"
	"reflect"

	server2 "github.com/RunningShrimp/sim-go/server/internal/route"
)

type RouteObject struct{}

func NewRouteObject() *RouteObject {
	return &RouteObject{}
}

// RestGroup 直接注册符合restful风格的api
func (r *RouteObject) RestGroup(patten string, controller IController) {
	methodValue := reflect.ValueOf(controller)
	getMethod := methodValue.MethodByName("Get")
	postMethod := methodValue.MethodByName("Post")
	putMethod := methodValue.MethodByName("Put")
	deleteMethod := methodValue.MethodByName("Delete")
	if getMethod.IsValid() {
		r.Get(patten, getMethod)
	}
	if postMethod.IsValid() {
		r.Post(patten, postMethod)
	}
	if putMethod.IsValid() {
		r.Put(patten, putMethod)
	}
	if deleteMethod.IsValid() {
		r.Delete(patten, deleteMethod)
	}
}

// Get http-get
func (r *RouteObject) Get(patten string, handler any) {
	server2.AddRoute(http.MethodGet, patten, handler)
}

// Post http-post
func (r *RouteObject) Post(patten string, handler any) {
	server2.AddRoute(http.MethodPost, patten, handler)
}

// Put http-put
func (r *RouteObject) Put(patten string, handler any) {
	server2.AddRoute(http.MethodPut, patten, handler)
}

// Delete http-delete
func (r *RouteObject) Delete(patten string, handler any) {
	server2.AddRoute(http.MethodDelete, patten, handler)
}

func (r *RouteObject) DispatchHandler(urlPattern, httpMethod string) *server2.Handler {
	return server2.GetRoute(urlPattern, httpMethod)
}
