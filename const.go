package radixtree

import "errors"

const (
	UnInitRadixTreeStr = "un init radix tree"
	RootNodeVal = "#######################"
)
var (
	UnInitRadixTree = errors.New(UnInitRadixTreeStr)
)
