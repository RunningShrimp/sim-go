package server

import (
	"reflect"
	"slices"
	"strings"

	"github.com/RunningShrimp/sim-go/utils"
)

type handlerFunc struct {
	In    []*reflect.Type // 入参列表,按照
	Out   []*reflect.Type // 出参列表
	HFunc reflect.Value   // 处理方法
}

type node struct {
	part     string  // 存储单词片段
	isEnd    bool    // 是否是一个完整的单词u
	children []*node // 子节点
	param    string
	routes   map[string]*handlerFunc
}

// insert
//
//	@Description: 插入节点
//	@receiver t
//	@param word： 待插入单词
func (t *node) insert(word, httpMethod string, routeHandler *handlerFunc) {
	if word == "/" {
		if !t.isEnd {
			t.isEnd = true
		}
		t.addHandler(httpMethod, routeHandler)
		return

	}
	if strings.HasPrefix(word, "/") { // 防止拆出共同前缀"/",无故多出一个节点
		word = word[utils.FindCommonPrefix("/", word):]
	}
	existNode := isExistNode(word, t)
	if existNode == nil {
		createdNode := insertRecursively(word, t)
		createdNode.addHandler(httpMethod, routeHandler)
		return
	}
	existNode.addHandler(httpMethod, routeHandler)

}

func (t *node) search(word string) (*node, map[string]string) {
	if word == "/" {
		return t, nil
	}
	if strings.HasPrefix(word, "/") {
		word = word[utils.FindCommonPrefix("/", word):]
	}
	paramValueMap := make(map[string]string)
	return searchRecursively(word, t, paramValueMap)
}
func (t *node) addHandler(httpMethod string, routeHandler *handlerFunc) {
	if t.routes == nil {
		t.routes = map[string]*handlerFunc{
			httpMethod: routeHandler,
		}
		return
	}
	_, ok := t.routes[httpMethod]
	if ok {
		panic("路由的处理方法已添加，请确认")
	}
	t.routes[httpMethod] = routeHandler
}

func isExistNode(word string, cur *node) *node {
	if word == "" {
		return cur
	}
	for _, childNode := range cur.children {
		commonPrefixLen := utils.FindCommonPrefix(childNode.part, word)
		if commonPrefixLen == 0 {
			continue
		}
		nextWord := word[commonPrefixLen:]
		return isExistNode(nextWord, childNode)
	}
	return nil
}
func searchRecursively(word string, cur *node, paramValueMap map[string]string) (*node, map[string]string) {
	if len(word) <= 0 && cur.isEnd {
		return cur, paramValueMap
	}

	// 记录带有参数的子节点
	existParamNodeIndex := -1
	for index, childNode := range cur.children {
		if len(childNode.param) > 0 {
			existParamNodeIndex = index
		}
		commonPrefixLen := utils.FindCommonPrefix(childNode.part, word)
		if commonPrefixLen == 0 || word[:commonPrefixLen] == "/" {
			continue
		}
		nextWord := word[commonPrefixLen:]
		return searchRecursively(nextWord, childNode, paramValueMap)
	}
	if existParamNodeIndex > -1 {
		tempWord := word[1:]
		paramValue := word[strings.Index(word, "/")+1:]
		if strings.Contains(tempWord, "/") {
			paramValue = tempWord[:strings.Index(tempWord, "/")]
		}
		nextWord := tempWord[utils.FindCommonPrefix(tempWord, paramValue):]
		paramValueMap[cur.children[existParamNodeIndex].param] = paramValue
		return searchRecursively(nextWord, cur.children[existParamNodeIndex], paramValueMap)
	}
	return nil, paramValueMap
}

// insertRecursively
//
//	@Description: 将单词插入到n的后代节点中
//	@receiver t
//	@param word：带插入单词
//	@param n： 节点
func insertRecursively(word string, cur *node) *node {
	if len(word) == 0 {
		cur.isEnd = true
		if cur.children == nil {
			cur.children = make([]*node, 0)
		}
		if cur.routes == nil {
			cur.routes = make(map[string]*handlerFunc)
		}
		return cur
	}
	if strings.HasPrefix(word, "/:") || strings.HasPrefix(word, "/{") {

		tempWord := word[2:]
		param := tempWord
		part := word
		nextWord := ""
		if strings.Contains(param, "/") {
			param = tempWord[:strings.Index(tempWord, "/")]
			part = word[:strings.Index(tempWord, "/")+2]
			nextWord = tempWord[strings.Index(tempWord, "/"):]
		}
		if strings.Contains(param, "}") {
			param = param[:strings.Index(param, "}")]
		}

		isExistParamNode := slices.ContainsFunc(cur.children, func(n *node) bool { // 检测子节点是否已经存在路径参数节点
			return param != n.param && len(n.param) > 0
		})
		if isExistParamNode {
			panic("多个参数节点，无法再次添加路径，请确认路径配置")
		}

		// 是否存在节点node
		existNodeIndex := slices.IndexFunc(cur.children, func(n *node) bool {
			return n.part == part
		})
		if existNodeIndex > -1 {
			return insertRecursively(nextWord, cur.children[existNodeIndex])

		}

		childNode := &node{
			part:     part,
			param:    param,
			children: make([]*node, 0),
			isEnd:    false,
			routes:   make(map[string]*handlerFunc),
		}

		cur.children = append(cur.children, childNode)
		return insertRecursively(nextWord, childNode)

	}

	// 1. 循环遍历n的子节点child，找到child.part与单词一样的前缀
	// 2. 如果找到相同前缀
	// 	2.1 从child.part截取前缀，分割子节点的单词为前缀部分a和后缀部分b
	//	2.2 重新初始前缀a的节点和后缀b的节点，将后缀节点b变为前缀a节点的子节点
	//  2.3 删除原节点，将新节点追加到n的子节点中
	//  2.4 继续截取word不相同部分的后缀c，并将后缀c追加到a节点的子节点中（递归可实现）
	// 3. 如果未找到具有相同前缀的子节点，则初始化新的节点到n的节点中
	for index, childNode := range cur.children {
		commonPrefixLen := utils.FindCommonPrefix(childNode.part, word)
		if commonPrefixLen == 0 || word[:commonPrefixLen] == "/" { // 没有公共前缀
			continue
		}

		// 找到相同前缀，需要childNode.part需要拆分
		prefix := childNode.part[:commonPrefixLen]
		suffix := childNode.part[commonPrefixLen:]

		tempNode := childNode
		if len(suffix) > 0 {
			tempNode = &node{
				part:     prefix,
				isEnd:    false, // 如果没有后缀，说明节点就是一个完成单词
				children: childNode.children,
			}
			subNode := &node{
				part:     suffix,
				isEnd:    true,
				children: make([]*node, 0),
				routes:   childNode.routes,
			}

			// 节点加入切分子节点和新节点
			tempNode.children = append(tempNode.children, subNode)
			// 删除原节点
			cur.children = slices.Delete(cur.children, index, index+1)
			cur.children = append(cur.children, tempNode)
		}
		return insertRecursively(word[commonPrefixLen:], tempNode)

	}

	commonPrefixLen := utils.FindCommonPrefix(cur.part, word)
	newNode := &node{
		part:     word[commonPrefixLen:],
		isEnd:    true,
		children: make([]*node, 0),
		routes:   make(map[string]*handlerFunc),
	}

	if strings.HasPrefix(word, "/") && strings.HasPrefix(cur.part, "/") { // 防止只具有公共前缀"/"时，"/"被删除
		newNode.part = "/" + newNode.part
	}
	cur.children = append(cur.children, newNode)
	return newNode
}
