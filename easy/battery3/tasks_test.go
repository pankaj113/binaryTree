package battery3

import "testing"

func mk(v int, l, r *Node) *Node { return &Node{Val: v, Left: l, Right: r} }
func leaf(v int) *Node           { return &Node{Val: v} }

func eq(a, b []int) bool {
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

// sample is the first tree from the README:
//
//	    1
//	   / \
//	  2   3
//	 / \
//	4   5
func sample() *Node {
	return mk(1, mk(2, leaf(4), leaf(5)), leaf(3))
}

// bst is the second tree from the README:
//
//	    8
//	   / \
//	  3   10
//	 / \    \
//	1   6    14
//	   / \   /
//	  4   7 13
func bst() *Node {
	return mk(8,
		mk(3, leaf(1), mk(6, leaf(4), leaf(7))),
		mk(10, nil, mk(14, leaf(13), nil)))
}

func rightChain() *Node { return mk(1, nil, mk(2, nil, leaf(3))) }

type travCase struct {
	name string
	root *Node
	want []int
}

func run(t *testing.T, fname string, f func(*Node) []int, cases []travCase) {
	t.Helper()
	for _, c := range cases {
		got := f(c.root)
		if !eq(got, c.want) {
			t.Errorf("%s: %s = %v, want %v", c.name, fname, got, c.want)
		}
	}
}

func TestPreOrder(t *testing.T) {
	run(t, "PreOrder", PreOrder, []travCase{
		{"nil tree", nil, nil},
		{"one node", leaf(9), []int{9}},
		{"sample", sample(), []int{1, 2, 4, 5, 3}},
		{"bst", bst(), []int{8, 3, 1, 6, 4, 7, 10, 14, 13}},
		{"right chain", rightChain(), []int{1, 2, 3}},
	})
}

func TestInOrder(t *testing.T) {
	run(t, "InOrder", InOrder, []travCase{
		{"nil tree", nil, nil},
		{"one node", leaf(9), []int{9}},
		{"sample", sample(), []int{4, 2, 5, 1, 3}},
		{"bst", bst(), []int{1, 3, 4, 6, 7, 8, 10, 13, 14}},
		{"right chain", rightChain(), []int{1, 2, 3}},
	})
}

func TestPostOrder(t *testing.T) {
	run(t, "PostOrder", PostOrder, []travCase{
		{"nil tree", nil, nil},
		{"one node", leaf(9), []int{9}},
		{"sample", sample(), []int{4, 5, 2, 3, 1}},
		{"bst", bst(), []int{1, 4, 7, 6, 3, 13, 14, 10, 8}},
		{"right chain", rightChain(), []int{3, 2, 1}},
	})
}
