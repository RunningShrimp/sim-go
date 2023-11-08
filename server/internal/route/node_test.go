package server

import (
	"net/http"
	"testing"

	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
)

func Test_node_insert(t *testing.T) {
	convey.Convey("插入", t, func() {
		root := &node{
			part:     "",
			isEnd:    false,
			children: make([]*node, 0),
		}
		root.insert("/user", http.MethodGet, &handlerFunc{})
		root.insert("/route", http.MethodGet, &handlerFunc{})
		root.insert("/us/er", http.MethodGet, &handlerFunc{})
		root.insert("/use/route", http.MethodGet, &handlerFunc{})
		root.insert("/use/rou/er", http.MethodGet, &handlerFunc{})
		root.insert("/use/rou/es", http.MethodGet, &handlerFunc{})
		root.insert("/use/:id", http.MethodGet, &handlerFunc{})
		root.insert("/use/:id/:name", http.MethodGet, &handlerFunc{})
		root.insert("/use/id", http.MethodGet, &handlerFunc{})
		root.insert("/use/:id/age", http.MethodGet, &handlerFunc{})
		n, paramValMap := root.search("/user")
		if assert.NotNil(t, n) {
			assert.Equal(t, "r", n.part)
			assert.NotNil(t, n.routes[http.MethodGet])
			assert.Empty(t, paramValMap)
		}

		n, paramValMap = root.search("/us/er")
		if assert.NotNil(t, n) {
			assert.Equal(t, "/er", n.part)
			assert.Empty(t, paramValMap)
		}
		n, paramValMap = root.search("/use/1")
		if assert.NotNil(t, n) && assert.NotEmpty(t, paramValMap) {
			assert.Equal(t, "id", n.param)
			assert.Equal(t, "/:id", n.part)
			assert.EqualValues(t, "1", paramValMap["id"])
		}
		n, paramValMap = root.search("/use/1/test")
		if assert.NotNil(t, n) && assert.NotEmpty(t, paramValMap) {
			assert.Equal(t, "name", n.param)
			assert.Equal(t, "/:name", n.part)
			assert.EqualValues(t, "1", paramValMap["id"])
			assert.EqualValues(t, "test", paramValMap["name"])
		}
		root.insert("/user/{age}", http.MethodGet, &handlerFunc{})
		n, paramValMap = root.search("/user/12")
		if assert.NotNil(t, n) && assert.NotEmpty(t, paramValMap) {
			assert.Equal(t, "/{age}", n.part)
			assert.Equal(t, "age", n.param)
			assert.EqualValues(t, "12", paramValMap["age"])

		}
		root.insert("/user/{age}/{name}", http.MethodGet, &handlerFunc{})
		n, paramValMap = root.search("/user/12/test")
		if assert.NotNil(t, n) && assert.NotEmpty(t, paramValMap) {
			assert.Equal(t, "/{name}", n.part)
			assert.Equal(t, "name", n.param)
			assert.EqualValues(t, "12", paramValMap["age"])
			assert.EqualValues(t, "test", paramValMap["name"])
		}

		n, paramValMap = root.search("/use/1/age")
		if assert.NotNil(t, n) && assert.NotEmpty(t, paramValMap) {
			assert.Equal(t, "/age", n.part)
			assert.EqualValues(t, "1", paramValMap["id"])
		}
	})
}
