package problem2

import (
	"fmt"
	"testing"
)

func mk(v int, l, r *Node) *Node { return &Node{Val: v, Left: l, Right: r} }
func leaf(v int) *Node           { return &Node{Val: v} }

func shape(n *Node) string {
	if n == nil {
		return "."
	}
	if n.Left == nil && n.Right == nil {
		return fmt.Sprint(n.Val)
	}
	return fmt.Sprintf("%d(%s,%s)", n.Val, shape(n.Left), shape(n.Right))
}

func eq(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}

// sample is the tree from the README:
//
//	       1
//	     /   \
//	    2     3
//	   / \     \
//	  4   5     6
//	 /     \   / \
//	7       8 9   10
func sample() *Node {
	return mk(1,
		mk(2, mk(4, leaf(7), nil), mk(5, nil, leaf(8))),
		mk(3, nil, mk(6, leaf(9), leaf(10))))
}

func TestZigZag(t *testing.T) {
	cases := []struct {
		name string
		root *Node
		want [][]int
	}{
		{"nil tree", nil, nil},
		{"one node", leaf(1), [][]int{{1}}},
		{"two levels", mk(1, leaf(2), leaf(3)), [][]int{{1}, {3, 2}}},
		{"sample", sample(), [][]int{{1}, {3, 2}, {4, 5, 6}, {10, 9, 8, 7}}},
		{"left chain", mk(1, mk(2, mk(3, leaf(4), nil), nil), nil), [][]int{{1}, {2}, {3}, {4}}},
		{"only right children", mk(1, nil, mk(2, nil, leaf(3))), [][]int{{1}, {2}, {3}}},
		{"full tree of 15", full(1, 4), [][]int{
			{1},
			{3, 2},
			{4, 5, 6, 7},
			{15, 14, 13, 12, 11, 10, 9, 8},
		}},
		{"gaps in a level", mk(1, mk(2, nil, leaf(4)), mk(3, leaf(5), nil)), [][]int{{1}, {3, 2}, {4, 5}}},
		{"repeated values", mk(5, leaf(5), mk(5, leaf(1), leaf(2))), [][]int{{5}, {5, 5}, {1, 2}}},
	}
	for _, c := range cases {
		got := ZigZag(c.root)
		if !eq(got, c.want) {
			t.Errorf("%s: ZigZag(%s) = %v, want %v", c.name, shape(c.root), got, c.want)
		}
	}
}

// full builds a full tree numbered level by level, left to right, starting at i.
func full(i, depth int) *Node {
	if depth == 0 {
		return nil
	}
	return mk(i, full(2*i, depth-1), full(2*i+1, depth-1))
}

func TestZigZagDoesNotModify(t *testing.T) {
	root := sample()
	before := shape(root)
	ZigZag(root)
	if shape(root) != before {
		t.Errorf("ZigZag modified its input: was %s, now %s", before, shape(root))
	}
}
