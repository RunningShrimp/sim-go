package server

import (
	server "github.com/RunningShrimp/sim-go/server/internal/route"
)

type SimHTTPServer struct {
	Route *server.RouterObject
}

// 全局server实例只会有一个
var serverInstance *SimHTTPServer
