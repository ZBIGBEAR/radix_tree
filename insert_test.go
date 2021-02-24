package radixtree

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsert(t *testing.T) {
	rt := NewRadixTree()

	// 1.插入tester
	word1 := "tester"
	err := rt.Insert(word1)
	assert.Equal(t,nil, err)
	// 手动生成radix_tree，用于比对结果
	root := &radix_node{
		Val: RootNodeVal,
		Childs: []*radix_node{
			&radix_node{
				Val:    word1,
				Childs: nil,
			},
		},
	}
	assert.Equal(t, true, compare(root, rt.Root))
	// 打印结果
	resultStr, err := rt.String()
	assert.Equal(t, nil, err)
	fmt.Println("result1:",resultStr)
	/*
	{
		"Root": {
			"Val": "#######################",
			"Childs": [
				{
					"Val": "tester",
					"Childs": null
				}
			]
		}
	}
	*/

	// 2.插入slow
	word2 := "slow"
	err = rt.Insert(word2)
	assert.Equal(t, nil, err)
	root.Childs = append(root.Childs, &radix_node{
		Val:    word2,
		Childs: nil,
	})
	assert.Equal(t, true, compare(root, rt.Root))
	// 打印结果
	resultStr, err = rt.String()
	assert.Equal(t, nil, err)
	fmt.Println("result2:", resultStr)
	/*
	{
		"Root": {
			"Val": "#######################",
			"Childs": [
				{
					"Val": "tester",
					"Childs": null
				},
				{
					"Val": "slow",
					"Childs": null
				}
			]
		}
	}
	*/

	// 3.插入water
	word3 := "water"
	err = rt.Insert(word3)
	assert.Equal(t, nil, err)
	root.Childs = append(root.Childs, &radix_node{
		Val:    word3,
		Childs: nil,
	})
	assert.Equal(t, true, compare(root, rt.Root))
	// 打印结果
	resultStr, err = rt.String()
	assert.Equal(t, nil, err)
	fmt.Println("result3:", resultStr)
	/*
	{
		"Root": {
			"Val": "#######################",
			"Childs": [
				{
					"Val": "tester",
					"Childs": null
				},
				{
					"Val": "slow",
					"Childs": null
				},
				{
					"Val": "water",
					"Childs": null
				}
			]
		}
	}
	*/

	// 4.插入slower
	word4 := "slower"
	err = rt.Insert(word4)
	assert.Equal(t, nil, err)
	root.Childs[1].Childs = append(root.Childs[1].Childs, &radix_node{
		Val:    "er",
		Childs: nil,
	})
	assert.Equal(t, true, compare(root, rt.Root))
	// 打印结果
	resultStr, err = rt.String()
	assert.Equal(t, nil, err)
	fmt.Println("result4:", resultStr)
	/*
	{
		"Root": {
			"Val": "#######################",
			"Childs": [
				{
					"Val": "tester",
					"Childs": null
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
					"Val": "water",
					"Childs": null
				}
			]
		}
	}
	*/

	// 5.插入slowerly
	word5 := "slowerly"
	err = rt.Insert(word5)
	assert.Equal(t, nil, err)
	root.Childs[1].Childs[0].Childs = append(root.Childs[1].Childs[0].Childs, &radix_node{
		Val:    "ly",
		Childs: nil,
	})
	assert.Equal(t, true, compare(root, rt.Root))
	// 打印结果
	resultStr, err = rt.String()
	assert.Equal(t, nil, err)
	fmt.Println("result5:", resultStr)
	/*
	{
		"Root": {
			"Val": "#######################",
			"Childs": [
				{
					"Val": "tester",
					"Childs": null
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
					"Val": "water",
					"Childs": null
				}
			]
		}
	}
	*/

	// 6.插入test
	word6 := "test"
	err = rt.Insert(word6)
	assert.Equal(t, nil, err)
	root.Childs[0].Val = word6
	root.Childs[0].Childs = append(root.Childs[0].Childs, &radix_node{
		Val:    "er",
		Childs: nil,
	})
	assert.Equal(t, true, compare(root, rt.Root))
	// 打印结果
	resultStr, err = rt.String()
	assert.Equal(t, nil, err)
	fmt.Println("result6:", resultStr)
	/*
	{
		"Root": {
			"Val": "#######################",
			"Childs": [
				{
					"Val": "test",
					"Childs": [
						{
							"Val": "er",
							"Childs": null
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
					"Val": "water",
					"Childs": null
				}
			]
		}
	}
	*/

	// 7.插入watly
	word7 := "watly"
	err = rt.Insert(word7)
	assert.Equal(t, nil, err)
	root.Childs[2].Val = "wat"
	root.Childs[2].Childs = append(root.Childs[2].Childs, &radix_node{
		Val:    "er",
		Childs: nil,
	})
	root.Childs[2].Childs = append(root.Childs[2].Childs, &radix_node{
		Val:    "ly",
		Childs: nil,
	})
	assert.Equal(t, true, compare(root, rt.Root))
	// 打印结果
	resultStr, err = rt.String()
	assert.Equal(t, nil, err)
	fmt.Println("result7:", resultStr)
	/*
	{
		"Root": {
			"Val": "#######################",
			"Childs": [
				{
					"Val": "test",
					"Childs": [
						{
							"Val": "er",
							"Childs": null
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

	// 8.插入team
	word8 := "team"
	err = rt.Insert(word8)
	assert.Equal(t, nil, err)
	root.Childs[0].Val = "te"
	root.Childs[0].Childs = []*radix_node{
		&radix_node{
			Val: "st",
			Childs: []*radix_node{
				&radix_node{
					Val:    "er",
					Childs: nil,
					},
				},
		},&radix_node{
			Val:    "am",
			Childs: nil,
		},
	}
	assert.Equal(t, true, compare(root, rt.Root))
	// 打印结果
	resultStr, err = rt.String()
	assert.Equal(t, nil, err)
	fmt.Println("result8:", resultStr)
	/*
	{
		"Root": {
			"Val": "#######################",
			"Childs": [
				{
					"Val": "te",
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

	// 9.插入toast
	word9 := "toast"
	err = rt.Insert(word9)
	assert.Equal(t, nil, err)
	root.Childs[0].Val = "t"
	root.Childs[0].Childs = []*radix_node{
		&radix_node{
			Val: "e",
			Childs: []*radix_node{
				&radix_node{
					Val: "st",
					Childs: []*radix_node{
						&radix_node{
							Val:    "er",
							Childs: nil,
						},
					},
				},&radix_node{
					Val:    "am",
					Childs: nil,
				},
			},
		},&radix_node{
			Val:    "oast",
			Childs: nil,
		},
	}
	assert.Equal(t, true, compare(root, rt.Root))
	// 打印结果
	resultStr, err = rt.String()
	assert.Equal(t, nil, err)
	fmt.Println("result9:", resultStr)
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
						},
						{
							"Val": "oast",
							"Childs": null
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
}