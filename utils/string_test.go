package utils

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
)

func TestFindCommonPrefix(t *testing.T) {
	Convey("具有公共前缀", t, func() {
		commonPrefixLen := FindCommonPrefix("useradd", "userdelete")
		assert.Equal(t, 4, commonPrefixLen)
		commonPrefixLen = FindCommonPrefix("/", "/user")
		assert.Equal(t, 1, commonPrefixLen)
	})
	Convey("没有公共前缀", t, func() {
		length := FindCommonPrefix("user", "add")
		assert.Equal(t, 0, length)
		length = FindCommonPrefix("/er", "e/rou/er")
		assert.Equal(t, 0, length)

	})
}
