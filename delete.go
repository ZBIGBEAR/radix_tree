package radixtree

func (rt *RadixTree)Delete(str string) error {
	if rt == nil {
		return UnInitRadixTree
	}

	rt.deleteFromNode(nil, rt.Root, str)
	return nil
}

func (rt *RadixTree) deleteFromNode(parentNode *radix_node, node *radix_node, str string) bool {
	if node.isRootNode(){
		return rt.deleteFromChilds(node, str)
	}
	// 找出str与当前节点的公共前缀
	preFixIndex := calcPreFix(str, node.Val)
	if preFixIndex < 0 {
		// 没有公共前缀，不在当前节点上
		return false
	}
	if preFixIndex == len(str)-1 {
		// 完全匹配
		if len(node.Childs) == 0 {
			// 如果当前节点没有孩子，则删除当前节点
			var newChilds []*radix_node
			for i:=range parentNode.Childs {
				if parentNode.Childs[i].Val != node.Val {
					newChilds = append(newChilds, parentNode.Childs[i])
				}
			}
			parentNode.Childs = newChilds
		}
		return true
	}else{
		// 部分匹配
		return rt.deleteFromChilds(node, str[preFixIndex+1:])
	}
}

func (rt *RadixTree) deleteFromChilds(node *radix_node, str string) bool{
	for i:=range node.Childs {
		if rt.deleteFromNode(node, node.Childs[i], str) {
			return true
		}
	}
	return false
}