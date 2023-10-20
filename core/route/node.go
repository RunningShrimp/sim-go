package route

import (
	"reflect"
)

type handlerMethod struct {
	in    *reflect.Type   // 入参列表,按照
	out   []*reflect.Type // 出参列表
	value reflect.Value   // 处理方法
}
type route struct {
	method   string
	path     string
	reqCount *uint64
	handler  *handlerMethod
}

type node struct {
	component string
	wildChild bool
	children  []*node
	route     *route
}

// insert 插入节点
//   - @components 路径拆出的单词
//   - @route 路由信息
func (n *node) insert(components []string, route *route) {
	// 1. 如果是根路径即为："/"
	if len(components) == 0 {
		n.route = route
		return
	}
	// 2. 根节点的子孙节点
	// 存储当前遍历到的节点
	curNode := n
	for _, component := range components {
		//  2.1 查询当前节点的子节点是否有对应的单词
		for _, childNode := range curNode.children {
			if component == childNode.component { // 若路径完全匹配，则继续往下
				curNode = childNode
				break
			}
		}

	}
}

// search 搜索
//
//   - @components 路径拆出的单词
func (n *node) search(components []string) *route {
	if len(components) == 0 {
		return n.route
	}

	c := components[0]

	// Search for an exact match
	for _, child := range n.children {
		if child.component == c {
			r := child.search(components[1:])
			if r != nil {
				return r
			}
			break
		}
	}

	// Search for a wildcard match
	for _, child := range n.children {
		if child.wildChild {
			r := child.search(components[1:])
			if r != nil {
				return r
			}
			break
		}
	}

	return nil
}
