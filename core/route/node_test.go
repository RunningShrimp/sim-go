package route

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestInsert(t *testing.T) {
	routeTree := NewRadixTree()

	routeTree.insert("/user", http.MethodGet, nil)
	routeTree.insert("/user", http.MethodPost, &route{})
	routeTree.insert("/us/er", http.MethodGet, &route{})
	r, _ := routeTree.search("/us/er", http.MethodGet)
	assert.NotNil(t, r)
	r, _ = routeTree.search("/user", http.MethodPost)
	assert.NotNil(t, r)
	r, _ = routeTree.search("/user", http.MethodGet)
	assert.Nil(t, r)
	//routeTree.insert("/user/:id", http.MethodGet, &route{})
	//r, _ = routeTree.search("/user/1", http.MethodGet)
	//assert.NotNil(t, r)
	//if assert.NotNil(t, paramValue) {
	//	assert.Equal(t, paramValue["id"], "1")
	//}
	//routeTree.insert("/user/:id", http.MethodPost, nil)
	//r, paramValue = routeTree.search("/user/1", http.MethodPost)
	//assert.NotNil(t, r)
	//if assert.NotNil(t, paramValue) {
	//	assert.Equal(t, paramValue["id"], "1")
	//}
	routeTree.insert("/user/:ids/:name", http.MethodGet, &route{})
	r, _ = routeTree.search("/user/1/test", http.MethodGet)
	assert.NotNil(t, r)
	//assert.NotNil(t, r)
	//if assert.NotNil(t, paramValue) {
	//	assert.Equal(t, paramValue["id"], "1")
	//	assert.Equal(t, paramValue["name"], "test")
	//}
	//r, paramValue = routeTree.search("/user/1/test/test", http.MethodGet)
	//assert.Nil(t, r)

}
