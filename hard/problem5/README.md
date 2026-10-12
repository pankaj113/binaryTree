# hard / problem5 — MaxPathSum

Theme: **binary trees**. One task, written in `task.go`.
Run `go test ./problem5` from the `hard/` folder to check.

## Instructions

A hiking map is drawn as a binary tree: each node is a checkpoint, and its
value is how many points a hiker earns there (negative for a checkpoint that
costs points, like a toll). Trails only run along the edges of the tree, and
no checkpoint can be visited twice.

Write a function `MaxPathSum` that returns the best score a single hike can
make.

## Allowed packages

None.

## Expected type and function

```go
package problem5

type Node struct {
    Val         int
    Left, Right *Node
}

func MaxPathSum(root *Node) (int, bool)
```

## Rules

1. A **path** is a sequence of nodes where each next node is the parent or a
   child of the previous one, and no node appears twice.
2. A path has at least one node. It can start and end at any node: it does not
   have to go through the root, and it does not have to end at a leaf.
3. The score of a path is the sum of its values.
4. Return the largest score of any path in the tree, and `true`.
5. `MaxPathSum(nil)` returns `0, false`.
6. Values can be negative. You do not need to handle integer overflow.
7. The tree must not be modified.

## Examples

```text
     -10
     /  \
    9    20
        /  \
      15    7
```

```go
MaxPathSum(root)   // -> 42, true    (15 -> 20 -> 7)
```

```text
          5
         / \
        4   8
       /   / \
      11  13  4
     /  \      \
    7    2      1
```

```go
MaxPathSum(root)   // -> 48, true    (7 -> 11 -> 4 -> 5 -> 8 -> 13)
```

```text
     1
    / \
   2   3
```

```go
MaxPathSum(root)   // -> 6, true     (2 -> 1 -> 3)
```

```text
     -3
     / \
   -2   -1
```

```go
MaxPathSum(root)   // -> -1, true    (the path of just one node, -1)
MaxPathSum(nil)    // -> 0, false
```

## Hint

At each node, work out two numbers: the best path that **ends** at this node
going down into one side (this is what the parent can extend), and the best
path that **bends** through this node, using both sides. A side that only
lowers the score can be left out.
