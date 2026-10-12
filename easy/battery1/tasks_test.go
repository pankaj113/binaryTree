package battery1

import "testing"

func mk(v int, l, r *Node) *Node { return &Node{Val: v, Left: l, Right: r} }
func leaf(v int) *Node           { return &Node{Val: v} }

// sample is the tree from the README:
//
//	    5
//	   / \
//	  3   8
//	 / \   \
//	1   4   9
func sample() *Node {
	return mk(5, mk(3, leaf(1), leaf(4)), mk(8, nil, leaf(9)))
}

// leftChain is 1 -> 2 -> 3 -> 4 -> 5, every node a left child.
func leftChain() *Node {
	return mk(1, mk(2, mk(3, mk(4, leaf(5), nil), nil), nil), nil)
}

// zigzag goes left, right, left, right.
func zigzag() *Node {
	return mk(10, mk(20, nil, mk(30, mk(40, nil, leaf(50)), nil)), nil)
}

func TestSize(t *testing.T) {
	cases := []struct {
		name string
		root *Node
		want int
	}{
		{"nil tree", nil, 0},
		{"one node", leaf(7), 1},
		{"sample", sample(), 6},
		{"left chain", leftChain(), 5},
		{"zigzag", zigzag(), 5},
		{"only a right child", mk(1, nil, leaf(2)), 2},
	}
	for _, c := range cases {
		if got := Size(c.root); got != c.want {
			t.Errorf("%s: Size = %d, want %d", c.name, got, c.want)
		}
	}
}

func TestSum(t *testing.T) {
	cases := []struct {
		name string
		root *Node
		want int
	}{
		{"nil tree", nil, 0},
		{"one node", leaf(7), 7},
		{"sample", sample(), 30},
		{"left chain", leftChain(), 15},
		{"negative values", mk(-5, leaf(-3), leaf(2)), -6},
		{"zeros", mk(0, leaf(0), leaf(0)), 0},
	}
	for _, c := range cases {
		if got := Sum(c.root); got != c.want {
			t.Errorf("%s: Sum = %d, want %d", c.name, got, c.want)
		}
	}
}

func TestHeight(t *testing.T) {
	cases := []struct {
		name string
		root *Node
		want int
	}{
		{"nil tree", nil, 0},
		{"one node", leaf(7), 1},
		{"sample", sample(), 3},
		{"left chain", leftChain(), 5},
		{"zigzag", zigzag(), 5},
		{"deeper on the right", mk(1, leaf(2), mk(3, nil, mk(4, nil, leaf(5)))), 4},
	}
	for _, c := range cases {
		if got := Height(c.root); got != c.want {
			t.Errorf("%s: Height = %d, want %d", c.name, got, c.want)
		}
	}
}
