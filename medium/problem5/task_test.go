package problem5

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

// sample is the tree from the README:
//
//	       60
//	     /    \
//	   20      80
//	  /  \       \
//	10    40      90
//	     /  \
//	   30    50
func sample() *Node {
	return mk(60,
		mk(20, leaf(10), mk(40, leaf(30), leaf(50))),
		mk(80, nil, leaf(90)))
}

func TestCommonAncestor(t *testing.T) {
	cases := []struct {
		name   string
		root   *Node
		a, b   int
		want   int
		wantOK bool
	}{
		{"nil tree", nil, 1, 2, 0, false},
		{"two leaves far apart", sample(), 10, 90, 60, true},
		{"two leaves under 40", sample(), 30, 50, 40, true},
		{"leaf and its cousin", sample(), 10, 50, 20, true},
		{"order does not matter", sample(), 50, 10, 20, true},
		{"one is the ancestor of the other", sample(), 20, 30, 20, true},
		{"root and a leaf", sample(), 60, 30, 60, true},
		{"parent and child", sample(), 80, 90, 80, true},
		{"same value twice", sample(), 40, 40, 40, true},
		{"same value twice, root", sample(), 60, 60, 60, true},
		{"a is missing", sample(), 35, 50, 0, false},
		{"b is missing", sample(), 30, 45, 0, false},
		{"both missing", sample(), 1, 2, 0, false},
		{"same missing value twice", sample(), 70, 70, 0, false},
		{"missing, would split at root", sample(), 10, 85, 0, false},
		{"missing, bigger than all", sample(), 90, 100, 0, false},
		{"one node", leaf(5), 5, 5, 5, true},
		{"one node, other missing", leaf(5), 5, 6, 0, false},
		{"negatives", mk(0, mk(-10, leaf(-20), leaf(-5)), leaf(10)), -20, -5, -10, true},
	}
	for _, c := range cases {
		got, ok := CommonAncestor(c.root, c.a, c.b)
		if got != c.want || ok != c.wantOK {
			t.Errorf("%s: CommonAncestor(%s, %d, %d) = %d, %v, want %d, %v",
				c.name, shape(c.root), c.a, c.b, got, ok, c.want, c.wantOK)
		}
	}
}

func TestCommonAncestorDoesNotModify(t *testing.T) {
	root := sample()
	before := shape(root)
	CommonAncestor(root, 30, 90)
	if shape(root) != before {
		t.Errorf("CommonAncestor modified its input: was %s, now %s", before, shape(root))
	}
}
