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

func TestMaxPathSum(t *testing.T) {
	cases := []struct {
		name   string
		root   *Node
		want   int
		wantOK bool
	}{
		{"nil tree", nil, 0, false},
		{"one node", leaf(5), 5, true},
		{"one negative node", leaf(-5), -5, true},
		{"through the root", mk(1, leaf(2), leaf(3)), 6, true},
		{"README", mk(-10, leaf(9), mk(20, leaf(15), leaf(7))), 42, true},
		{"README 2", mk(5, mk(4, mk(11, leaf(7), leaf(2)), nil), mk(8, leaf(13), mk(4, nil, leaf(1)))), 48, true},
		{"all negative", mk(-3, leaf(-2), leaf(-1)), -1, true},
		{"all negative, best is the root", mk(-1, leaf(-2), leaf(-3)), -1, true},
		{"skip a negative child", mk(10, leaf(-5), leaf(3)), 13, true},
		{"path does not need a leaf", mk(10, mk(2, leaf(-20), nil), nil), 12, true},
		{"bend below the root", mk(-50, mk(10, leaf(4), leaf(6)), leaf(1)), 20, true},
		{"cannot use both sides twice", mk(1, mk(2, leaf(3), leaf(4)), leaf(5)), 12, true},
		{"chain", mk(1, mk(-2, mk(3, mk(4, nil, nil), nil), nil), nil), 7, true},
		{"zeros", mk(0, leaf(0), leaf(0)), 0, true},
		{"negative root, two good sides", mk(-1, leaf(10), leaf(10)), 19, true},
	}
	for _, c := range cases {
		before := shape(c.root)
		got, ok := MaxPathSum(c.root)
		if got != c.want || ok != c.wantOK {
			t.Errorf("%s: MaxPathSum(%s) = %d, %v, want %d, %v", c.name, before, got, ok, c.want, c.wantOK)
		}
		if shape(c.root) != before {
			t.Errorf("%s: MaxPathSum modified its input: was %s, now %s", c.name, before, shape(c.root))
		}
	}
}
