package battery3

// Node is one node of a binary tree.
type Node struct {
	Val         int
	Left, Right *Node
}

// PreOrder returns the values in the order: node, left subtree, right subtree.
func PreOrder(root *Node) []int {
	// if root == nil {
	// 	return nil
	// }
	// var result []int
	// result = append(result, root.Val)
	// l := root.Left
	// r := root.Right
	// if l != nil {
	// 	result = append(result, PreOrder(l)...)
	// }
	// if r != nil {
	// 	result = append(result, PreOrder(r)...)
	// }
	// return result

	var finalResult []int
	PreOrderHelper(root, &finalResult)
	return finalResult
}

func PreOrderHelper(node *Node, accumulator *[]int) {
	if node == nil {
		return
	}
	*accumulator = append(*accumulator, node.Val)

	PreOrderHelper(node.Left, accumulator)
	PreOrderHelper(node.Right, accumulator)
}

// InOrder returns the values in the order: left subtree, node, right subtree.
func InOrder(root *Node) []int {
	// if root == nil {
	// 	return nil
	// }
	// var result []int

	// l := root.Left
	// if l != nil {
	// 	result = append(result, InOrder(l)...)
	// }

	// result = append(result, root.Val)

	// r := root.Right
	// if r != nil {
	// 	result = append(result, InOrder(r)...)
	// }

	// return result

	var finalResult []int
	InOrderHelper(root, &finalResult)
	return finalResult
}

func InOrderHelper(node *Node, accumulator *[]int) {
	if node == nil {
		return
	}
	InOrderHelper(node.Left, accumulator)
	*accumulator = append(*accumulator, node.Val)
	InOrderHelper(node.Right, accumulator)
}

// PostOrder returns the values in the order: left subtree, right subtree, node.
func PostOrder(root *Node) []int {
	// if root == nil {
	// 	return nil
	// }
	// var result []int
	// l := root.Left
	// if l != nil {
	// 	result = append(result, PostOrder(l)...)
	// }
	// r := root.Right
	// if r != nil {
	// 	result = append(result, PostOrder(r)...)
	// }
	// result = append(result, root.Val)
	// return result
	var finalResult []int
	PostOrderHelper(root, &finalResult)
	return finalResult
}

func PostOrderHelper(node *Node, accumulator *[]int) {
	if node == nil {
		return
	}
	PostOrderHelper(node.Left, accumulator)
	PostOrderHelper(node.Right, accumulator)
	*accumulator = append(*accumulator, node.Val)
}
