package server

import (
	server2 "github.com/RunningShrimp/sim-go/server/internal/log"
	server "github.com/RunningShrimp/sim-go/server/internal/route"
)

type SimHTTPServer struct {
	Route  *server.RouterObject
	Logger *server2.InnerStdLogger
}

// 全局server实例只会有一个
var serverInstance *SimHTTPServer
