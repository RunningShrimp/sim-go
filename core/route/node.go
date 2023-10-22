package route

import (
	"fmt"
	"reflect"
	"strings"
)

type handlerMethod struct {
	in    *reflect.Type   // 入参列表,按照
	out   []*reflect.Type // 出参列表
	value reflect.Value   // 处理方法
}
type route struct {
	handler *handlerMethod
}

type radixNode struct {
	part     string
	isEnd    bool
	children map[string]*radixNode
	routes   map[string]*route
	param    string // 存储参数名
}
type radixTree struct {
	root *radixNode
}

func NewRadixTree() *radixTree {
	return &radixTree{
		root: &radixNode{
			part:     "",
			children: make(map[string]*radixNode),
			routes:   make(map[string]*route),
		},
	}
}

func (t *radixTree) insert(path, method string, r *route) {
	if path == "/" {
		t.root.isEnd = true
		t.root.routes[method] = r
		return
	}
	t.insertRecursively(t.root, path, method, r)
}

func (t *radixTree) insertRecursively(node *radixNode, path, method string, r *route) {
	if len(path) == 0 {
		node.isEnd = true
		node.routes[method] = r
		return
	}

	// Check for parameter placeholder
	if strings.HasPrefix(path, "/:") || strings.HasPrefix(path, "/{") {
		paramName := ""
		// Find the parameter name
		if strings.HasPrefix(path, "/:") {
			paramName = path[2:]
		} else {
			paramName = path[2:strings.Index(path, "}")]
		}

		// Check if the node already has a parameter
		if node.param == "" {
			param := paramName[:strings.Index(paramName, "/")]
			subPath := paramName[strings.Index(paramName, "/"):]

			// Split the node and create a new one
			paramNode := &radixNode{
				isEnd:    false,
				children: make(map[string]*radixNode),
				routes:   make(map[string]*route),
				param:    param,
			}
			if strings.HasPrefix(path, "/:") {
				paramNode.part = "/:" + param
			} else if strings.HasPrefix(path, "/{") {
				paramNode.part = fmt.Sprintf("/{%s}", param)

			}
			node.part = node.part[:2]
			node.children[paramNode.part] = paramNode
			if subPath != "" {
				n := &radixNode{
					children: make(map[string]*radixNode),
					isEnd:    false,
					routes:   make(map[string]*route),
					part:     subPath,
				}
				var tempPath string
				if strings.HasPrefix(subPath, "/:") {
					tempPath = subPath[2:]
				} else if strings.HasPrefix(subPath, "/{") {
					tempPath = subPath[2:strings.Index(subPath, "}")]
				}
				n.param = tempPath
				paramNode.children[n.part] = n

				t.insertRecursively(n, subPath, method, r)
			}
		} else {
			node.param = paramName
			node.routes[method] = r
			node.isEnd = true
		}
		return
	}

	for _, child := range node.children {
		commonPrefix := findCommonPrefix(path, child.part)
		if commonPrefix > 0 {
			if commonPrefix == len(child.part) {
				// Common prefix matches the child's key
				t.insertRecursively(child, path[commonPrefix:], method, r)
			} else if commonPrefix == len(path) {
				// Common prefix matches the path to be inserted
				t.insertRecursively(child, child.part[commonPrefix:], method, r)
				child.part = path
				child.isEnd = true
				child.routes[method] = r
			} else {
				// Split the node where the common prefix ends
				newNode := &radixNode{
					part:     child.part[commonPrefix:],
					isEnd:    child.isEnd,
					children: child.children,
					routes:   child.routes,
					param:    child.param,
				}
				child.part = child.part[:commonPrefix]
				child.isEnd = false
				child.children[newNode.part] = newNode
				if commonPrefix < len(path) {
					t.insertRecursively(child, path[commonPrefix:], method, r)
				} else {
					child.isEnd = true
				}
			}
			return
		}
	}
	// No common prefix found, create a new node
	node.children[path] = &radixNode{
		part:     path,
		isEnd:    true,
		children: make(map[string]*radixNode),
		routes: map[string]*route{
			method: r,
		},
	}
}

func (t *radixTree) search(path, method string) (*route, map[string]string) {
	if path == "/" {
		if r, ok := t.root.routes[method]; ok {
			return r, nil
		}
		return nil, nil
	}
	return t.searchRecursively(t.root, path, method, make(map[string]string))
}

func (t *radixTree) searchRecursively(node *radixNode, path, method string, paramValue map[string]string) (*route, map[string]string) {
	if len(path) == 0 && node.isEnd {
		if r, ok := node.routes[method]; ok {
			return r, nil
		}
		return nil, nil
	}

	for _, child := range node.children {
		commonPrefix := findCommonPrefix(path, child.part)
		if commonPrefix > 0 {
			if commonPrefix == len(child.part) {
				return t.searchRecursively(child, path[commonPrefix:], method, paramValue)
			} else if commonPrefix == len(path) {
				return nil, nil
			}
		}
		if child.param != "" {
			// Handle parameter placeholder
			if commonPrefix == 0 {
				continue
			}
			//if child.isEnd {
			//	if r, ok := child.routes[method]; ok {
			//		return r
			//	}
			//}
			//
			//subPath := path[commonPrefix:]
			//// Recursively search for parameter values
			//subNode, ok := child.children[subPath]

		}
	}
	return nil, nil
}

// findCommonPrefix finds the common prefix of two strings
func findCommonPrefix(s1, s2 string) int {
	minLen := len(s1)
	if len(s2) < minLen {
		minLen = len(s2)
	}
	for i := 0; i < minLen; i++ {
		if s1[i] != s2[i] {
			return i
		}
	}
	return minLen
}
