package battery2

// Node is one node of a binary tree.
type Node struct {
	Val         int
	Left, Right *Node
}

// CountLeaves returns how many nodes have no children.
func CountLeaves(root *Node) int {
	if root == nil {
		return 0
	}
	l := root.Left
	r := root.Right

	if l == nil && r == nil {
		return 1
	}
	return CountLeaves(l) + CountLeaves(r)
}

// MaxValue returns the largest value in the tree, and false if the tree is empty.
func MaxValue(root *Node) (int, bool) {
	if root == nil {
		return 0, false
	}
	max := root.Val
	l := root.Left
	r := root.Right

	if l != nil {
		lM, _ := MaxValue(l)
		if lM > max {
			max = lM
		}
	}
	if r != nil {
		rM, _ := MaxValue(r)
		if rM > max {
			max = rM
		}
	}
	return max, true
}

// Contains reports whether v is somewhere in the tree.
func Contains(root *Node, v int) bool {
	if root == nil {
		return false
	}
	if root.Val == v {
		return true
	}
	l := root.Left
	r := root.Right
	if Contains(l, v) || Contains(r, v) {
		return true
	}
	return false
}
