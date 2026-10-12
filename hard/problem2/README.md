# hard / problem2 — Build

Theme: **binary trees**. One task, written in `task.go`.
Run `go test ./problem2` from the `hard/` folder to check.

## Instructions

A debugging tool printed two lists for a tree: its values in **pre-order**
(node, left subtree, right subtree) and in **in-order** (left subtree, node,
right subtree). The tree itself is gone. When all values are different, these
two lists are enough to rebuild the exact tree.

Write a function `Build` that rebuilds the tree from its two traversals, and
detects lists that cannot belong to any tree.

## Allowed packages

None.

## Expected type and function

```go
package problem2

type Node struct {
    Val         int
    Left, Right *Node
}

func Build(preorder, inorder []int) (*Node, bool)
```

## Rules

1. Return the tree whose pre-order traversal is `preorder` and whose in-order
   traversal is `inorder`, and `true`.
2. Two empty lists (`nil` or zero-length) give the empty tree: `nil, true`.
3. Return `nil, false` if:
   - the two lists have different lengths
   - a value appears more than once in a list
   - no binary tree has exactly these two traversals
4. The tree is a plain binary tree, not a binary search tree: values can be in
   any order.
5. `preorder` and `inorder` must not be modified.

## Examples

```text
    3
   / \
  9   20
     /  \
    15   7
```

```go
Build([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})  // -> the tree above, true

Build([]int{1, 2, 3}, []int{3, 2, 1})   // -> true:     1
                                        //             /
                                        //            2
                                        //           /
                                        //          3

Build([]int{1, 2, 3}, []int{1, 2, 3})   // -> true:  1
                                        //            \
                                        //             2
                                        //              \
                                        //               3

Build(nil, nil)                         // -> nil, true
Build([]int{7}, []int{7})               // -> a single node 7, true

Build([]int{1, 2}, []int{1})            // -> nil, false   (different lengths)
Build([]int{1, 1}, []int{1, 1})         // -> nil, false   (1 appears twice)
Build([]int{1, 2}, []int{1, 3})         // -> nil, false   (not the same values)
Build([]int{1, 2, 3}, []int{3, 1, 2})   // -> nil, false   (no tree has these traversals)
```

## Hint

The first value of `preorder` is the root. Find it in `inorder`: everything to
its left there is the left subtree, everything to its right is the right
subtree.
