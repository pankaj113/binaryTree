# medium / problem3 — Delete

Theme: **binary search trees**. One task, written in `task.go`.
Run `go test ./problem3` from the `medium/` folder to check.

## Instructions

A parking garage keeps the numbers of the occupied spots in a binary search
tree. When a car leaves, its spot number has to be removed, and the tree must
stay a valid binary search tree afterwards.

Write a function `Delete` that removes a value from a binary search tree and
returns the root of the resulting tree.

## Allowed packages

None.

## Expected type and function

```go
package problem3

type Node struct {
    Val         int
    Left, Right *Node
}

func Delete(root *Node, v int) *Node
```

## Rules

1. The input is always a valid binary search tree with unique values.
2. If `v` is not in the tree, nothing changes. `Delete(nil, v)` returns `nil`.
3. A node with **no children** is simply removed.
4. A node with **one child** is replaced by that child (and the child's whole
   subtree comes with it).
5. A node with **two children** keeps its place in the tree, but takes the
   value of its **successor**: the smallest value in its right subtree. Then
   that successor is deleted from the right subtree, using these same rules.
6. Return the root of the tree. It changes when the root itself is deleted.
7. `Delete` may modify the tree it is given.

## Examples

```text
          50
        /    \
      30      70
     /  \    /  \
   20   40  60   80
             \
              65
```

```go
Delete(root, 20)   // a leaf: removed (rule 3)

          50
        /    \
      30      70
        \    /  \
        40  60   80
             \
              65
```

```go
Delete(root, 60)   // one child: 65 takes its place (rule 4)

          50
        /    \
      30      70
     /  \    /  \
   20   40  65   80
```

```go
Delete(root, 50)   // two children: the successor is 60 (rule 5)

          60
        /    \
      30      70
     /  \    /  \
   20   40  65   80
```

```go
Delete(root, 55)          // -> the same tree, nothing changes
Delete(&Node{Val: 5}, 5)  // -> nil
```
