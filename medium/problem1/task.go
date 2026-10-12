package problem1

// Node is one node of a binary tree.
type Node struct {
	Val         int
	Left, Right *Node
}

// IsValidBST reports whether the tree is a valid binary search tree.
func IsValidBST(root *Node) bool {
	if root == nil {
		return true
	}

	l := root.Left
	r := root.Right

	if l == nil && r == nil {
		return true
	}

	if l != nil {
		if l.Val >= root.Val {
			return false
		}
	}

	if r != nil {
		if root.Val >= r.Val {
			return false
		}
	}
	return IsValidBST(l) && IsValidBST(r)
}

func MinMax(root *Node) (min, max int) {
	if root == nil {
		return 0, 0
	}

}
