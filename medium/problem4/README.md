# medium / problem4 — Range, KthSmallest

Theme: **binary search trees**. One task with two functions, written in
`task.go`. Run `go test ./problem4` from the `medium/` folder to check.

## Instructions

A shop keeps all its product prices in a binary search tree. Customers filter
products by price ("between 25 and 55 euros"), and the shop's front page shows
"the 3rd cheapest product".

Write two functions:

- `Range` returns every price between `lo` and `hi`.
- `KthSmallest` returns the `k`-th smallest price.

## Allowed packages

None.

## Expected type and functions

```go
package problem4

type Node struct {
    Val         int
    Left, Right *Node
}

func Range(root *Node, lo, hi int) []int
func KthSmallest(root *Node, k int) (int, bool)
```

## Rules

1. The input is always a valid binary search tree with unique values.
2. `Range` returns every value `v` with `lo <= v <= hi`, smallest first.
3. If no value is in range (including when `lo > hi`, or the tree is `nil`),
   `Range` returns an empty slice (`nil` is fine).
4. `KthSmallest(root, 1)` returns the smallest value and `true`,
   `KthSmallest(root, 2)` the second smallest, and so on.
5. If `k < 1`, or the tree has fewer than `k` values, `KthSmallest` returns
   `0, false`.
6. The tree must not be modified.

## Examples

```text
            40
          /    \
        20      60
       /  \    /  \
     10   30  50   80
       \          /
       15        70
```

```go
Range(root, 25, 55)    // -> [30 40 50]
Range(root, 15, 30)    // -> [15 20 30]      (both ends are included)
Range(root, 31, 39)    // -> []
Range(root, 60, 20)    // -> []              (lo > hi)
Range(nil, 0, 100)     // -> []

KthSmallest(root, 1)   // -> 10, true
KthSmallest(root, 5)   // -> 40, true
KthSmallest(root, 9)   // -> 80, true
KthSmallest(root, 10)  // -> 0, false        (only 9 values)
KthSmallest(root, 0)   // -> 0, false
```
