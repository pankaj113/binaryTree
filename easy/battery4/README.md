# easy / battery4 — IsSameTree, Mirror, IsSymmetric

Theme: **binary trees**. Three small tasks, all written in `tasks.go`.
Run `go test ./battery4` from the `easy/` folder to check.

## Instructions

These three tasks are about the **shape** of a tree: where each node sits, and
not only which values the tree holds.

## Allowed packages

None.

## Expected type and functions

```go
package battery4

type Node struct {
    Val         int
    Left, Right *Node
}

func IsSameTree(a, b *Node) bool
func Mirror(root *Node) *Node
func IsSymmetric(root *Node) bool
```

## Rules

1. `IsSameTree` returns `true` when `a` and `b` have the same shape and the
   same value in every position. Two `nil` trees are the same.
2. `Mirror` returns the mirror image of the tree: at every node, the left and
   right children swap places.
3. `Mirror` builds a **new** tree. The input tree must not be modified, and the
   result must not share any node with it.
4. `Mirror(nil)` returns `nil`.
5. `IsSymmetric` returns `true` when the tree is the same as its own mirror
   image. A `nil` tree is symmetric.

## Examples

```text
     a          b          c
     1          1          1
    /            \        / \
   2              2      2   3
```

```go
IsSameTree(a, a)    // -> true
IsSameTree(a, b)    // -> false    (same values, but 2 is on the other side)
IsSameTree(c, c)    // -> true
IsSameTree(nil, nil) // -> true
IsSameTree(a, nil)  // -> false
```

```text
    root                Mirror(root)
      4                     4
     / \                   / \
    2   7                 7   2
   / \   \               /   / \
  1   3   9             9   3   1
```

```text
      s1               s2               s3
      1                1                1
     / \              / \              / \
    2   2            2   2            2   2
   / \ / \            \   \            \  /
  3  4 4  3            3   3           3 3
```

```go
IsSymmetric(s1)    // -> true
IsSymmetric(s2)    // -> false    (both 3s hang on the right: not mirrored)
IsSymmetric(s3)    // -> true     (one 3 on the right, the other on the left)
IsSymmetric(root)  // -> false    (the 4-2-7 tree above)
IsSymmetric(nil)   // -> true
```
