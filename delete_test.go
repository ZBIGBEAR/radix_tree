package radixtree

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDelete(t *testing.T) {
	datas := []string{"tester","slow","water","slower","slowerly","test","watly","team","toast"}
	rt := NewRadixTree()
	var err error
	for i:=range datas{
		err = rt.Insert(datas[i])
		assert.Equal(t, nil, err)
	}
	copyRoot := rt.Root.DeepCopy()
	copyRoot.Childs[0].Childs = []*radix_node{
		copyRoot.Childs[0].Childs[0],
	}
	// 1.删除"toast"
	rt.Delete(datas[8])
	assert.Equal(t, true, rt.Root.compare(copyRoot))
	resultStr, _ := rt.String()
	fmt.Println("result1:", resultStr)
	/*
	{
		"Root": {
			"Val": "#######################",
			"Childs": [
				{
					"Val": "t",
					"Childs": [
						{
							"Val": "e",
							"Childs": [
								{
									"Val": "st",
									"Childs": [
										{
											"Val": "er",
											"Childs": null
										}
									]
								},
								{
									"Val": "am",
									"Childs": null
								}
							]
						}
					]
				},
				{
					"Val": "slow",
					"Childs": [
						{
							"Val": "er",
							"Childs": [
								{
									"Val": "ly",
									"Childs": null
								}
							]
						}
					]
				},
				{
					"Val": "wat",
					"Childs": [
						{
							"Val": "er",
							"Childs": null
						},
						{
							"Val": "ly",
							"Childs": null
						}
					]
				}
			]
		}
	}
	*/
	// 2.删除"team"
	rt.Delete(datas[7])
	copyRoot.Childs[0].Childs[0].Childs = []*radix_node{
		copyRoot.Childs[0].Childs[0].Childs[0],
	}
	assert.Equal(t, true, rt.Root.compare(copyRoot))
	resultStr, _ = rt.String()
	fmt.Println("result2:", resultStr)
	/*
	{
		"Root": {
			"Val": "#######################",
			"Childs": [
				{
					"Val": "t",
					"Childs": [
						{
							"Val": "e",
							"Childs": [
								{
									"Val": "st",
									"Childs": [
										{
											"Val": "er",
											"Childs": null
										}
									]
								}
							]
						}
					]
				},
				{
					"Val": "slow",
					"Childs": [
						{
							"Val": "er",
							"Childs": [
								{
									"Val": "ly",
									"Childs": null
								}
							]
						}
					]
				},
				{
					"Val": "wat",
					"Childs": [
						{
							"Val": "er",
							"Childs": null
						},
						{
							"Val": "ly",
							"Childs": null
						}
					]
				}
			]
		}
	}
	*/
	// 3.删除"watly"
	rt.Delete(datas[6])
	copyRoot.Childs[2].Childs = []*radix_node{
		copyRoot.Childs[2].Childs[0],
	}
	assert.Equal(t, true, rt.Root.compare(copyRoot))
	resultStr, _ = rt.String()
	fmt.Println("result3:", resultStr)
	/*
	{
		"Root": {
			"Val": "#######################",
			"Childs": [
				{
					"Val": "t",
					"Childs": [
						{
							"Val": "e",
							"Childs": [
								{
									"Val": "st",
									"Childs": [
										{
											"Val": "er",
											"Childs": null
										}
									]
								}
							]
						}
					]
				},
				{
					"Val": "slow",
					"Childs": [
						{
							"Val": "er",
							"Childs": [
								{
									"Val": "ly",
									"Childs": null
								}
							]
						}
					]
				},
				{
					"Val": "wat",
					"Childs": [
						{
							"Val": "er",
							"Childs": null
						}
					]
				}
			]
		}
	}
	*/
	// 4.删除"test"
	rt.Delete(datas[5])
	assert.Equal(t, true, rt.Root.compare(copyRoot))
	resultStr, _ = rt.String()
	fmt.Println("result4:", resultStr)
	/*
	{
		"Root": {
			"Val": "#######################",
			"Childs": [
				{
					"Val": "t",
					"Childs": [
						{
							"Val": "e",
							"Childs": [
								{
									"Val": "st",
									"Childs": [
										{
											"Val": "er",
											"Childs": null
										}
									]
								}
							]
						}
					]
				},
				{
					"Val": "slow",
					"Childs": [
						{
							"Val": "er",
							"Childs": [
								{
									"Val": "ly",
									"Childs": null
								}
							]
						}
					]
				},
				{
					"Val": "wat",
					"Childs": [
						{
							"Val": "er",
							"Childs": null
						}
					]
				}
			]
		}
	}
	*/
	// 5.删除"slowerly"
	rt.Delete(datas[4])
	copyRoot.Childs[1].Childs[0].Childs = nil
	assert.Equal(t, true, rt.Root.compare(copyRoot))
	resultStr, _ = rt.String()
	fmt.Println("result5:", resultStr)
	/*
	{
		"Root": {
			"Val": "#######################",
			"Childs": [
				{
					"Val": "t",
					"Childs": [
						{
							"Val": "e",
							"Childs": [
								{
									"Val": "st",
									"Childs": [
										{
											"Val": "er",
											"Childs": null
										}
									]
								}
							]
						}
					]
				},
				{
					"Val": "slow",
					"Childs": [
						{
							"Val": "er",
							"Childs": null
						}
					]
				},
				{
					"Val": "wat",
					"Childs": [
						{
							"Val": "er",
							"Childs": null
						}
					]
				}
			]
		}
	}
	*/
	// 6.删除"slower"
	rt.Delete(datas[3])
	copyRoot.Childs[1].Childs = nil
	assert.Equal(t, true, rt.Root.compare(copyRoot))
	resultStr, _ = rt.String()
	fmt.Println("result6:", resultStr)
	/*
	{
		"Root": {
			"Val": "#######################",
			"Childs": [
				{
					"Val": "t",
					"Childs": [
						{
							"Val": "e",
							"Childs": [
								{
									"Val": "st",
									"Childs": [
										{
											"Val": "er",
											"Childs": null
										}
									]
								}
							]
						}
					]
				},
				{
					"Val": "slow",
					"Childs": null
				},
				{
					"Val": "wat",
					"Childs": [
						{
							"Val": "er",
							"Childs": null
						}
					]
				}
			]
		}
	}
	*/
	rt.Delete("tester")
	copyRoot.Childs[0].Childs[0].Childs[0].Childs = nil
	assert.Equal(t, true, rt.Root.compare(copyRoot))

	rt.Delete("test")
	copyRoot.Childs[0].Childs[0].Childs = nil
	assert.Equal(t, true, rt.Root.compare(copyRoot))

	rt.Delete("te")
	copyRoot.Childs[0].Childs = nil
	assert.Equal(t, true, rt.Root.compare(copyRoot))


	rt.Delete("t")
	copyRoot.Childs = []*radix_node{
		copyRoot.Childs[1],
		copyRoot.Childs[2],
	}
	assert.Equal(t, true, rt.Root.compare(copyRoot))

	rt.Delete("water")
	copyRoot.Childs[1].Childs = nil
	assert.Equal(t, true, rt.Root.compare(copyRoot))

	rt.Delete("wat")
	copyRoot.Childs = []*radix_node{
		copyRoot.Childs[0],
	}
	assert.Equal(t, true, rt.Root.compare(copyRoot))

	rt.Delete("slow")
	copyRoot.Childs = nil
	assert.Equal(t, true, rt.Root.compare(copyRoot))
}
