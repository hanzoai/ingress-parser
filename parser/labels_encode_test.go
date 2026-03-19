package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodeNode(t *testing.T) {
	testCases := []struct {
		desc     string
		node     *Node
		expected map[string]string
	}{
		{
			desc: "1 label",
			node: &Node{
				Name: "ingress",
				Children: []*Node{
					{Name: "aaa", Value: "bar"},
				},
			},
			expected: map[string]string{
				"ingress.aaa": "bar",
			},
		},
		{
			desc: "2 labels",
			node: &Node{
				Name: "ingress",
				Children: []*Node{
					{Name: "aaa", Value: "bar"},
					{Name: "bbb", Value: "bur"},
				},
			},
			expected: map[string]string{
				"ingress.aaa": "bar",
				"ingress.bbb": "bur",
			},
		},
		{
			desc: "2 labels, 1 disabled",
			node: &Node{
				Name: "ingress",
				Children: []*Node{
					{Name: "aaa", Value: "bar"},
					{Name: "bbb", Value: "bur", Disabled: true},
				},
			},
			expected: map[string]string{
				"ingress.aaa": "bar",
			},
		},
		{
			desc: "2 levels",
			node: &Node{
				Name: "ingress",
				Children: []*Node{
					{Name: "foo", Children: []*Node{
						{Name: "aaa", Value: "bar"},
					}},
				},
			},
			expected: map[string]string{
				"ingress.foo.aaa": "bar",
			},
		},
		{
			desc: "3 levels",
			node: &Node{
				Name: "ingress",
				Children: []*Node{
					{Name: "foo", Children: []*Node{
						{Name: "bar", Children: []*Node{
							{Name: "aaa", Value: "bar"},
						}},
					}},
				},
			},
			expected: map[string]string{
				"ingress.foo.bar.aaa": "bar",
			},
		},
		{
			desc: "2 levels, same root",
			node: &Node{
				Name: "ingress",
				Children: []*Node{
					{Name: "foo", Children: []*Node{
						{Name: "bar", Children: []*Node{
							{Name: "aaa", Value: "bar"},
							{Name: "bbb", Value: "bur"},
						}},
					}},
				},
			},
			expected: map[string]string{
				"ingress.foo.bar.aaa": "bar",
				"ingress.foo.bar.bbb": "bur",
			},
		},
		{
			desc: "several levels, different root",
			node: &Node{
				Name: "ingress",
				Children: []*Node{
					{Name: "bar", Children: []*Node{
						{Name: "ccc", Value: "bir"},
					}},
					{Name: "foo", Children: []*Node{
						{Name: "bar", Children: []*Node{
							{Name: "aaa", Value: "bar"},
						}},
					}},
				},
			},
			expected: map[string]string{
				"ingress.foo.bar.aaa": "bar",
				"ingress.bar.ccc":     "bir",
			},
		},
		{
			desc: "multiple labels, multiple levels",
			node: &Node{
				Name: "ingress",
				Children: []*Node{
					{Name: "bar", Children: []*Node{
						{Name: "ccc", Value: "bir"},
					}},
					{Name: "foo", Children: []*Node{
						{Name: "bar", Children: []*Node{
							{Name: "aaa", Value: "bar"},
							{Name: "bbb", Value: "bur"},
						}},
					}},
				},
			},
			expected: map[string]string{
				"ingress.foo.bar.aaa": "bar",
				"ingress.foo.bar.bbb": "bur",
				"ingress.bar.ccc":     "bir",
			},
		},
		{
			desc: "slice of struct syntax",
			node: &Node{
				Name: "ingress",
				Children: []*Node{
					{Name: "foo", Children: []*Node{
						{Name: "[0]", Children: []*Node{
							{Name: "aaa", Value: "bar0"},
							{Name: "bbb", Value: "bur0"},
						}},
						{Name: "[1]", Children: []*Node{
							{Name: "aaa", Value: "bar1"},
							{Name: "bbb", Value: "bur1"},
						}},
					}},
				},
			},
			expected: map[string]string{
				"ingress.foo[0].aaa": "bar0",
				"ingress.foo[0].bbb": "bur0",
				"ingress.foo[1].aaa": "bar1",
				"ingress.foo[1].bbb": "bur1",
			},
		},
		{
			desc: "raw value, level 1",
			node: &Node{
				Name: "ingress",
				Children: []*Node{
					{Name: "aaa", RawValue: map[string]interface{}{
						"bbb": "test1",
						"ccc": "test2",
					}},
				},
			},
			expected: map[string]string{
				"ingress.aaa.bbb": "test1",
				"ingress.aaa.ccc": "test2",
			},
		},
		{
			desc: "raw value, level 2",
			node: &Node{
				Name: "ingress",
				Children: []*Node{
					{Name: "aaa", RawValue: map[string]interface{}{
						"bbb": "test1",
						"ccc": map[string]interface{}{
							"ddd": "test2",
						},
					}},
				},
			},
			expected: map[string]string{
				"ingress.aaa.bbb":     "test1",
				"ingress.aaa.ccc.ddd": "test2",
			},
		},
		{
			desc: "raw value, slice of struct",
			node: &Node{
				Name: "ingress",
				Children: []*Node{
					{Name: "aaa", RawValue: map[string]interface{}{
						"bbb": []interface{}{
							map[string]interface{}{
								"ccc": "test1",
								"ddd": "test2",
							},
						},
					}},
				},
			},
			expected: map[string]string{
				"ingress.aaa.bbb[0].ccc": "test1",
				"ingress.aaa.bbb[0].ddd": "test2",
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			labels := EncodeNode(test.node)

			assert.Equal(t, test.expected, labels)
		})
	}
}
