package server

import (
	"go.uber.org/zap"
)

type InnerLogger struct {
	zap.Logger
	loggFilePath string
}
