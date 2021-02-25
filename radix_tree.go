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
func (rn *radix_node)isRootNode() bool {
	return rn != nil && rn.Val == RootNodeVal
}

// 比较两个radix_tree是否相等
func (rn *radix_node) compare(node *radix_node) bool {
	if rn == nil {
		return node == nil
	}
	if node == nil {
		return rn == nil
	}
	if rn.Val != node.Val {
		return false
	}
	if len(rn.Childs) != len(node.Childs) {
		return false
	}
	for i:=range rn.Childs {
		find := false
		for j:=range node.Childs {
			if rn.Childs[j].Val == node.Childs[i].Val {
				find = true
				if !rn.Childs[i].compare(node.Childs[j]){
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

// 深度拷贝
func (rt *radix_node) DeepCopy() *radix_node {
	if rt == nil {
		return nil
	}
	newRoot := &radix_node{
		Val:    rt.Val,
		Childs: nil,
	}
	for i:=range rt.Childs{
		newRoot.Childs = append(newRoot.Childs, rt.Childs[i].DeepCopy())
	}
	return newRoot
}