# hard / problem3 — IsBalanced, Balance

Theme: **binary search trees**. One task with two functions, written in
`task.go`. Run `go test ./problem3` from the `hard/` folder to check.

## Instructions

A binary search tree is fast only when it is bushy. If values are inserted in
sorted order, every new node goes to the right, the tree becomes a long chain,
and a search has to walk through every node. A balanced tree with the same
values is much shorter.

Write `IsBalanced`, which checks whether a tree is balanced, and `Balance`,
which builds a balanced binary search tree holding the same values as a given
one.

## Allowed packages

None.

## Expected type and functions

```go
package problem3

type Node struct {
    Val         int
    Left, Right *Node
}

func IsBalanced(root *Node) bool
func Balance(root *Node) *Node
```

## Rules

1. The **height** of a tree is the number of nodes on its longest path from the
   root down to a leaf. A `nil` tree has height `0`.
2. A tree is **balanced** when, at **every** node, the heights of its left and
   right subtrees differ by at most `1`. A `nil` tree is balanced.
3. `IsBalanced` works on any binary tree.
4. `Balance` is given a valid binary search tree with unique values. It returns
   a new binary search tree holding exactly the same values, built like this:
   - put all the values in sorted order
   - the root is the value in the middle; with an even count, take the
     **left** one of the two middle values (index `(len-1)/2`)
   - the left subtree is built the same way from the values before the middle,
     and the right subtree from the values after it
5. `Balance` must not modify its input, and the result must not share any node
   with it. `Balance(nil)` returns `nil`.

## Examples

```text
  balanced           not balanced
       5                   5
      / \                 / \
     2   7               2   7
    / \  /              / \
   1   3 6             1   3
        \                   \
         4                   4
```

```go
IsBalanced(balanced)       // -> true
IsBalanced(notBalanced)    // -> false   (at 5: the left side has height 3, the right side 1)
IsBalanced(nil)            // -> true
```

```text
  chain                   Balance(chain)
  1
   \                            4
    2                         /   \
     \                       2     6
      3          ->         / \   / \
       \                   1   3 5   7
        ...
          \
           7
```

```go
Balance(chain)   // -> the tree on the right: 1..7 sorted, the middle (4) is the root

// Four values 1 2 3 4: the middle is index (4-1)/2 = 1, so 2 is the root.
//      2
//     / \
//    1   3
//         \
//          4
```
