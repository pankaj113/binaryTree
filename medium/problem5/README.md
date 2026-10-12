# medium / problem5 — CommonAncestor

Theme: **binary search trees**. One task, written in `task.go`.
Run `go test ./problem5` from the `medium/` folder to check.

## Instructions

A network of cables is laid out as a binary search tree: each node is a
junction box with a number, and the boxes are numbered so that the tree is in
order. When two boxes need to talk to each other, the signal goes up to the
closest junction they share, and down again.

Write a function `CommonAncestor` that finds that shared junction: the lowest
node of the tree that has both `a` and `b` in its subtree.

## Allowed packages

None.

## Expected type and function

```go
package problem5

type Node struct {
    Val         int
    Left, Right *Node
}

func CommonAncestor(root *Node, a, b int) (int, bool)
```

## Rules

1. The input is always a valid binary search tree with unique values.
2. A node's subtree includes the node itself. So if `a` is above `b` in the
   tree, the answer is `a`.
3. Return the value of the lowest node whose subtree holds both `a` and `b`,
   and `true`.
4. The order of `a` and `b` does not matter. `a` and `b` may be equal.
5. If `a` or `b` is not in the tree (or the tree is `nil`), return `0, false`.
6. The tree must not be modified.

## Examples

```text
          60
        /    \
      20      80
     /  \       \
   10    40      90
        /  \
      30    50
```

```go
CommonAncestor(root, 30, 50)   // -> 40, true
CommonAncestor(root, 10, 50)   // -> 20, true
CommonAncestor(root, 50, 10)   // -> 20, true    (order does not matter)
CommonAncestor(root, 10, 90)   // -> 60, true
CommonAncestor(root, 20, 30)   // -> 20, true    (20 is above 30)
CommonAncestor(root, 40, 40)   // -> 40, true
CommonAncestor(root, 35, 50)   // -> 0, false    (there is no 35)
CommonAncestor(nil, 1, 2)      // -> 0, false
```
