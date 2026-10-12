# medium / problem2 — ZigZag

Theme: **binary trees**. One task, written in `task.go`.
Run `go test ./problem2` from the `medium/` folder to check.

## Instructions

A company draws its org chart as a binary tree: the CEO at the top, then each
level of managers below. The chart is printed on a long paper strip by a
plotter that saves time by never moving back: it prints one level left to
right, the next level right to left, then left to right again, and so on.

Write a function `ZigZag` that returns the values of the tree level by level,
in the order the plotter prints them.

## Allowed packages

None.

## Expected type and function

```go
package problem2

type Node struct {
    Val         int
    Left, Right *Node
}

func ZigZag(root *Node) [][]int
```

## Rules

1. Level `0` is the root alone. The children of the nodes on level `n` make up
   level `n+1`.
2. The result has one slice per level, from the top level down.
3. Level `0` is read **left to right**, level `1` **right to left**, level `2`
   left to right again, and so on.
4. "Left to right" means the order the nodes appear on the level when the tree
   is drawn, and missing children are simply skipped.
5. A `nil` tree returns an empty result (`nil` is fine).
6. The tree must not be modified.

## Examples

```text
          1
        /   \
       2     3
      / \     \
     4   5     6
    /     \   / \
   7       8 9   10
```

```go
ZigZag(root)
// -> [][]int{
//     {1},
//     {3, 2},          right to left
//     {4, 5, 6},       left to right
//     {10, 9, 8, 7},   right to left
// }
```

```go
ZigZag(&Node{Val: 1})  // -> [][]int{{1}}
ZigZag(nil)            // -> [][]int{}   (or nil)
```
