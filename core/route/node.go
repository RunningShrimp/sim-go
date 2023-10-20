package route

import (
	"reflect"
)

type handlerMethod struct {
	in    *reflect.Type   // 入参列表,按照
	out   []*reflect.Type // 出参列表
	value reflect.Value   // 处理方法
}
type routes struct {
	method   string
	path     string
	reqCount *uint64
	handler  *handlerMethod
}

type node struct {
	component string // 分割后的单词
	wildChild bool   // 标识是不是一个完整的url路径
	children  []*node
	routes    []*routes //切片用来支持统一路由多个HTTP方法
}

// insert 插入节点
//   - @components 路径拆出的单词
//   - @routes 路由信息
func (n *node) insert(components []string, route *routes) {
	// 1. 如果是根路径即为："/"
	if len(components) == 0 {
		if n.findRoute(route.method, route.path) != nil {
			n.routes = append(n.routes, route)
			n.wildChild = true
		}
		return
	}
	// 2. 根节点的子孙节点
	// 存储当前遍历到的节点
	curNode := n
	for index, component := range components {
		isMatch := false // 是否有匹配的单词
		//  2.1 查询当前节点的子节点是否有对应的单词
		for _, childNode := range curNode.children {
			if component == childNode.component { // 若路径完全匹配，则继续往下
				curNode = childNode
				isMatch = true
				break
			}
		}
		if !isMatch { // 没有匹配的单词，则创建节点加入
			temp := &node{
				component: component,
				wildChild: index == len(components)-1, // 判断是不是最终节点
				children:  make([]*node, 0),
			}
			if index == len(components)-1 && curNode.findRoute(route.method, route.path) == nil {
				temp.routes = append(temp.routes, route)
			}

			curNode.children = append(curNode.children, temp)
		}

		if isMatch && index == len(components)-1 && curNode.findRoute(route.method, route.path) == nil { // 有匹配的节点，且已经到达最末位单词，
			curNode.wildChild = true
			curNode.routes = append(curNode.routes, route)
		}

	}
}

// findRoute
func (n *node) findRoute(method, path string) *routes {
	if len(n.routes) <= 0 {
		return nil
	}
	for _, route := range n.routes {
		if route.path == path && route.method == method {
			return route
		}
	}
	return nil
}

// search 搜索
//
//   - @components 路径拆出的单词
func (n *node) search(components []string) *routes {

}
