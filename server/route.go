package server

import (
	"net/http"
	"reflect"

	server "github.com/RunningShrimp/sim-go/server/internal"
)

// RestGroup 直接注册符合restful风格的api
func RestGroup(patten string, controller IController) {
	methodValue := reflect.ValueOf(controller)
	getMethod := methodValue.MethodByName("Get")
	postMethod := methodValue.MethodByName("Post")
	putMethod := methodValue.MethodByName("Put")
	deleteMethod := methodValue.MethodByName("Delete")
	if getMethod.IsValid() {
		Get(patten, getMethod)
	}
	if postMethod.IsValid() {
		Post(patten, postMethod)
	}
	if putMethod.IsValid() {
		Put(patten, putMethod)
	}
	if deleteMethod.IsValid() {
		Delete(patten, deleteMethod)
	}
}

// Get http-get
func Get(patten string, handler any) {
	server.AddRoute(http.MethodGet, patten, handler)
}

// Post http-post
func Post(patten string, handler any) {
	server.AddRoute(http.MethodPost, patten, handler)
}

// Put http-put
func Put(patten string, handler any) {
	server.AddRoute(http.MethodPut, patten, handler)
}

// Delete http-delete
func Delete(patten string, handler any) {
	server.AddRoute(http.MethodDelete, patten, handler)
}
