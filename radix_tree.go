package radixtree

// 定义节点
type radix_node struct {
	val string
	childs []*radix_node
}

// 定义RadixTree结构体
type RadixTree struct {
	root *radix_node
}

func NewRadixTree() *RadixTree {
	return &RadixTree{
		root:&radix_node{
			val:RootNodeVal,
		},
	}
}

func isRootNode(node *radix_node) bool {
	return node != nil && node.val == RootNodeVal
}