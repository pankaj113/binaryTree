package problem4

// Node is one node of a binary search tree.
type Node struct {
	Val         int
	Left, Right *Node
}

// Range returns every value v of the tree with lo <= v <= hi, smallest first.
func Range(root *Node, lo, hi int) []int {
	// TODO: implement
	return []int{-1}
}

// KthSmallest returns the k-th smallest value of the tree (k = 1 is the smallest).
func KthSmallest(root *Node, k int) (int, bool) {
	// TODO: implement
	return -1, true
}
