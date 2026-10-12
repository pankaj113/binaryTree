package battery5

// Node is one node of a binary search tree.
type Node struct {
	Val         int
	Left, Right *Node
}

// Insert adds v to the binary search tree and returns the root of the tree.
func Insert(root *Node, v int) *Node {

	if root == nil {
		result := &Node{
			Val: v,
		}
		root = result
	}
	if root.Val > v {
		root.Left = Insert(root.Left, v)
	} else if root.Val < v {
		root.Right = Insert(root.Right, v)
	}
	return root
}

// Search reports whether v is in the binary search tree.
func Search(root *Node, v int) bool {
	if root == nil {
		return false
	}

	if root.Val == v {
		return true
	}

	if root.Val < v {
		return Search(root.Right, v)
	}

	if root.Val > v {
		return Search(root.Left, v)
	}
	return false
}

// MinMax returns the smallest and the largest value of the binary search tree.
func MinMax(root *Node) (min, max int, ok bool) {
	if root == nil {
		return 0, 0, false
	}
	max = root.Val
	min = root.Val

	l := root.Left
	r := root.Right

	if l != nil {
		lMin, _, _ := MinMax(l)
		if lMin < min {
			min = lMin
		}
	}
	if r != nil {
		_, rMax, _ := MinMax(r)
		if rMax > max {
			max = rMax
		}
	}
	return min, max, true
}
