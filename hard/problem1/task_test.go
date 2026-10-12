package problem1

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
//	  1
//	 / \
//	2   3
//	   / \
//	  4   5
func sample() *Node { return mk(1, leaf(2), mk(3, leaf(4), leaf(5))) }

var serializeCases = []struct {
	name string
	root *Node
	want string
}{
	{"nil tree", nil, "#"},
	{"one node", leaf(7), "7,#,#"},
	{"sample", sample(), "1,2,#,#,3,4,#,#,5,#,#"},
	{"left chain", mk(1, mk(2, leaf(3), nil), nil), "1,2,3,#,#,#,#"},
	{"right chain", mk(1, nil, mk(2, nil, leaf(3))), "1,#,2,#,3,#,#"},
	{"negatives and zero", mk(-1, leaf(0), leaf(-20)), "-1,0,#,#,-20,#,#"},
	{"big numbers", mk(9223372036854775807, leaf(-9223372036854775808), nil), "9223372036854775807,-9223372036854775808,#,#,#"},
	{"repeated values", mk(5, leaf(5), leaf(5)), "5,5,#,#,5,#,#"},
}

func TestSerialize(t *testing.T) {
	for _, c := range serializeCases {
		if got := Serialize(c.root); got != c.want {
			t.Errorf("%s: Serialize(%s) = %q, want %q", c.name, shape(c.root), got, c.want)
		}
	}
}

func TestDeserializeValid(t *testing.T) {
	for _, c := range serializeCases {
		got, ok := Deserialize(c.want)
		if !ok || shape(got) != shape(c.root) {
			t.Errorf("%s: Deserialize(%q) = %s, %v, want %s, true", c.name, c.want, shape(got), ok, shape(c.root))
		}
	}
	extra := []struct {
		name string
		in   string
		want string
	}{
		{"leading zeros are fine", "007,#,#", "7"},
		{"minus zero is zero", "-0,#,#", "0"},
	}
	for _, c := range extra {
		got, ok := Deserialize(c.in)
		if !ok || shape(got) != c.want {
			t.Errorf("%s: Deserialize(%q) = %s, %v, want %s, true", c.name, c.in, shape(got), ok, c.want)
		}
	}
}

func TestDeserializeInvalid(t *testing.T) {
	cases := []struct {
		name string
		in   string
	}{
		{"empty string", ""},
		{"value without children", "1"},
		{"missing right child", "1,#"},
		{"missing deep", "1,2,#,#,3,#"},
		{"extra after the tree", "1,#,#,#"},
		{"two trees", "#,#"},
		{"extra value after the tree", "1,#,#,2,#,#"},
		{"empty token", "1,,#"},
		{"trailing comma", "1,#,#,"},
		{"leading comma", ",1,#,#"},
		{"space", "1, #,#"},
		{"spaces around", " 1,#,# "},
		{"plus sign", "+1,#,#"},
		{"lone minus", "-,#,#"},
		{"double minus", "--1,#,#"},
		{"letters", "a,#,#"},
		{"digits then letters", "1a,#,#"},
		{"decimal", "1.5,#,#"},
		{"double hash", "1,##,#"},
		{"too big for int", "9223372036854775808,#,#"},
		{"too small for int", "-9223372036854775809,#,#"},
		{"other separator", "1;#;#"},
	}
	for _, c := range cases {
		got, ok := Deserialize(c.in)
		if ok || got != nil {
			t.Errorf("%s: Deserialize(%q) = %s, %v, want nil, false", c.name, c.in, shape(got), ok)
		}
	}
}

func TestRoundTrip(t *testing.T) {
	root := mk(10, mk(-3, nil, mk(4, leaf(4), nil)), mk(0, leaf(100), mk(7, nil, leaf(-8))))
	s := Serialize(root)
	back, ok := Deserialize(s)
	if !ok || shape(back) != shape(root) {
		t.Errorf("round trip: Deserialize(Serialize(%s)) = %s, %v", shape(root), shape(back), ok)
	}
	if Serialize(back) != s {
		t.Errorf("round trip: Serialize(Deserialize(%q)) = %q", s, Serialize(back))
	}
}
