package battery4

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

// sample:
//
//	    4
//	   / \
//	  2   7
//	 / \   \
//	1   3   9
func sample() *Node { return mk(4, mk(2, leaf(1), leaf(3)), mk(7, nil, leaf(9))) }

// symmetric:
//
//	    1
//	   / \
//	  2   2
//	 / \ / \
//	3  4 4  3
func symmetric() *Node {
	return mk(1, mk(2, leaf(3), leaf(4)), mk(2, leaf(4), leaf(3)))
}

func TestIsSameTree(t *testing.T) {
	cases := []struct {
		name string
		a, b *Node
		want bool
	}{
		{"both nil", nil, nil, true},
		{"one nil", leaf(1), nil, false},
		{"other nil", nil, leaf(1), false},
		{"same single", leaf(1), leaf(1), true},
		{"different single", leaf(1), leaf(2), false},
		{"same sample", sample(), sample(), true},
		{"same values, child on other side", mk(1, leaf(2), nil), mk(1, nil, leaf(2)), false},
		{"one value differs deep", sample(), mk(4, mk(2, leaf(1), leaf(3)), mk(7, nil, leaf(8))), false},
		{"extra node", sample(), mk(4, mk(2, leaf(1), leaf(3)), mk(7, leaf(5), leaf(9))), false},
		{"tree and its mirror", symmetric(), symmetric(), true},
	}
	for _, c := range cases {
		if got := IsSameTree(c.a, c.b); got != c.want {
			t.Errorf("%s: IsSameTree(%s, %s) = %v, want %v", c.name, shape(c.a), shape(c.b), got, c.want)
		}
	}
}

func TestMirror(t *testing.T) {
	cases := []struct {
		name string
		root *Node
		want string
	}{
		{"nil tree", nil, "."},
		{"one node", leaf(5), "5"},
		{"sample", sample(), "4(7(9,.),2(3,1))"},
		{"left chain", mk(1, mk(2, leaf(3), nil), nil), "1(.,2(.,3))"},
		{"symmetric", symmetric(), "1(2(3,4),2(4,3))"},
	}
	for _, c := range cases {
		before := shape(c.root)
		got := Mirror(c.root)
		if shape(got) != c.want {
			t.Errorf("%s: Mirror(%s) = %s, want %s", c.name, before, shape(got), c.want)
		}
		if shape(c.root) != before {
			t.Errorf("%s: Mirror modified its input: was %s, now %s", c.name, before, shape(c.root))
		}
		orig := map[*Node]bool{}
		collect(c.root, orig)
		res := map[*Node]bool{}
		collect(got, res)
		for n := range res {
			if orig[n] {
				t.Errorf("%s: Mirror reused a node of the input (value %d); it must build new nodes", c.name, n.Val)
				break
			}
		}
	}
}

func TestIsSymmetric(t *testing.T) {
	cases := []struct {
		name string
		root *Node
		want bool
	}{
		{"nil tree", nil, true},
		{"one node", leaf(1), true},
		{"symmetric", symmetric(), true},
		{"sample", sample(), false},
		{"same shape, values differ", mk(1, mk(2, leaf(3), leaf(4)), mk(2, leaf(3), leaf(4))), false},
		{"same values, shape differs", mk(1, mk(2, nil, leaf(3)), mk(2, nil, leaf(3))), false},
		{"mirrored shape", mk(1, mk(2, nil, leaf(3)), mk(2, leaf(3), nil)), true},
		{"only one child", mk(1, leaf(2), nil), false},
		{"two equal children", mk(1, leaf(2), leaf(2)), true},
	}
	for _, c := range cases {
		if got := IsSymmetric(c.root); got != c.want {
			t.Errorf("%s: IsSymmetric(%s) = %v, want %v", c.name, shape(c.root), got, c.want)
		}
	}
}
