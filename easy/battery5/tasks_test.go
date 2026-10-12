package battery5

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
//	    8
//	   / \
//	  3   10
//	 / \    \
//	1   6    14
func sample() *Node {
	return mk(8, mk(3, leaf(1), leaf(6)), mk(10, nil, leaf(14)))
}

func build(vals ...int) *Node {
	var root *Node
	for _, v := range vals {
		root = Insert(root, v)
	}
	return root
}

func TestInsert(t *testing.T) {
	cases := []struct {
		name string
		vals []int
		want string
	}{
		{"nothing", nil, "."},
		{"one value", []int{5}, "5"},
		{"README order", []int{8, 3, 10, 1, 6, 14}, "8(3(1,6),10(.,14))"},
		{"sorted input makes a chain", []int{1, 2, 3, 4}, "1(.,2(.,3(.,4)))"},
		{"reverse sorted", []int{4, 3, 2, 1}, "4(3(2(1,.),.),.)"},
		{"duplicates are ignored", []int{5, 3, 5, 7, 3, 7}, "5(3,7)"},
		{"negatives", []int{0, -5, 5, -10, -1}, "0(-5(-10,-1),5)"},
		{"zigzag", []int{50, 20, 40, 30}, "50(20(.,40(30,.)),.)"},
	}
	for _, c := range cases {
		if got := shape(build(c.vals...)); got != c.want {
			t.Errorf("%s: inserting %v gives %s, want %s", c.name, c.vals, got, c.want)
		}
	}
}

func TestInsertIntoExisting(t *testing.T) {
	root := sample()
	got := Insert(root, 7)
	if got != root {
		t.Errorf("Insert into a non-empty tree must return the same root")
	}
	if want := "8(3(1,6(.,7)),10(.,14))"; shape(got) != want {
		t.Errorf("Insert(sample, 7) = %s, want %s", shape(got), want)
	}
	got = Insert(got, 13)
	if want := "8(3(1,6(.,7)),10(.,14(13,.)))"; shape(got) != want {
		t.Errorf("Insert(.., 13) = %s, want %s", shape(got), want)
	}
	got = Insert(got, 8)
	if want := "8(3(1,6(.,7)),10(.,14(13,.)))"; shape(got) != want {
		t.Errorf("Insert(.., 8) (already there) = %s, want %s", shape(got), want)
	}
}

func TestSearch(t *testing.T) {
	cases := []struct {
		name string
		root *Node
		v    int
		want bool
	}{
		{"nil tree", nil, 1, false},
		{"root", sample(), 8, true},
		{"leaf on the left", sample(), 1, true},
		{"inner node", sample(), 3, true},
		{"leaf on the right", sample(), 14, true},
		{"missing, smaller than all", sample(), 0, false},
		{"missing, bigger than all", sample(), 100, false},
		{"missing, in between", sample(), 7, false},
		{"missing, under a node with one child", sample(), 9, false},
		// 2 sits on the wrong side of 5: a real search never goes there.
		{"only one path is followed", mk(5, leaf(3), mk(8, leaf(2), nil)), 2, false},
		{"only one path is followed (right)", mk(5, mk(3, nil, leaf(9)), leaf(8)), 9, false},
	}
	for _, c := range cases {
		if got := Search(c.root, c.v); got != c.want {
			t.Errorf("%s: Search(%s, %d) = %v, want %v", c.name, shape(c.root), c.v, got, c.want)
		}
	}
}

func TestMinMax(t *testing.T) {
	cases := []struct {
		name             string
		root             *Node
		wantMin, wantMax int
		wantOK           bool
	}{
		{"nil tree", nil, 0, 0, false},
		{"one node", leaf(4), 4, 4, true},
		{"sample", sample(), 1, 14, true},
		{"root is the min", mk(1, nil, mk(5, leaf(3), nil)), 1, 5, true},
		{"root is the max", mk(9, mk(2, nil, leaf(7)), nil), 2, 9, true},
		{"negatives", mk(-5, leaf(-20), leaf(-1)), -20, -1, true},
		// The min is not the left-most leaf: 3 has no left child.
		{"min has a right child", mk(10, mk(3, nil, leaf(6)), leaf(12)), 3, 12, true},
	}
	for _, c := range cases {
		mn, mx, ok := MinMax(c.root)
		if mn != c.wantMin || mx != c.wantMax || ok != c.wantOK {
			t.Errorf("%s: MinMax(%s) = %d, %d, %v, want %d, %d, %v",
				c.name, shape(c.root), mn, mx, ok, c.wantMin, c.wantMax, c.wantOK)
		}
	}
}
