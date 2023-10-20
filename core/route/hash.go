package route

// 当一个url访问超过这个数，就直接去map中取
const reqNum = 1000

// 自适应hash，访问超出这个限制就把相关信息存放在map
var adaptiveHashRequestMap = make(map[string]map[string]*handlerMethod)

// 快速获取节点
func fastGetHandler(path, httpMethod string) *handlerMethod {
	if n, ok := adaptiveHashRequestMap[httpMethod][path]; ok {
		return n
	}
	return nil
}

func putNode(route *route) {
	if _, ok := adaptiveHashRequestMap[route.method][route.path]; !ok && *route.reqCount >= reqNum {
		adaptiveHashRequestMap[route.method][route.path] = route.handler
	}
}
