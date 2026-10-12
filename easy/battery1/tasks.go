package battery1

// Node is one node of a binary tree.
type Node struct {
	Val         int
	Left, Right *Node
}

// Size returns how many nodes the tree has.
func Size(root *Node) int {

	if root == nil {
		return 0
	}

	l := root.Left
	r := root.Right

	return 1 + Size(l) + Size(r)

}

// Sum returns the sum of every value in the tree.
func Sum(root *Node) int {

	if root == nil {
		return 0
	}
	l := root.Left
	r := root.Right

	return root.Val + Sum(l) + Sum(r)
}

// Height returns the number of nodes on the longest path from the root down to a leaf.
func Height(root *Node) int {
	if root == nil {
		return 0
	}

	l := root.Left
	r := root.Right

	return 1 + Max(Height(l), Height(r))
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
