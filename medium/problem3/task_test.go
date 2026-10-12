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

// sample is the tree from the README:
//
//	       50
//	     /    \
//	   30      70
//	  /  \    /  \
//	20   40  60   80
//	         \
//	          65
func sample() *Node {
	return mk(50,
		mk(30, leaf(20), leaf(40)),
		mk(70, mk(60, nil, leaf(65)), leaf(80)))
}

func TestDelete(t *testing.T) {
	cases := []struct {
		name string
		root *Node
		v    int
		want string
	}{
		{"nil tree", nil, 5, "."},
		{"only node", leaf(5), 5, "."},
		{"missing value", sample(), 55, "50(30(20,40),70(60(.,65),80))"},
		{"missing, bigger than all", sample(), 99, "50(30(20,40),70(60(.,65),80))"},
		{"a leaf", sample(), 20, "50(30(.,40),70(60(.,65),80))"},
		{"another leaf", sample(), 80, "50(30(20,40),70(60(.,65),.))"},
		{"one child (right)", sample(), 60, "50(30(20,40),70(65,80))"},
		{"one child (left)", mk(10, mk(5, leaf(2), nil), nil), 5, "10(2,.)"},
		{"two children, successor is a leaf", sample(), 30, "50(40(20,.),70(60(.,65),80))"},
		{"two children, successor has a right child", sample(), 50, "60(30(20,40),70(65,80))"},
		{"two children, successor is the right child", mk(10, leaf(5), mk(15, nil, leaf(20))), 10, "15(5,20)"},
		{"root with only a left child", mk(10, mk(5, leaf(3), leaf(7)), nil), 10, "5(3,7)"},
		{"root with only a right child", mk(10, nil, mk(15, leaf(12), nil)), 10, "15(12,.)"},
		{"deep successor", mk(10, leaf(5), mk(20, mk(15, mk(12, nil, leaf(13)), nil), leaf(25))), 10,
			"12(5,20(15(13,.),25))"},
	}
	for _, c := range cases {
		before := shape(c.root)
		got := Delete(c.root, c.v)
		if shape(got) != c.want {
			t.Errorf("%s: Delete(%s, %d) = %s, want %s", c.name, before, c.v, shape(got), c.want)
		}
	}
}

func TestDeleteEverything(t *testing.T) {
	root := sample()
	order := []int{50, 20, 70, 30, 65, 40, 80, 60}
	want := []string{
		"60(30(20,40),70(65,80))",
		"60(30(.,40),70(65,80))",
		"60(30(.,40),80(65,.))",
		"60(40,80(65,.))",
		"60(40,80)",
		"60(.,80)",
		"60",
		".",
	}
	for i, v := range order {
		root = Delete(root, v)
		if shape(root) != want[i] {
			t.Fatalf("after deleting %v: got %s, want %s", order[:i+1], shape(root), want[i])
		}
	}
}
