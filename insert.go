package radixtree

func (rt *RadixTree) Insert(str string) error {
	if rt == nil {
		return UnInitRadixTree
	}
	rt.insertToChild(rt.Root, str)
	return nil
}

func (rt *RadixTree) insertToNode(node *radix_node, str string) {
	if str==""{
		return
	}
	if isRootNode(node) {
		rt.insertToChild(node, str)
	}
	// 当前节点一定与str有公共前缀
	preFixIndex := calcPreFix(str, node.Val)
	// 当前节点剩余字符串，需要作为当前节点的孩子节点
	if preFixIndex < len(node.Val)-1{
		oldVal := node.Val
		node.Val = oldVal[0:preFixIndex+1]
		if len(node.Childs) == 0{
			rt.insertToChild(node, oldVal[preFixIndex+1:])
		}else{
			newNode := &radix_node{
				Val:    oldVal[preFixIndex+1:],
				Childs: node.Childs,
			}
			node.Childs = []*radix_node{
				newNode,
			}
		}
	}
	// str剩余字符串，需要作为当前节点的孩子节点
	if preFixIndex < len(str)-1{
		rt.insertToChild(node,str[preFixIndex+1:])
	}
}

func (rt *RadixTree) insertToChild(node *radix_node, str string){
	if str == "" {
		return
	}
	if len(node.Childs) == 0 {
		// 当前节点为空，没有孩子节点，则直接初始化一个孩子节点并保存str
		node.Childs = append(node.Childs, &radix_node{
			Val:    str,
			Childs: nil,
		})
		return
	}
	for i:=range node.Childs {
		if isPartPreFix(str, node.Childs[i].Val){
			// 找到某个孩子节点与str有公共前缀，则
			rt.insertToNode(node.Childs[i], str)
			return
		}
	}
	// 在当前节点节点的孩子节点中没有找到前缀，则新初始化一个孩子节点并保存str
	node.Childs = append(node.Childs, &radix_node{
		Val:    str,
		Childs: nil,
	})
	return
}