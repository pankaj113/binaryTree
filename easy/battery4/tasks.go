package battery4

// Node is one node of a binary tree.
type Node struct {
	Val         int
	Left, Right *Node
}

// IsSameTree reports whether a and b have the same shape and the same values.
func IsSameTree(a, b *Node) bool {
	if a == nil && b == nil {
		return true
	}
	// if (a != nil && b == nil) || (a == nil && b != nil) {
	// 	return false
	// }
	if a == nil || b ==
		nil {
		return false
	}

	if a.Val != b.Val {
		return false
	}

	// if a.Val == b.Val {
	// 	return IsSameTree(a.Left, b.Left) && IsSameTree(a.Right, b.Right)
	// }
	return IsSameTree(a.Left, b.Left) && IsSameTree(a.Right, b.Right)
	// return false
}

// Mirror returns a new tree that is the mirror image of root. root is not modified.
func Mirror(root *Node) *Node {
	if root == nil {
		return nil
	}

	mirror := &Node{
		Val:   root.Val,
		Left:  Mirror(root.Right),
		Right: Mirror(root.Left),
	}

	return mirror
}

// IsSymmetric reports whether the tree is its own mirror image.
func IsSymmetric(root *Node) bool {

	return IsSameTree(root, Mirror(root))
}
