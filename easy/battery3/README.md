# easy / battery3 — PreOrder, InOrder, PostOrder

Theme: **binary trees**. Three small tasks, all written in `tasks.go`.
Run `go test ./battery3` from the `easy/` folder to check.

## Instructions

To **traverse** a tree is to visit every node once. With a binary tree there
are three classic orders, and they differ only in *when* the node itself is
visited compared to its two subtrees.

Write one function for each order. Each returns the values in the order they
are visited.

## Allowed packages

None.

## Expected type and functions

```go
package battery3

type Node struct {
    Val         int
    Left, Right *Node
}

func PreOrder(root *Node) []int
func InOrder(root *Node) []int
func PostOrder(root *Node) []int
```

## Rules

1. `PreOrder`: the node first, then its whole left subtree, then its whole
   right subtree.
2. `InOrder`: the whole left subtree, then the node, then the whole right
   subtree.
3. `PostOrder`: the whole left subtree, then the whole right subtree, then the
   node.
4. A `nil` tree returns an empty slice (`nil` is fine).

## Examples

```text
    1
   / \
  2   3
 / \
4   5
```

```go
PreOrder(root)   // -> [1 2 4 5 3]
InOrder(root)    // -> [4 2 5 1 3]
PostOrder(root)  // -> [4 5 2 3 1]
PreOrder(nil)    // -> []
```

```text
      8
     / \
    3   10
   / \    \
  1   6    14
     / \   /
    4   7 13
```

```go
PreOrder(root)   // -> [8 3 1 6 4 7 10 14 13]
InOrder(root)    // -> [1 3 4 6 7 8 10 13 14]   (this tree is sorted: in-order gives the values small to big)
PostOrder(root)  // -> [1 4 7 6 3 13 14 10 8]
```
