package problem1

import (
	"fmt"
	"math"
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

func TestIsValidBST(t *testing.T) {
	const minInt, maxInt = math.MinInt, math.MaxInt
	cases := []struct {
		name string
		root *Node
		want bool
	}{
		{"nil tree", nil, true},
		{"one node", leaf(5), true},
		{"README valid", mk(20, mk(10, leaf(5), leaf(15)), mk(30, leaf(25), leaf(40))), true},
		{"child on the wrong side", mk(10, leaf(15), nil), false},
		{"right child smaller", mk(10, nil, leaf(5)), false},
		{"grandchild breaks the root", mk(20, mk(10, leaf(5), leaf(25)), leaf(30)), false},
		{"grandchild breaks the root (right)", mk(20, leaf(10), mk(30, leaf(15), leaf(40))), false},
		{"deep violation", mk(50, mk(30, mk(20, nil, leaf(25)), mk(40, leaf(35), leaf(55))), leaf(60)), false},
		{"duplicate on the left", mk(10, leaf(10), nil), false},
		{"duplicate on the right", mk(10, nil, leaf(10)), false},
		{"duplicate deep", mk(10, mk(5, nil, leaf(10)), nil), false},
		{"chain to the right", mk(1, nil, mk(2, nil, mk(3, nil, leaf(4)))), true},
		{"chain to the left", mk(4, mk(3, mk(2, leaf(1), nil), nil), nil), true},
		{"zigzag valid", mk(50, mk(10, nil, mk(40, mk(20, nil, leaf(30)), nil)), nil), true},
		{"zigzag valid, between 20 and 40", mk(50, mk(10, nil, mk(40, mk(20, nil, leaf(35)), nil)), nil), true},
		{"zigzag invalid (crosses 40)", mk(50, mk(10, nil, mk(40, mk(20, nil, leaf(41)), nil)), nil), false},
		{"negatives", mk(-10, leaf(-20), leaf(0)), true},
		{"smallest int", mk(0, leaf(minInt), nil), true},
		{"largest int", mk(0, nil, leaf(maxInt)), true},
		{"smallest int at the root", mk(minInt, nil, leaf(maxInt)), true},
		{"largest int at the root", mk(maxInt, leaf(minInt), nil), true},
		{"two smallest ints", mk(minInt, leaf(minInt), nil), false},
		{"two largest ints", mk(maxInt, nil, leaf(maxInt)), false},
	}
	for _, c := range cases {
		if got := IsValidBST(c.root); got != c.want {
			t.Errorf("%s: IsValidBST(%s) = %v, want %v", c.name, shape(c.root), got, c.want)
		}
	}
}

func TestIsValidBSTDoesNotModify(t *testing.T) {
	root := mk(20, mk(10, leaf(5), leaf(25)), leaf(30))
	before := shape(root)
	IsValidBST(root)
	if shape(root) != before {
		t.Errorf("IsValidBST modified its input: was %s, now %s", before, shape(root))
	}
}
