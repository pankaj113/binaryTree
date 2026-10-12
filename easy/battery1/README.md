# easy / battery1 — Size, Sum, Height

Theme: **binary trees**. Three small tasks, all written in `tasks.go`.
Run `go test ./battery1` from the `easy/` folder to check.

## Instructions

A binary tree is made of nodes. Each node has a value and at most two
children, `Left` and `Right`. A missing child is `nil`. The tree is given by
its top node, the **root**. An empty tree is a `nil` root.

Write three functions that each look at every node of the tree once.

## Allowed packages

None.

## Expected type and functions

```go
package battery1

type Node struct {
    Val         int
    Left, Right *Node
}

func Size(root *Node) int
func Sum(root *Node) int
func Height(root *Node) int
```

## Rules

1. `Size` returns how many nodes the tree has.
2. `Sum` returns the sum of all values in the tree.
3. `Height` returns how many nodes are on the longest path from the root down
   to a leaf (a node with no children).
4. For a `nil` tree, all three return `0`.

## Examples

All the examples use this tree:

```text
      5
     / \
    3   8
   / \   \
  1   4   9
```

```go
Size(root)    // -> 6
Sum(root)     // -> 30
Height(root)  // -> 3     (5 -> 3 -> 1, or 5 -> 8 -> 9)

Size(nil)     // -> 0
Sum(nil)      // -> 0
Height(nil)   // -> 0

one := &Node{Val: 7}
Size(one)     // -> 1
Sum(one)      // -> 7
Height(one)   // -> 1
```
