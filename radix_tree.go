package radixtree

import (
	"bytes"
	"encoding/json"
)

// 定义节点
type radix_node struct {
	Val    string        `json:"Val"`
	Childs []*radix_node `json:"Childs"`
}

// 定义RadixTree结构体
type RadixTree struct {
	Root *radix_node `json:"Root"`
}

func NewRadixTree() *RadixTree {
	return &RadixTree{
		Root:&radix_node{
			Val: RootNodeVal,
		},
	}
}

// 判断是不是根节点
func isRootNode(node *radix_node) bool {
	return node != nil && node.Val == RootNodeVal
}

// 比较两个radix_tree是否相等
func compare(root1, root2 *radix_node) bool {
	if root1 == nil {
		return root2 == nil
	}
	if root2 == nil {
		return root1 == nil
	}
	if root1.Val != root2.Val {
		return false
	}
	if len(root1.Childs) != len(root2.Childs) {
		return false
	}
	for i:=range root1.Childs {
		find := false
		for j:=range root2.Childs {
			if root2.Childs[j].Val == root1.Childs[i].Val {
				find = true
				if !compare(root1.Childs[i], root2.Childs[j]){
					return false
				}
			}
		}
		if !find {
			return false
		}
	}
	return true
}

// 输出结构化Radix Tree
func (rt *RadixTree) String() (string, error) {
	b, err := json.Marshal(*rt)
	if err != nil {
		return "", err
	}
	var out bytes.Buffer
	err = json.Indent(&out, b, "", "	")
	if err != nil {
		return "", err
	}
	return out.String(), nil
}