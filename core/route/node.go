package route

import (
	"strings"

	"github.com/RunningShrimp/sim-go/utils"
)

type node struct {
	part     string  // 存储单词片段
	isEnd    bool    // 是否是一个完整的单词u
	children []*node // 子节点
}

// insert
//
//	@Description: 插入节点
//	@receiver t
//	@param word： 待插入单词
func (t *node) insert(word string) {
	if word == "/" {
		if !t.isEnd {
			t.isEnd = true
		}
		return

	}
	if strings.HasPrefix(word, "/") { // 防止拆出共同前缀"/",无故多出一个节点
		word = word[utils.FindCommonPrefix("/", word):]
	}
	t.insertRecursively(word, t)
}

// insertRecursively
//
//	@Description: 将单词插入到n的后代节点中
//	@receiver t
//	@param word：带插入单词
//	@param n： 节点
func (t *node) insertRecursively(word string, n *node) {
	if len(word) == 0 {
		t.isEnd = true
		t.children = make([]*node, 0)
		return
	}
	// 1. 循环遍历n的子节点child，找到child.part与单词一样的前缀
	// 2. 如果找到相同前缀
	// 	2.1 从child.part截取前缀，分割子节点的单词为前缀部分a和后缀部分b
	//	2.2 重新初始前缀a的节点和后缀b的节点，将后缀节点b变为前缀a节点的子节点
	//  2.3 删除原节点，将新节点追加到n的子节点中
	//  2.4 继续截取word不相同部分的后缀c，并将后缀c追加到a节点的子节点中（递归可实现）
	// 3. 如果未找到具有相同前缀的子节点，则初始化新的节点到n的节点中
	for index, childNode := range n.children {
		commonPrefixLen := utils.FindCommonPrefix(childNode.part, word)
		if commonPrefixLen == 0 { // 没有公共前缀
			continue
		}

		// 找到相同前缀，需要childNode.part需要拆分
		prefix := childNode.part[:commonPrefixLen]
		suffix := childNode.part[commonPrefixLen:]

		tempNode := &node{
			part:     prefix,
			isEnd:    suffix == "", // 如果没有后缀，说明节点就是一个完成单词
			children: childNode.children,
		}
		if len(suffix) > 0 {
			subNode := &node{
				part:     suffix,
				isEnd:    true,
				children: make([]*node, 0),
			}

			// 节点加入切分子节点和新节点
			tempNode.children = append(tempNode.children, subNode)
		}

		// 拆分后删除原节点
		n.children = append(n.children[:index], n.children[index+1:]...)
		n.children = append(n.children, tempNode)

		t.insertRecursively(word[commonPrefixLen:], tempNode)
		return

	}

	commonPrefixLen := utils.FindCommonPrefix(n.part, word)
	newNode := &node{
		part:     word[commonPrefixLen:],
		isEnd:    true,
		children: make([]*node, 0),
	}

	if strings.HasPrefix(word, "/") && strings.HasPrefix(n.part, "/") { // 防止只具有公共前缀"/"时，"/"被删除
		newNode.part = "/" + newNode.part
	}
	n.children = append(n.children, newNode)
}

func (t *node) search(word string) *node {
	if word == "/" {
		return t
	}
	return t.searchRecursively(word, t)
}
func (t *node) searchRecursively(word string, n *node) *node {
	if len(word) == 0 {
		return n
	}
	for _, childNode := range n.children {
		commonPrefixLen := utils.FindCommonPrefix(childNode.part, word)
		if commonPrefixLen == 0 {
			continue
		}
	}
	return nil
}
