package sim

import (
	"fmt"
	"net/http"
	"time"

	"github.com/RunningShrimp/sim-go/server"
	"go.uber.org/zap"
)

var log = server.Log

type Server struct {
	// 支持自定义 HTTP baseServer handler
	baseServer *http.Server

	serveHandler http.Handler
	// 有默认值
	name string
	port string

	timeOut time.Duration
	maxConn int64
}

func (g Server) Run() {

	log.Info(fmt.Sprintf("[Name-%s-Port-%s] HTTP server is running.", g.name, g.port))
	err := g.baseServer.ListenAndServe()
	if err != nil {
		log.Fatal("Server run failed", zap.String("err", err.Error()))
	}
}
