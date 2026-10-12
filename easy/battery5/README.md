# easy / battery5 — Insert, Search, MinMax

Theme: **binary search trees**. Three small tasks, all written in `tasks.go`.
Run `go test ./battery5` from the `easy/` folder to check.

## Instructions

A **binary search tree** (BST) is a binary tree kept in order. For every
node:

- every value in its **left** subtree is **smaller** than the node's value
- every value in its **right** subtree is **bigger** than the node's value

Because of this order, you never need to look at the whole tree: at each node
you know which side to go to.

## Allowed packages

None.

## Expected type and functions

```go
package battery5

type Node struct {
    Val         int
    Left, Right *Node
}

func Insert(root *Node, v int) *Node
func Search(root *Node, v int) bool
func MinMax(root *Node) (min, max int, ok bool)
```

## Rules

1. `Insert` adds `v` as a new leaf in the place where it keeps the tree in
   order, and returns the root.
2. `Insert(nil, v)` returns a new single node holding `v`.
3. If `v` is already in the tree, `Insert` changes nothing.
4. `Insert` may modify the tree it is given. For a non-empty tree, the root it
   returns is the same root it was given.
5. `Search` returns `true` if `v` is in the tree. It only follows one path
   down: left when `v` is smaller than the node, right when it is bigger.
6. `MinMax` returns the smallest value, the largest value and `true`.
7. `MinMax(nil)` returns `0, 0, false`.

## Examples

```text
      8
     / \
    3   10
   / \    \
  1   6    14
```

```go
var root *Node
for _, v := range []int{8, 3, 10, 1, 6, 14} {
    root = Insert(root, v)
}
// root is now the tree above

Insert(root, 7)     // -> 7 becomes the right child of 6
Insert(root, 8)     // -> nothing changes (8 is already there)

Search(root, 6)     // -> true
Search(root, 9)     // -> false
Search(nil, 9)      // -> false

MinMax(root)        // -> 1, 14, true
MinMax(nil)         // -> 0, 0, false
```

Rule 5: `Search` trusts the order. In this broken tree, `2` is on the wrong
side of `5`, so a search never finds it:

```text
    5
   / \
  3   8
     /
    2
```

```go
Search(broken, 2)   // -> false
```
