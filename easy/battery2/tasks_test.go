package battery2

import "testing"

func mk(v int, l, r *Node) *Node { return &Node{Val: v, Left: l, Right: r} }
func leaf(v int) *Node           { return &Node{Val: v} }

// sample is the tree from the README. It is NOT a binary search tree.
//
//	    2
//	   / \
//	  7   5
//	 / \   \
//	2   6   9
//	   / \  /
//	  5  11 4
func sample() *Node {
	return mk(2,
		mk(7, leaf(2), mk(6, leaf(5), leaf(11))),
		mk(5, nil, mk(9, leaf(4), nil)))
}

func negatives() *Node {
	return mk(-10, mk(-3, leaf(-50), nil), leaf(-7))
}

func TestCountLeaves(t *testing.T) {
	cases := []struct {
		name string
		root *Node
		want int
	}{
		{"nil tree", nil, 0},
		{"one node", leaf(1), 1},
		{"sample", sample(), 4},
		{"chain", mk(1, mk(2, mk(3, nil, nil), nil), nil), 1},
		{"full tree of 3", mk(1, leaf(2), leaf(3)), 2},
		{"negatives", negatives(), 2},
	}
	for _, c := range cases {
		if got := CountLeaves(c.root); got != c.want {
			t.Errorf("%s: CountLeaves = %d, want %d", c.name, got, c.want)
		}
	}
}

func TestMaxValue(t *testing.T) {
	cases := []struct {
		name   string
		root   *Node
		want   int
		wantOK bool
	}{
		{"nil tree", nil, 0, false},
		{"one node", leaf(4), 4, true},
		{"sample", sample(), 11, true},
		{"all negative", negatives(), -3, true},
		{"max at the root", mk(100, leaf(1), leaf(2)), 100, true},
		{"max deep on the left", mk(1, mk(2, mk(3, leaf(99), nil), nil), leaf(50)), 99, true},
		{"zero is the max", mk(-1, leaf(0), leaf(-2)), 0, true},
	}
	for _, c := range cases {
		got, ok := MaxValue(c.root)
		if got != c.want || ok != c.wantOK {
			t.Errorf("%s: MaxValue = %d, %v, want %d, %v", c.name, got, ok, c.want, c.wantOK)
		}
	}
}

func TestContains(t *testing.T) {
	cases := []struct {
		name string
		root *Node
		v    int
		want bool
	}{
		{"nil tree", nil, 0, false},
		{"root", sample(), 2, true},
		{"deep left", sample(), 11, true},
		{"deep right", sample(), 4, true},
		{"value appears twice", sample(), 5, true},
		{"missing", sample(), 3, false},
		{"missing zero", sample(), 0, false},
		{"negative found", negatives(), -50, true},
		{"negative missing", negatives(), -4, false},
		{"one node, found", leaf(8), 8, true},
		{"one node, missing", leaf(8), 9, false},
	}
	for _, c := range cases {
		if got := Contains(c.root, c.v); got != c.want {
			t.Errorf("%s: Contains(%d) = %v, want %v", c.name, c.v, got, c.want)
		}
	}
}
