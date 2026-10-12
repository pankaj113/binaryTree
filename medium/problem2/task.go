package problem2

// Node is one node of a binary tree.
type Node struct {
	Val         int
	Left, Right *Node
}

// ZigZag returns the values of the tree level by level, changing direction on every level.
func ZigZag(root *Node) [][]int {
	// TODO: implement
	return [][]int{{-1}}
}
