package server

func AddRoute(method, patten string, handler any) {
	if serverInstance == nil {
		panic("请先初始化服务实例")
	}

	serverInstance.Route.AddRoute(method, patten, handler)
}
