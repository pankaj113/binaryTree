package problem4

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

// prices is the tree from the README:
//
//	       40
//	     /    \
//	   20      60
//	  /  \    /  \
//	10   30  50   80
//	  \          /
//	  15        70
func prices() *Node {
	return mk(40,
		mk(20, mk(10, nil, leaf(15)), leaf(30)),
		mk(60, leaf(50), mk(80, leaf(70), nil)))
}

func TestRange(t *testing.T) {
	cases := []struct {
		name   string
		root   *Node
		lo, hi int
		want   []int
	}{
		{"nil tree", nil, 0, 100, nil},
		{"everything", prices(), 0, 100, []int{10, 15, 20, 30, 40, 50, 60, 70, 80}},
		{"middle", prices(), 25, 55, []int{30, 40, 50}},
		{"bounds are included", prices(), 15, 30, []int{15, 20, 30}},
		{"one value", prices(), 40, 40, []int{40}},
		{"lo == hi, missing", prices(), 41, 41, nil},
		{"nothing in range", prices(), 31, 39, nil},
		{"below everything", prices(), -10, 5, nil},
		{"above everything", prices(), 81, 1000, nil},
		{"lo > hi", prices(), 60, 20, nil},
		{"left part", prices(), 0, 20, []int{10, 15, 20}},
		{"right part", prices(), 61, 100, []int{70, 80}},
		{"negatives", mk(0, mk(-10, leaf(-20), leaf(-5)), leaf(10)), -15, 0, []int{-10, -5, 0}},
	}
	for _, c := range cases {
		got := Range(c.root, c.lo, c.hi)
		if !eq(got, c.want) {
			t.Errorf("%s: Range(%s, %d, %d) = %v, want %v", c.name, shape(c.root), c.lo, c.hi, got, c.want)
		}
	}
}

func TestKthSmallest(t *testing.T) {
	cases := []struct {
		name   string
		root   *Node
		k      int
		want   int
		wantOK bool
	}{
		{"nil tree", nil, 1, 0, false},
		{"smallest", prices(), 1, 10, true},
		{"second", prices(), 2, 15, true},
		{"third", prices(), 3, 20, true},
		{"the root", prices(), 5, 40, true},
		{"seventh", prices(), 7, 60, true},
		{"largest", prices(), 9, 80, true},
		{"k too big", prices(), 10, 0, false},
		{"k is zero", prices(), 0, 0, false},
		{"k negative", prices(), -1, 0, false},
		{"one node", leaf(7), 1, 7, true},
		{"one node, k = 2", leaf(7), 2, 0, false},
		{"right chain", mk(1, nil, mk(2, nil, mk(3, nil, leaf(4)))), 3, 3, true},
	}
	for _, c := range cases {
		got, ok := KthSmallest(c.root, c.k)
		if got != c.want || ok != c.wantOK {
			t.Errorf("%s: KthSmallest(%s, %d) = %d, %v, want %d, %v", c.name, shape(c.root), c.k, got, ok, c.want, c.wantOK)
		}
	}
}

func TestDoesNotModify(t *testing.T) {
	root := prices()
	before := shape(root)
	Range(root, 15, 70)
	KthSmallest(root, 4)
	if shape(root) != before {
		t.Errorf("the tree was modified: was %s, now %s", before, shape(root))
	}
}
