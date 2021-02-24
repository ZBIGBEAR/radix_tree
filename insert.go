package radixtree

func (rt *RadixTree) Insert(str string) error {
	if rt == nil {
		return UnInitRadixTree
	}
	rt.insertToChild(rt.root, str)
	return nil
}

func (rt *RadixTree) insertToNode(node *radix_node, str string) {
	if isRootNode(node) {
		rt.insertToChild(node, str)
	}
	// 当前节点一定与str有公共前缀
	preFixIndex := calcPreFix(str, node.val)
	// 当前节点剩余字符串，需要作为当前节点的孩子节点
	if preFixIndex < len(node.val)-1{
		rt.insertToChild(node, node.val[preFixIndex+1:])
	}
	// str剩余字符串，需要作为当前节点的孩子节点
	if preFixIndex < len(str)-1{
		rt.insertToChild(node,str[preFixIndex+1:])
	}
}

func (rt *RadixTree) insertToChild(node *radix_node, str string){
	if len(node.childs) == 0 {
		// 当前节点为空，没有孩子节点，则直接初始化一个孩子节点并保存str
		node.childs = append(node.childs, &radix_node{
			val:    str,
			childs: nil,
		})
		return
	}
	for i:=range node.childs{
		if isPartPreFix(str, node.childs[i].val){
			// 找到某个孩子节点与str有公共前缀，则
			rt.insertToNode(node.childs[i], str)
			return
		}
	}
	// 在当前节点节点的孩子节点中没有找到前缀，则新初始化一个孩子节点并保存str
	node.childs = append(node.childs, &radix_node{
		val:    str,
		childs: nil,
	})
	return
}