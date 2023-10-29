package route

import (
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
		root.insert("/user")
		root.insert("/route")
		root.insert("/us/er")
		root.insert("/use/route")
		root.insert("/use/rou/er")
		root.insert("/use/rou/es")
		n := root.search("/user")
		if assert.NotNil(t, n) {
			assert.Equal(t, "/user", n.part)
		}

	})
}
