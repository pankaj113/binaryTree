# easy / battery2 — CountLeaves, MaxValue, Contains

Theme: **binary trees**. Three small tasks, all written in `tasks.go`.
Run `go test ./battery2` from the `easy/` folder to check.

## Instructions

A binary tree is made of nodes. Each node has a value and at most two
children, `Left` and `Right`. A missing child is `nil`. An empty tree is a
`nil` root.

The trees in this battery are **not** sorted in any way: a value can be
anywhere, and the same value can appear more than once.

## Allowed packages

None.

## Expected type and functions

```go
package battery2

type Node struct {
    Val         int
    Left, Right *Node
}

func CountLeaves(root *Node) int
func MaxValue(root *Node) (int, bool)
func Contains(root *Node, v int) bool
```

## Rules

1. A **leaf** is a node with no children (`Left` and `Right` are both `nil`).
2. `CountLeaves` returns how many leaves the tree has. A `nil` tree has `0`.
3. `MaxValue` returns the largest value in the tree and `true`.
4. `MaxValue(nil)` returns `0, false`.
5. Values can be negative.
6. `Contains` returns `true` if any node of the tree has the value `v`.
   `Contains(nil, v)` is `false`.

## Examples

```text
        2
       / \
      7   5
     / \   \
    2   6   9
       / \  /
      5  11 4
```

```go
CountLeaves(root)  // -> 4      (2, 5, 11 and 4)
MaxValue(root)     // -> 11, true
Contains(root, 4)  // -> true
Contains(root, 3)  // -> false

CountLeaves(nil)   // -> 0
MaxValue(nil)      // -> 0, false
Contains(nil, 4)   // -> false
```

```text
      -10
      /  \
    -3    -7
    /
  -50
```

```go
MaxValue(root)     // -> -3, true
```
