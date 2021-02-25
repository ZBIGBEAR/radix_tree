/*
Select方法类似于Delete方法
*/
package radixtree

func (rt *RadixTree) Select(str string) bool {
	if rt == nil {
		return false
	}
	return rt.selectFromNode(rt.Root, str)
}

func (rt *RadixTree) selectFromNode(node *radix_node, str string) bool {
	if node.isRootNode(){
		return rt.selectFromChilds(node, str)
	}
	// 找出str与当前节点的公共前缀
	preFixIndex := calcPreFix(str, node.Val)
	if preFixIndex < 0 {
		// 没有公共前缀，不在当前节点上
		return false
	}
	if preFixIndex == len(str)-1 {
		// 完全匹配
		return true
	}else{
		// 部分匹配
		return rt.selectFromChilds(node, str[preFixIndex+1:])
	}
}

func (rt *RadixTree) selectFromChilds(node *radix_node, str string) bool {
	for i:=range node.Childs {
		if rt.selectFromNode(node.Childs[i], str) {
			return true
		}
	}
	return false
}