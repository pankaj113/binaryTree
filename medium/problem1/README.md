# medium / problem1 — IsValidBST

Theme: **binary search trees**. One task, written in `task.go`.
Run `go test ./problem1` from the `medium/` folder to check.

## Instructions

A library keeps its book index as a binary search tree of book IDs. After a
crash, a backup was restored, and nobody is sure the tree is still in order.
Before the index is used again, it has to be checked.

Write a function `IsValidBST` that reports whether a binary tree is a valid
binary search tree.

## Allowed packages

None.

## Expected type and function

```go
package problem1

type Node struct {
    Val         int
    Left, Right *Node
}

func IsValidBST(root *Node) bool
```

## Rules

1. A tree is a valid BST when, for **every** node:
   - every value in its whole left subtree is smaller than the node's value
   - every value in its whole right subtree is bigger than the node's value
2. Values must be unique: a value equal to an ancestor's value makes the tree
   invalid.
3. A `nil` tree and a single node are valid.
4. Values can be any `int`, including the smallest and largest possible `int`.
5. The tree must not be modified.

## Examples

```text
  valid                   invalid
       20                      20
      /  \                    /  \
    10    30                10    30
   / \   / \               / \
  5  15 25  40            5   25
```

```go
IsValidBST(valid)    // -> true
IsValidBST(invalid)  // -> false   (25 is in the left subtree of 20, but bigger than 20)
IsValidBST(nil)      // -> true
```

```text
    10        10        10
   /            \         \
  15             5         10
```

```go
// -> false   (15 is on the left but bigger)
// -> false   (5 is on the right but smaller)
// -> false   (duplicates are not allowed)
```

```go
root := &Node{Val: 0, Left: &Node{Val: math.MinInt}, Right: &Node{Val: math.MaxInt}}
IsValidBST(root)     // -> true
```
