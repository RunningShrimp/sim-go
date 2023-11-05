package route

var routeTreeRoot = &node{
	part:     "",
	isEnd:    false,
	children: make([]*node, 0),
}
