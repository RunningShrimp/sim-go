package server

import (
	"testing"
)

func TestInnerStdLogger_Info(t *testing.T) {
	logger := InnerStdLogger{}

	logger.Info("测试数据").Set("test", "test").Print()
}
