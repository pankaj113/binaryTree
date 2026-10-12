package problem3

// Node is one node of a binary search tree.
type Node struct {
	Val         int
	Left, Right *Node
}

// IsBalanced reports whether, at every node, the heights of the two subtrees differ by at most 1.
func IsBalanced(root *Node) bool {
	// TODO: implement
	return false
}

// Balance returns a new, balanced binary search tree holding the same values as root.
func Balance(root *Node) *Node {
	// TODO: implement
	return &Node{Val: -1}
}
