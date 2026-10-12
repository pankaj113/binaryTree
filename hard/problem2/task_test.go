package problem2

import (
	"fmt"
	"testing"
)

func mk(v int, l, r *Node) *Node { return &Node{Val: v, Left: l, Right: r} }
func leaf(v int) *Node           { return &Node{Val: v} }

// shape writes a tree as text: "v(left,right)", "." for nil, a leaf as just "v".
func shape(n *Node) string {
	if n == nil {
		return "."
	}
	if n.Left == nil && n.Right == nil {
		return fmt.Sprint(n.Val)
	}
	return fmt.Sprintf("%d(%s,%s)", n.Val, shape(n.Left), shape(n.Right))
}

func clone(s []int) []int {
	if s == nil {
		return nil
	}
	return append([]int{}, s...)
}

func same(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestBuildValid(t *testing.T) {
	cases := []struct {
		name    string
		pre, in []int
		want    string
	}{
		{"empty", nil, nil, "."},
		{"empty, not nil", []int{}, []int{}, "."},
		{"one node", []int{7}, []int{7}, "7"},
		{"README", []int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}, "3(9,20(15,7))"},
		{"left chain", []int{1, 2, 3}, []int{3, 2, 1}, "1(2(3,.),.)"},
		{"right chain", []int{1, 2, 3}, []int{1, 2, 3}, "1(.,2(.,3))"},
		{"zigzag", []int{1, 2, 3, 4}, []int{2, 4, 3, 1}, "1(2(.,3(4,.)),.)"},
		{"full tree", []int{4, 2, 1, 3, 6, 5, 7}, []int{1, 2, 3, 4, 5, 6, 7}, "4(2(1,3),6(5,7))"},
		{"negatives", []int{0, -5, 5}, []int{-5, 0, 5}, "0(-5,5)"},
		{"bigger", []int{8, 3, 1, 6, 4, 7, 10, 14, 13}, []int{1, 3, 4, 6, 7, 8, 10, 13, 14}, "8(3(1,6(4,7)),10(.,14(13,.)))"},
	}
	for _, c := range cases {
		pre, in := clone(c.pre), clone(c.in)
		got, ok := Build(pre, in)
		if !ok || shape(got) != c.want {
			t.Errorf("%s: Build(%v, %v) = %s, %v, want %s, true", c.name, c.pre, c.in, shape(got), ok, c.want)
		}
		if !same(pre, c.pre) || !same(in, c.in) {
			t.Errorf("%s: Build modified its input: now %v, %v", c.name, pre, in)
		}
	}
}

func TestBuildInvalid(t *testing.T) {
	cases := []struct {
		name    string
		pre, in []int
	}{
		{"lengths differ", []int{1, 2}, []int{1}},
		{"one empty", []int{1}, nil},
		{"other empty", nil, []int{1}},
		{"different values", []int{1, 2}, []int{1, 3}},
		{"duplicate in both", []int{1, 1}, []int{1, 1}},
		{"duplicate in preorder", []int{1, 2, 1}, []int{2, 1, 3}},
		{"duplicate in inorder", []int{1, 2, 3}, []int{2, 2, 1}},
		{"no such tree", []int{1, 2, 3}, []int{3, 1, 2}},
		{"no such tree, deeper", []int{1, 2, 3, 4}, []int{4, 2, 1, 3}},
		{"no such tree, mixed sides", []int{3, 9, 20, 15, 7}, []int{9, 3, 7, 20, 15}},
		{"README invalid", []int{1, 2, 3}, []int{2, 3, 1, 4}},
	}
	for _, c := range cases {
		pre, in := clone(c.pre), clone(c.in)
		got, ok := Build(pre, in)
		if ok || got != nil {
			t.Errorf("%s: Build(%v, %v) = %s, %v, want nil, false", c.name, c.pre, c.in, shape(got), ok)
		}
		if !same(pre, c.pre) || !same(in, c.in) {
			t.Errorf("%s: Build modified its input: now %v, %v", c.name, pre, in)
		}
	}
}
