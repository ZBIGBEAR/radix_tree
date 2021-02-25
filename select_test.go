package radixtree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSelect(t *testing.T) {
	datas := []string{"tester","slow","water","slower","slowerly","test","watly","team","toast"}
	rt := NewRadixTree()
	var err error
	for i:=range datas{
		err = rt.Insert(datas[i])
		assert.Equal(t, nil, err)
	}
	var find bool
	for i:=range datas{
		find = rt.Select(datas[i])
		assert.Equal(t, true, find)
	}

	find = rt.Select("t")
	assert.Equal(t, true, find)

	find = rt.Select("tt")
	assert.Equal(t, false, find)

	find = rt.Select("wat")
	assert.Equal(t, true, find)

	find = rt.Select("waterer")
	assert.Equal(t, false, find)

	find = rt.Select("te")
	assert.Equal(t, true, find)

	find = rt.Select("hello")
	assert.Equal(t, false, find)
}
