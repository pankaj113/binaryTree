package problem3

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

func collect(n *Node, seen map[*Node]bool) {
	if n == nil {
		return
	}
	seen[n] = true
	collect(n.Left, seen)
	collect(n.Right, seen)
}

// rightChain builds lo -> lo+1 -> ... -> hi, every node a right child.
func rightChain(lo, hi int) *Node {
	if lo > hi {
		return nil
	}
	return mk(lo, nil, rightChain(lo+1, hi))
}

func TestIsBalanced(t *testing.T) {
	cases := []struct {
		name string
		root *Node
		want bool
	}{
		{"nil tree", nil, true},
		{"one node", leaf(1), true},
		{"one child", mk(2, leaf(1), nil), true},
		{"chain of 3", rightChain(1, 3), false},
		{"full tree", mk(4, mk(2, leaf(1), leaf(3)), mk(6, leaf(5), leaf(7))), true},
		{"README balanced", mk(5, mk(2, leaf(1), mk(3, nil, leaf(4))), mk(7, leaf(6), nil)), true},
		{"README not balanced", mk(5, mk(2, leaf(1), mk(3, nil, leaf(4))), leaf(7)), false},
		// Same height on both sides of the root, but 2 and 6 are off balance.
		{"root ok, children not", mk(4, mk(3, mk(2, leaf(1), nil), nil), mk(5, nil, mk(6, nil, leaf(7)))), false},
		{"deep imbalance", mk(8, mk(4, mk(2, leaf(1), leaf(3)), mk(6, leaf(5), leaf(7))), mk(12, mk(10, leaf(9), nil), nil)), false},
		{"deep balanced", mk(8, mk(4, mk(2, leaf(1), leaf(3)), mk(6, leaf(5), leaf(7))), mk(12, mk(10, leaf(9), nil), leaf(13))), true},
	}
	for _, c := range cases {
		if got := IsBalanced(c.root); got != c.want {
			t.Errorf("%s: IsBalanced(%s) = %v, want %v", c.name, shape(c.root), got, c.want)
		}
	}
}

func TestBalance(t *testing.T) {
	cases := []struct {
		name string
		root *Node
		want string
	}{
		{"nil tree", nil, "."},
		{"one node", leaf(5), "5"},
		{"two nodes", rightChain(1, 2), "1(.,2)"},
		{"three nodes", rightChain(1, 3), "2(1,3)"},
		{"four nodes", rightChain(1, 4), "2(1,3(.,4))"},
		{"README", rightChain(1, 7), "4(2(1,3),6(5,7))"},
		{"six nodes", rightChain(1, 6), "3(1(.,2),5(4,6))"},
		{"ten nodes", rightChain(1, 10), "5(2(1,3(.,4)),8(6(.,7),9(.,10)))"},
		{"left chain", mk(4, mk(3, mk(2, leaf(1), nil), nil), nil), "2(1,3(.,4))"},
		{"already balanced", mk(4, mk(2, leaf(1), leaf(3)), mk(6, leaf(5), leaf(7))), "4(2(1,3),6(5,7))"},
		{"uneven", mk(10, mk(5, leaf(1), nil), mk(20, nil, mk(30, nil, mk(40, nil, leaf(50))))), "20(5(1,10),40(30,50))"},
	}
	for _, c := range cases {
		before := shape(c.root)
		got := Balance(c.root)
		if shape(got) != c.want {
			t.Errorf("%s: Balance(%s) = %s, want %s", c.name, before, shape(got), c.want)
		}
		if shape(c.root) != before {
			t.Errorf("%s: Balance modified its input: was %s, now %s", c.name, before, shape(c.root))
		}
		orig := map[*Node]bool{}
		collect(c.root, orig)
		res := map[*Node]bool{}
		collect(got, res)
		for n := range res {
			if orig[n] {
				t.Errorf("%s: Balance reused a node of the input (value %d); it must build new nodes", c.name, n.Val)
				break
			}
		}
	}
}

func TestBalanceIsBalanced(t *testing.T) {
	for n := 1; n <= 64; n++ {
		got := Balance(rightChain(1, n))
		if !IsBalanced(got) {
			t.Errorf("Balance of a chain of %d nodes is not balanced: %s", n, shape(got))
		}
	}
}
